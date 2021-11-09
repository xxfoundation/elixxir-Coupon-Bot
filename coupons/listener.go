package coupons

import (
	"git.xx.network/elixxir/coupons/storage"
	"github.com/golang/protobuf/proto"
	jww "github.com/spf13/jwalterweatherman"
	"gitlab.com/elixxir/client/api"
	"gitlab.com/elixxir/client/interfaces/message"
	"gitlab.com/elixxir/client/interfaces/params"
)

type listener struct {
	s *storage.Storage
	c *api.Client
}

// Hear messages from users to the coupon bot & respond appropriately
func (l *listener) Hear(item message.Receive) {
	// Confirm that authenticated channels
	if !l.c.HasAuthenticatedChannel(item.Sender) {
		jww.ERROR.Printf("No authenticated channel exists to %+v", item.Sender)
	}

	// Parse the trigger
	in := &CMIXText{}
	trigger := ""
	err := proto.Unmarshal(item.Payload, in)
	if err != nil {
		jww.ERROR.Printf("Could not unmartial message from messenger: %+v", err)
	} else {
		trigger = in.Text
	}

	// Retrieve coupon code for trigger if it exists
	strResponse, err := l.s.GetCouponCode(trigger)
	if err != nil {
		jww.DEBUG.Printf("No coupon code for trigger %s: %+v", trigger, err)
		strResponse = "No coupon found for trigger " + trigger
	}

	payload := &CMIXText{
		Text: strResponse,
	}
	marshalled, err := proto.Marshal(payload)
	if err != nil {
		jww.ERROR.Printf("Failed to marshal payload: %+v", err)
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
