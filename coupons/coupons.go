////////////////////////////////////////////////////////////////////////////////
// Copyright © 2022 xx foundation                                             //
//                                                                            //
// Use of this source code is governed by a license that can be found in the  //
// LICENSE file.                                                              //
////////////////////////////////////////////////////////////////////////////////

package coupons

import (
	"git.xx.network/elixxir/coupons/storage"
	"gitlab.com/elixxir/client/v5/cmix"
)

// Impl struct wraps the listener for coupons
type Impl struct {
	*listener
}

// New initializes a listener with passed in storage and client
func New(s *storage.Storage, c *cmix.Client) *Impl {
	return &Impl{
		&listener{
			s: s,
			c: c,
		},
	}
}
