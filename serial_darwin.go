// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.

package modbus

import (
	"fmt"
)

var errNotImplemented = fmt.Errorf("not implemented")

// serialPort is included in serialTransporter.
type serialPort struct {
}

// Connect opens serial port. Device must be set before calling this method.
func (mb *serialTransporter) Connect() (err error) {
	return errNotImplemented
}

func (mb *serialTransporter) Close() (err error) {
	return errNotImplemented
}

// isConnected returns true if serial port has been opened
func (mb *serialTransporter) isConnected() bool {
	return false
}

// read reads from serial port, blocks until data received or times out
func (mb *serialTransporter) read(b []byte) (n int, err error) {
	return 0, errNotImplemented
}

// write sends data to serial port
func (mb *serialTransporter) write(b []byte) (n int, err error) {
	return 0, errNotImplemented
}
