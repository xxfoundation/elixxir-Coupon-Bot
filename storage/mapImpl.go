package storage

import (
	"errors"
	"fmt"
)

func (m *MapImpl) GetCouponCode(trigger string) (string, int, error) {
	c, ok := m.coupons[trigger]
	if !ok {
		return "", 0, errors.New(fmt.Sprintf("No coupon for trigger %s", trigger))
	}
	uses := int(c.Uses)
	if c.Uses <= 0 {
		return "", uses, errors.New("No uses left for requested coupon")
	}
	c.Uses = c.Uses - 1
	m.coupons[trigger] = c
	return c.Code, uses, nil
}

func (m *MapImpl) InsertCoupon(c Coupon) error {
	m.coupons[c.Trigger] = c
	return nil
}
