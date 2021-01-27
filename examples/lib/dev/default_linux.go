package dev

import (
	"github.com/markdrayton/ble"
	"github.com/markdrayton/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
