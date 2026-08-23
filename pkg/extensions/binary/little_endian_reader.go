package xbinary

import (
	"encoding/binary"
	"unsafe"
)

type LittleEndianReader struct {
	index int
	buf   []byte
}

func NewLittleEndianReader(buf []byte) *LittleEndianReader {
	return &LittleEndianReader{
		index: 0,
		buf:   buf,
	}
}

func (r *LittleEndianReader) ReadUint8() uint8 {
	v := r.buf[r.index]
	r.index++

	return v
}
func (r *LittleEndianReader) ReadUint8x4() [4]uint8 {
	result := [4]uint8{
		r.buf[r.index],
		r.buf[r.index+1],
		r.buf[r.index+2],
		r.buf[r.index+3],
	}
	r.index += 4

	return result
}

func (r *LittleEndianReader) ReadByte() byte {
	v := r.buf[r.index]
	r.index++

	return v
}

func (r *LittleEndianReader) Read(receiver []byte) {
	copy(receiver, r.buf[r.index:])
	r.index += len(receiver)
}

func (r *LittleEndianReader) ReadInt8() int8 {
	v := r.buf[r.index]
	r.index++

	return int8(v)
}
func (r *LittleEndianReader) ReadUint16() uint16 {
	v := binary.LittleEndian.Uint16(r.buf[r.index:])
	r.index += int(unsafe.Sizeof(v))

	return v
}
func (r *LittleEndianReader) ReadInt16() int16 {
	v := int16(binary.LittleEndian.Uint16(r.buf[r.index:]))
	r.index += int(unsafe.Sizeof(v))

	return v
}
func (r *LittleEndianReader) ReadUint32() uint32 {
	v := binary.LittleEndian.Uint32(r.buf[r.index:])
	r.index += int(unsafe.Sizeof(v))

	return v
}
func (r *LittleEndianReader) ReadFloat32() float32 {
	v := float32(binary.LittleEndian.Uint32(r.buf[r.index:]))
	r.index += int(unsafe.Sizeof(v))

	return v
}

func (r *LittleEndianReader) ReadUint16x4() [4]uint16 {
	return [4]uint16{
		r.ReadUint16(),
		r.ReadUint16(),
		r.ReadUint16(),
		r.ReadUint16(),
	}
}

func (r *LittleEndianReader) ReadFloat32x4() [4]float32 {
	return [4]float32{
		r.ReadFloat32(),
		r.ReadFloat32(),
		r.ReadFloat32(),
		r.ReadFloat32(),
	}
}

func (r *LittleEndianReader) ReadFloat64() float64 {
	v := float64(binary.LittleEndian.Uint64(r.buf[r.index:]))
	r.index += int(unsafe.Sizeof(v))

	return v
}
func (r *LittleEndianReader) ReadUint64() uint64 {
	v := binary.LittleEndian.Uint64(r.buf[r.index:])
	r.index += int(unsafe.Sizeof(v))

	return v
}
