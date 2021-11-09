package storage

import "testing"

func TestDatabase(t *testing.T) {
	s, err := NewStorage(Params{
		Username: "",
		Password: "",
		DBName:   "",
		Address:  "0.0.0.0",
		Port:     "5432",
	})
	if err != nil {
		t.Errorf("Failed to initialize storage: %+v", err)
	}
	err = s.InsertCoupon(Coupon{
		Trigger: "up up down down left right left right b a start",
		Code:    "helpimtrappedinacouponfactory",
		Uses:    1,
	})
	if err != nil {
		t.Errorf("Failed to insert coupon: %+v", err)
	}

	c, err := s.GetCouponCode("up up down down left right left right b a start")
	if err != nil {
		t.Errorf("Failed to get coupon code: %+v", err)
	}

	if c != "helpimtrappedinacouponfactory" {
		t.Errorf("Did not get expected coupon")
	}

	_, err = s.GetCouponCode("up up down down left right left right b a start")
	if err == nil {
		t.Error("Should have given error w/ no uses left")
	}
}
