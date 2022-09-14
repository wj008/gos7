package gos7

import (
	"bytes"
	"encoding/binary"
	"math"
)

type DataReader struct {
	buffer []byte
}

func NewDataReader(buffer []byte) *DataReader {
	return &DataReader{
		buffer: buffer,
	}
}

func (dr *DataReader) GetValue(pos int, value interface{}) error {
	buf := bytes.NewReader(dr.buffer[pos:])
	if err := binary.Read(buf, binary.BigEndian, value); err != nil {
		return err
	}
	return nil
}

func (dr *DataReader) GetUint8(pos int) (uint8, error) {
	var value uint8
	if err := dr.GetValue(pos, &value); err != nil {
		return 0, err
	}
	return value, nil
}

func (dr *DataReader) GetReal(pos int) (float32, error) {
	var value uint32
	if err := dr.GetValue(pos, &value); err != nil {
		return 0, err
	}
	float := math.Float32frombits(value)
	return float, nil
}

func (dr *DataReader) GetInt(pos int) (int16, error) {
	var value int16
	if err := dr.GetValue(pos, &value); err != nil {
		return 0, err
	}
	return value, nil
}

func (dr *DataReader) GetWord(pos int) (uint16, error) {
	var value uint16
	if err := dr.GetValue(pos, &value); err != nil {
		return 0, err
	}
	return value, nil
}

func (dr *DataReader) GetDInt(pos int) (int32, error) {
	var value int32
	if err := dr.GetValue(pos, &value); err != nil {
		return 0, err
	}
	return value, nil
}

func (dr *DataReader) GetDWord(pos int) (uint32, error) {
	var value uint32
	if err := dr.GetValue(pos, &value); err != nil {
		return 0, err
	}
	return value, nil
}

func (dr *DataReader) GetBool(pos1 int, pos2 uint) (bool, error) {
	var b byte
	if err := dr.GetValue(pos1, &b); err != nil {
		return false, err
	}
	return b&(1<<pos2) != 0, nil
}

func (dr *DataReader) Read(typ string, index1 int, index2 int) (value interface{}, err error) {
	switch typ {
	case "bool":
		value, err = dr.GetBool(index1, uint(index2))
		return
	case "byte":
		value, err = dr.GetUint8(index1)
		return
	case "word":
		value, err = dr.GetWord(index1)
		return
	case "int":
		value, err = dr.GetInt(index1)
		return
	case "dword":
		value, err = dr.GetDWord(index1)
		return
	case "dint":
		value, err = dr.GetDInt(index1)
		return
	case "real":
		value, err = dr.GetReal(index1)
		return
	default:
		value = nil
		return
	}
}
