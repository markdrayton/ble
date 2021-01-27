package dev

import (
	"github.com/markdrayton/ble"
	"github.com/markdrayton/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
