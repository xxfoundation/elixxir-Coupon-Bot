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
	trigger := "up up down down left right left right b a start"
	err = s.InsertCoupon(Coupon{
		Trigger: trigger,
		Code:    "helpimtrappedinacouponfactory",
		Uses:    1,
	})
	if err != nil {
		t.Errorf("Failed to insert coupon: %+v", err)
	}

	c, _, err := s.GetCouponCode("up up down down left right left right b a start")
	if err != nil {
		t.Errorf("Failed to get coupon code: %+v", err)
	}

	if c != "helpimtrappedinacouponfactory" {
		t.Errorf("Did not get expected coupon")
	}

	err = s.UseCode("zezima", trigger)
	if err != nil {
		t.Errorf("Failed to use code: %+v", err)
	}

	_, uses, err := s.GetCouponCode("up up down down left right left right b a start")
	if uses > 0 {
		t.Error("Should not have any uses left")
	}
}
