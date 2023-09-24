package uc776revb

import (
	"github.com/d2r2/go-i2c"
	"time"
)

const (
	Ssd1306I2cAddress = 0x3c
	Ssd1306Bus        = 1
	COMMAND           = 0
	DATA              = 1
)

type Lcd struct {
	i2c *i2c.I2C
}

func NewLcd(i2c *i2c.I2C) (*Lcd, error) {
	this := &Lcd{i2c: i2c}
	initByteSeq := []byte{
		0xAE, // disable display
		0x40,
		0xB0,
		0xC8,
		0x81,
		0xff,
		0xa1,
		0xa6,
		0xa8,
		0x1f,
		0xd3,
		0x00,
		0xd5,
		0xf0,
		0xd9,
		0x22,
		0xda,
		0x02,
		0xdb,
		0x49,
		0x8d,
		0x14,
		0xaf,
	}
	for _, b := range initByteSeq {
		err := this.writeByte(b, COMMAND)
		if err != nil {
			return nil, err
		}
	}
	err := this.Clear()
	if err != nil {
		return nil, err
	}
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

func (this *Lcd) Clear() error {
	for i := 0; i < 5; i++ {
		bytes := []byte{
			0xb0, 0xb1, 0xb2, 0xb3, 0xb4,
			0x00,
			0x10,
		}
		for _, value := range bytes {
			this.writeByte(value, COMMAND)
		}
		for n := 0; i < 128; n++ {
			this.writeByte(0, DATA)
		}
	}

	return nil
}

//{
//unsigned char i,n;
//for(i=0;i<4;i++)
//{
//OLED_WR_Byte (0xb0+i,OLED_CMD);
//OLED_WR_Byte (0x00,OLED_CMD);
//OLED_WR_Byte (0x10,OLED_CMD);
//for(n=0;n<128;n++)OLED_WR_Byte(0,OLED_DATA);
//}
//}
