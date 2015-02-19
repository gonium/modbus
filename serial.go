// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.

package modbus

import (
	"log"
	"time"
)

const (
	// Default timeout
	serialTimeoutMillis = 5000
)

// SerialConfig is the common configuration for serial port.
type SerialConfig struct {
	// Device path (/dev/ttyS0)
	Address string
	// Baud rate (default 19200)
	BaudRate int
	// Data bits: 5, 6, 7 or 8 (default 8)
	DataBits int
	// Stop bits: 1 or 2 (default 1)
	StopBits int
	// Parity: N - None, E - Even, O - Odd (default E)
	// (The use of no parity requires 2 stop bits.)
	Parity string
}

type serialController interface {
	Connect() (err error)
	Close() (err error)

	isConnected() bool
	read(b []byte) (n int, err error)
	write(b []byte) (n int, err error)
}

// serialTransporter has configuration and I/O controller and must
// implicitly implement serialController interface.
type serialTransporter struct {
	// Serial port configuration.
	SerialConfig
	// serialPort is platform-dependent data structure for serial port.
	serialPort

	// Read timeout
	Timeout time.Duration
	Logger  *log.Logger
}

// Ensure serialTransporter also implements serialController interface.
var _ serialController = (*serialTransporter)(nil)

// SetSerialConfig allows a same configuration can be set for both RTU & ASCII handler.
func (serial *serialTransporter) SetSerialConfig(config SerialConfig) {
	serial.SerialConfig = config
}
