package main

import (
	"fmt"
	"net"
	"time"

	"github.com/ugorji/go/codec"
)

// https://github.com/fluent/fluent-bit-docs/blob/1.9/development/msgpack-format.md

func getMessageBytes() []byte {
	var mh codec.MsgpackHandle

	// https://github.com/fluent/fluent-bit/blob/1fa0e94a09e4155f8a6d8a0efe36a5668cdc074e/plugins/in_forward/fw_prot.c#L417
	var v [3]interface{}
	v[0] = "iamatag"
	v[1] = uint64(time.Now().Unix())
	v[2] = map[string]interface{}{"key": "value"}

	var b []byte

	enc := codec.NewEncoderBytes(&b, &mh)
	err := enc.Encode(v)
	if err != nil {
		panic(err)
	}

	return b
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:24224")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {
		length, err := conn.Write(getMessageBytes())
		if err != nil {
			panic(err)
		}
		fmt.Println(length, "bytes sent on iteration", i)
		time.Sleep(2 * time.Second)
	}
}
