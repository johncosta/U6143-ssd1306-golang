package uc776revb

import (
	"github.com/d2r2/go-i2c"
	"time"
)

type Lcd struct {
	i2c *i2c.I2C
}

func NewLcd(i2c *i2c.I2C) (*Lcd, error) {
	this := &Lcd{i2c: i2c}
	initByteSeq := []byte{
		0xAE, // disable display
	}
	for _, b := range initByteSeq {
		err := this.writeByte(b, 0)
		if err != nil {
			return nil, err
		}
	}
	//err := this.Clear()
	//if err != nil {
	//	return nil, err
	//}
	//err = this.Home()
	//if err != nil {
	//	return nil, err
	//}
	return this, nil
}

func (this *Lcd) writeByte(data byte, controlPins byte) error {
	err := this.writeDataWithStrobe(data&0xF0 | controlPins)
	if err != nil {
		return err
	}
	err = this.writeDataWithStrobe((data<<4)&0xF0 | controlPins)
	if err != nil {
		return err
	}
	return nil
}

func (this *Lcd) writeDataWithStrobe(data byte) error {
	seq := []rawData{
		{data, 0}, // send data
	}
	return this.writeRawDataSeq(seq)
}

type rawData struct {
	Data  byte
	Delay time.Duration
}

func (this *Lcd) writeRawDataSeq(seq []rawData) error {
	for _, item := range seq {
		_, err := this.i2c.WriteBytes([]byte{item.Data})
		if err != nil {
			return err
		}
		time.Sleep(item.Delay)
	}
	return nil
}
