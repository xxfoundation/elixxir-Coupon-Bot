package storage

func (db *DatabaseImpl) GetCouponCode(trigger string) (string, error) {
	var c = &Coupon{}
	return c.Code, db.db.First(&c, "trigger = ?", trigger).Error
}

func (db *DatabaseImpl) InsertCoupon(c Coupon) error {
	return db.db.Create(c).Error
}
