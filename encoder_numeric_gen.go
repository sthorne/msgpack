package msgpack

// Auto-generated by internal/cmd/genencoder/genencoder.go. DO NOT EDIT!

import (
	"math"

	"github.com/pkg/errors"
)

func (enl *encoderNL) EncodeInt(v int) error {
	if inNegativeFixNumRange(int64(v)) {
		return enl.encodeNegativeFixNum(int8(byte(0xff & v)))
	}

	if err := enl.dst.WriteByteUint64(Int64.Byte(), uint64(v)); err != nil {
		return errors.Wrap(err, `msgpack: failed to write Int64`)
	}
	return nil
}

func (enl *encoderNL) EncodeInt8(v int8) error {
	if inNegativeFixNumRange(int64(v)) {
		return enl.encodeNegativeFixNum(v)
	}

	if err := enl.dst.WriteByteUint8(Int8.Byte(), uint8(v)); err != nil {
		return errors.Wrap(err, `msgpack: failed to write Int8`)
	}
	return nil
}

func (enl *encoderNL) EncodeInt16(v int16) error {
	if inNegativeFixNumRange(int64(v)) {
		return enl.encodeNegativeFixNum(int8(byte(0xff & v)))
	}

	if err := enl.dst.WriteByteUint16(Int16.Byte(), uint16(v)); err != nil {
		return errors.Wrap(err, `msgpack: failed to write Int16`)
	}
	return nil
}

func (enl *encoderNL) EncodeInt32(v int32) error {
	if inNegativeFixNumRange(int64(v)) {
		return enl.encodeNegativeFixNum(int8(byte(0xff & v)))
	}

	if err := enl.dst.WriteByteUint32(Int32.Byte(), uint32(v)); err != nil {
		return errors.Wrap(err, `msgpack: failed to write Int32`)
	}
	return nil
}

func (enl *encoderNL) EncodeInt64(v int64) error {
	if inNegativeFixNumRange(int64(v)) {
		return enl.encodeNegativeFixNum(int8(byte(0xff & v)))
	}

	if err := enl.dst.WriteByteUint64(Int64.Byte(), uint64(v)); err != nil {
		return errors.Wrap(err, `msgpack: failed to write Int64`)
	}
	return nil
}

func (enl *encoderNL) EncodeUint(v uint) error {
	if inPositiveFixNumRange(int64(v)) {
		return enl.encodePositiveFixNum(uint8(0xff & v))
	}

	if err := enl.dst.WriteByteUint64(Uint64.Byte(), uint64(v)); err != nil {
		return errors.Wrap(err, `msgpack: failed to write Uint64`)
	}
	return nil
}

func (enl *encoderNL) EncodeUint8(v uint8) error {
	if inPositiveFixNumRange(int64(v)) {
		return enl.encodePositiveFixNum(uint8(0xff & v))
	}

	if err := enl.dst.WriteByteUint8(Uint8.Byte(), v); err != nil {
		return errors.Wrap(err, `msgpack: failed to write Uint8`)
	}
	return nil
}

func (enl *encoderNL) EncodeUint16(v uint16) error {
	if inPositiveFixNumRange(int64(v)) {
		return enl.encodePositiveFixNum(uint8(0xff & v))
	}

	if err := enl.dst.WriteByteUint16(Uint16.Byte(), v); err != nil {
		return errors.Wrap(err, `msgpack: failed to write Uint16`)
	}
	return nil
}

func (enl *encoderNL) EncodeUint32(v uint32) error {
	if inPositiveFixNumRange(int64(v)) {
		return enl.encodePositiveFixNum(uint8(0xff & v))
	}

	if err := enl.dst.WriteByteUint32(Uint32.Byte(), v); err != nil {
		return errors.Wrap(err, `msgpack: failed to write Uint32`)
	}
	return nil
}

func (enl *encoderNL) EncodeUint64(v uint64) error {
	if inPositiveFixNumRange(int64(v)) {
		return enl.encodePositiveFixNum(uint8(0xff & v))
	}

	if err := enl.dst.WriteByteUint64(Uint64.Byte(), v); err != nil {
		return errors.Wrap(err, `msgpack: failed to write Uint64`)
	}
	return nil
}

func (enl *encoderNL) EncodeFloat32(f float32) error {
	if err := enl.dst.WriteByteUint32(Float.Byte(), math.Float32bits(f)); err != nil {
		return errors.Wrap(err, `msgpack: failed to write Float`)
	}
	return nil
}

func (enl *encoderNL) EncodeFloat64(f float64) error {
	if err := enl.dst.WriteByteUint64(Double.Byte(), math.Float64bits(f)); err != nil {
		return errors.Wrap(err, `msgpack: failed to write Double`)
	}
	return nil
}
