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
	Command int32
	Args    []interface {
		Encode() []byte
	}
}

func (s Command) Encode() (ret []byte) {
	ret = make([]byte, 0, 12+(len(s.Args)*12))
	ret = binary.BigEndian.AppendUint32(ret, uint32(COMMAND))
	ret = binary.BigEndian.AppendUint32(ret, uint32(s.Command))
	ret = binary.BigEndian.AppendUint32(ret, uint32(len(s.Args)))
	for _, v := range s.Args {
		ret = append(ret, v.Encode()...)
	}
	return
}

type Node struct {
	Freq uint32
	Time uint32
}

func (s Node) Encode() (ret []byte) {
	ret = make([]byte, 0, 12)
	ret = binary.BigEndian.AppendUint32(ret, uint32(DATA))
	ret = binary.BigEndian.AppendUint32(ret, s.Freq)
	ret = binary.BigEndian.AppendUint32(ret, s.Time)
	return
}
