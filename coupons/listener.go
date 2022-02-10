package coupons

import (
	"errors"
	"fmt"
	"git.xx.network/elixxir/coupons/storage"
	jww "github.com/spf13/jwalterweatherman"
	"gitlab.com/elixxir/client/api"
	"gitlab.com/elixxir/client/interfaces/message"
	"gitlab.com/elixxir/client/interfaces/params"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type listener struct {
	delay time.Duration
	s     *storage.Storage
	c     *api.Client
}

var validResponse = "Thank you! You will receive %d xx in %s within 2 weeks!"
var noWalletFound = "Thank you! You will receive %d xx. The xx team does not have a wallet address for your purchase, please contact %s."
var reusedCode = "Your code %s has already been used. Your wallet %s will receive %d xx within 2 weeks."
var invalidCode = "That is not a valid code. Please send the code you received in your email for participating in the November Community Sale. Please contact %s if you need support."
var accountUsed = "Your messenger account can only redeem once."
var supportEmail = "distribution@xx-coin.io"

// Hear messages from users to the coupon bot & respond appropriately
func (l *listener) Hear(item message.Receive) {
	// Confirm that authenticated channels
	if !l.c.HasAuthenticatedChannel(item.Sender) {
		jww.ERROR.Printf("No authenticated channel exists to %+v", item.Sender)
		return
	}

	// Parse the trigger
	in := &CMIXText{}
	var trigger string
	err := proto.Unmarshal(item.Payload, in)
	if err != nil {
		jww.ERROR.Printf("Could not unmarshal message from messenger: %+v", err)
		return
	} else {
		trigger = in.Text
	}

	jww.INFO.Printf("Received trigger %s [%+v]", trigger, in)
	var strResponse string

	// Check if the user is in the db, if it is return the trigger
	usedTrigger, err := l.s.CheckUser(item.Sender.String())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // Exit if we get error other than record not found
		jww.DEBUG.Printf("Failed to check user with DB: %+v", err)
		return
	}

	// Retrieve coupon code for trigger if it exists
	codeResp, uses, err := l.s.GetCouponCode(trigger)
	if err != nil {
		jww.DEBUG.Printf("No coupon code for trigger %s: %+v", trigger, err)
		strResponse = fmt.Sprintf(invalidCode, supportEmail)
	} else {
		// Split the data in code column
		rl := strings.Split(codeResp, ",")
		// Assign wallet to var
		wallet := rl[1]
		// Parse num coins
		num, err := strconv.Atoi(rl[0])
		if err != nil {
			jww.ERROR.Printf("Could not parse num coins: %+v", err)
			return
		}
		jww.DEBUG.Printf("Found %s with %d uses", codeResp, uses)

		// Response logic
		if usedTrigger != "" && usedTrigger != trigger { // If account has a code used that isn't the one passed in
			jww.INFO.Println("Case 1")
			strResponse = accountUsed
		} else if uses < 1 { // If the code has been used
			jww.INFO.Println("Case 2")
			strResponse = fmt.Sprintf(reusedCode, trigger, wallet, num)
		} else { // Code is good and can be used by this account
			jww.INFO.Println("Case 3")
			err = l.s.UseCode(item.Sender.String(), trigger)
			if err != nil {
				jww.ERROR.Printf("Failed to commit code use to db: %+v", err)
			}
			if wallet == "" { //
				strResponse = fmt.Sprintf(noWalletFound, num, supportEmail)
			} else {
				strResponse = fmt.Sprintf(validResponse, num, wallet)
			}
		}
	}

	payload := &CMIXText{
		Version: 0,
		Text:    strResponse,
		Reply: &TextReply{
			MessageId: item.ID.Marshal(),
			SenderId:  item.Sender.Marshal(),
		},
	}
	marshalled, err := proto.Marshal(payload)
	if err != nil {
		jww.ERROR.Printf("Failed to marshal payload: %+v", err)
		return
	}
	// Create response message
	resp := message.Send{
		Recipient:   item.Sender,
		Payload:     marshalled,
		MessageType: message.XxMessage,
	}

	// Send response message to sender over cmix
	rids, mid, t, err := l.c.SendE2E(resp, params.GetDefaultE2E())
	if err != nil {
		jww.ERROR.Printf("Failed to send message: %+v", err)
	} else {
		jww.INFO.Printf("Sent response %s [%+v] to %+v on rounds %+v [%+v]", strResponse, mid, item.Sender.String(), rids, t)
	}
}

// Name returns a name, used for debugging
func (l *listener) Name() string {
	return "Coupons-bot-listener"
}
