package storage

import (
	"errors"
	"fmt"
)

func (m *MapImpl) GetCouponCode(trigger string) (string, error) {
	c, ok := m.coupons[trigger]
	if !ok {
		return "", errors.New(fmt.Sprintf("No coupon for trigger %s", trigger))
	}
	if c.Uses <= 0 {
		return "", errors.New("No uses left for requested coupon")
	}
	c.Uses = c.Uses - 1
	m.coupons[trigger] = c
	return c.Code, nil
}

func (m *MapImpl) InsertCoupon(c Coupon) error {
	m.coupons[c.Trigger] = c
	return nil
}
