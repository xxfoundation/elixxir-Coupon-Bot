////////////////////////////////////////////////////////////////////////////////
// Copyright © 2022 xx foundation                                             //
//                                                                            //
// Use of this source code is governed by a license that can be found in the  //
// LICENSE file.                                                              //
////////////////////////////////////////////////////////////////////////////////

package storage

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (db *DatabaseImpl) GetCouponCode(trigger string) (string, int, error) {
	var c = &Coupon{}
	err := db.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&c, "trigger = ?", trigger).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", 0, err
	}

	return c.Code, int(c.Uses), nil
}

func (db *DatabaseImpl) InsertCoupon(c Coupon) error {
	return db.db.Create(c).Error
}

func (db *DatabaseImpl) CheckUser(id string) (string, error) {
	u := &User{}
	err := db.db.Where("id = ?", id).Take(&u).Error
	if err != nil {
		return "", err
	}
	return u.Trigger, nil
}

func (db *DatabaseImpl) UseCode(id, trigger string) error {
	return db.db.Transaction(func(tx *gorm.DB) error {
		u := &User{
			ID:      id,
			Trigger: trigger,
		}
		err := tx.Create(&u).Error
		if err != nil {
			return errors.WithMessage(err, "Failed to add user")
		}

		c := &Coupon{}
		err = tx.Model(&c).Where("trigger = ?", trigger).
			Update("uses", gorm.Expr("uses - ?", 1)).Error
		if err != nil {
			return errors.WithMessage(err, "Failed to use code")
		}
		return nil
	})
}
