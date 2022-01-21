package coupons

import (
	"fmt"
	"git.xx.network/elixxir/coupons/storage"
	"github.com/golang/protobuf/proto"
	jww "github.com/spf13/jwalterweatherman"
	"gitlab.com/elixxir/client/api"
	"gitlab.com/elixxir/client/interfaces/message"
	"gitlab.com/elixxir/client/interfaces/params"
	"strconv"
	"strings"
	"time"
)

type listener struct {
	delay time.Duration
	s     *storage.Storage
	c     *api.Client
}

var validResponse = "Thank you! You will receive %d in %s within 2 weeks!"
var noWalletFound = "Thank you! You will receive %d. The xx team does not have a wallet address for your purchase, please contact %s."
var reusedCode = "Your code %s has already been used. Your wallet %s will receive %d within 2 weeks."
var invalidCode = "That is not a valid code. Please send the code you received in your email for participating in the November Community Sale. Please contact %s if you need support."
var supportEmail = "support@xx.network"

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

	var strResponse string
	// Retrieve coupon code for trigger if it exists
	codeResp, uses, err := l.s.GetCouponCode(trigger)
	rl := strings.Split(codeResp, ",")
	num, err := strconv.Atoi(rl[0])
	if err != nil {
		jww.ERROR.Printf("Could not parse num coins: %+v", err)
		return
	}
	wallet := rl[1]
	if err != nil {
		jww.DEBUG.Printf("No coupon code for trigger %s: %+v", trigger, err)
		strResponse = fmt.Sprintf(invalidCode, supportEmail)
	} else {
		if uses < 1 {
			strResponse = fmt.Sprintf(reusedCode, trigger, wallet, num)
		} else if wallet == "" {
			strResponse = fmt.Sprintf(noWalletFound, num, supportEmail)
		} else {
			strResponse = fmt.Sprintf(validResponse, num, wallet)
		}
	}

	payload := &CMIXText{
		Text: strResponse,
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
		MessageType: message.Text,
	}

	// Send response message to sender over cmix
	rids, mid, t, err := l.c.SendE2E(resp, params.GetDefaultE2E())
	if err != nil {
		jww.ERROR.Printf("Failed to send message: %+v", err)
	}
	jww.INFO.Printf("Sent response %s [%+v] to %+v on rounds %+v [%+v]", strResponse, mid, item.Sender.String(), rids, t)
}

// Name returns a name, used for debugging
func (l *listener) Name() string {
	return "Coupons-bot-listener"
}
