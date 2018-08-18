package hash

import (
	"log"
	"testing"
	"unsafe"
)

type T struct {
	Name [6]byte
	Age  uint8
}

type TT struct {
	Name []byte
	Age  uint8
}

func Test_XdagHash(t *testing.T) {
	testStruct := T{
		Name: [6]byte{'a', 'a', 'a', 'a', 'a', 'a'},
		Age:  20,
	}

	h, err := XdagHash(unsafe.Pointer(&testStruct), unsafe.Sizeof(testStruct))
	if err != nil {
		log.Fatalln("err: ", err)
		return
	}
	log.Printf("%016x%016x%016x%016x\n", h[3], h[2], h[1], h[0])
}
