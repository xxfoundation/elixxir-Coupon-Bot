package storage

import (
	"errors"
	"gorm.io/gorm"
)

func (db *DatabaseImpl) GetCouponCode(trigger string) (string, error) {
	var c = &Coupon{}
	err := db.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&c, "trigger = ?", trigger).Error
		if err != nil {
			return err
		}
		if c.Uses <= 0 {
			return errors.New("requested code is out of uses")
		}
		err = tx.Model(&c).Where("trigger = ?", trigger).
			Update("uses", gorm.Expr("uses - ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return c.Code, nil
}

func (db *DatabaseImpl) InsertCoupon(c Coupon) error {
	return db.db.Create(c).Error
}
