package storage

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
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
	m.coupons[trigger] = c
	return c.Code, uses, nil
}

func (m *MapImpl) InsertCoupon(c Coupon) error {
	m.coupons[c.Trigger] = &c
	return nil
}

func (m *MapImpl) CheckUser(id string) (string, error) {
	u, ok := m.users[id]
	if !ok {
		return "", gorm.ErrRecordNotFound
	}
	return u.Trigger, nil
}

func (m *MapImpl) UseCode(id, trigger string) error {
	m.users[id] = m.coupons[trigger]
	c := m.coupons[trigger]
	fmt.Println(m.coupons)
	c.Uses = c.Uses - 1
	return nil
}
