package transport

import "time"

type null struct {
}

func (self *null) Read(m []byte) (int, error) {
	return 0, nil
}

func (self *null) Write(m []byte) (int, error) {
	return 0, nil
}

func (self *null) SetTimeout(d time.Duration) {
}

func (self *null) Open() error {
	return nil
}

func (self *null) Close() error {
	return nil
}

func (self *null) Flush() error {
	return nil
}

var TNullTransport = &null{}
