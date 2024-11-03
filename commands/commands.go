package commands

import (
	"encoding/binary"
)

const (
	COMMAND int32 = iota
	DATA
)

const (
	INVALID int32 = iota
	HELP
	PLAY
)

type Command struct {
	command int32
	argnum  int32
	args    []interface {
		Encode() []byte
	}
}

func (s Command) Encode() (ret []byte) {
	ret = make([]byte, 0, 12+(s.argnum*12))
	ret = binary.BigEndian.AppendUint32(ret, uint32(COMMAND))
	ret = binary.BigEndian.AppendUint32(ret, uint32(s.command))
	ret = binary.BigEndian.AppendUint32(ret, uint32(s.argnum))
	for _, v := range s.args {
		ret = append(ret, v.Encode()...)
	}
	return
}

type Node struct {
	freq uint32
	time uint32
}

func (s Node) Encode() (ret []byte) {
	ret = make([]byte, 0, 12)
	ret = binary.BigEndian.AppendUint32(ret, uint32(DATA))
	ret = binary.BigEndian.AppendUint32(ret, s.freq)
	ret = binary.BigEndian.AppendUint32(ret, s.time)
	return
}
