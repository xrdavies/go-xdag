package hash

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"log"
	"reflect"
	"unsafe"
	"xdag/sha256"
)

func init() {
	log.Println("xdagHash init")
}

// XdagHash double sha256
func XdagHash(data interface{}, size uintptr) (hashData []uint64, err error) {
	var buf bytes.Buffer
	ptr := reflect.ValueOf(data).Pointer()
	for i := uintptr(0); i < size; i++ {
		b := *(*byte)(unsafe.Pointer(ptr + i))
		err = buf.WriteByte(b)
		if err != nil {
			log.Println("[xdag_hash] data to bytes err: ", err)
			return
		}
	}

	hash256 := sha256.Sha256init(nil)
	sha256.Sha256update(hash256, buf.Bytes())
	tmpHashData := sha256.Sha256final(hash256)

	sha256.Sha256init(hash256)
	sha256.Sha256update(hash256, tmpHashData)
	finalHashData := sha256.Sha256final(hash256)

	hashData = make([]uint64, 4)

	var v = make([]byte, 0)
	var hashIdx, hashLen int
	for _, value := range finalHashData {
		hashTmp := value
		v = append(v, hashTmp)
		hashIdx++

		if hashIdx%8 == 0 {
			hashData[hashLen] = binary.LittleEndian.Uint64(v)
			hashLen++
			hashIdx = 0
			v = make([]byte, 0)
		}
	}
	return
}

// XdagHashInit ...
func XdagHashInit(hash256 interface{}) {
	sha256.Sha256init(hash256)
}

// XdagHashUpdate ...
func XdagHashUpdate(hash256 interface{}, data interface{}, size uintptr) {
	var buf bytes.Buffer
	ptr := reflect.ValueOf(data).Pointer()
	for i := uintptr(0); i < size; i++ {
		b := *(*byte)(unsafe.Pointer(ptr + i))
		err := buf.WriteByte(b)
		if err != nil {
			log.Println("[xdag_hash] data to bytes err: ", err)
			return
		}
	}
	sha256.Sha256update(hash256, buf.Bytes())
}

// XdagHashFinal ...
func XdagHashFinal(hash256 interface{}, data interface{}, size uintptr) (hashData []uint64, err error) {
	newHash256 := sha256.Sha256init(nil)

	if binaryMarshaler, ok := hash256.(encoding.BinaryMarshaler); ok {
		binaryData, e := binaryMarshaler.MarshalBinary()
		if e != nil {
			log.Println("[xdag_hash_final] marshal binary err: ", e)
			return nil, e
		}

		if binaryUnMarshaler, ok := newHash256.(encoding.BinaryUnmarshaler); ok {
			binaryUnMarshaler.UnmarshalBinary(binaryData)
		}
	}

	var buf bytes.Buffer
	ptr := reflect.ValueOf(data).Pointer()
	for i := uintptr(0); i < size; i++ {
		b := *(*byte)(unsafe.Pointer(ptr + i))
		err = buf.WriteByte(b)
		if err != nil {
			log.Println("[xdag_hash] data to bytes err: ", err)
			return
		}
	}

	sha256.Sha256update(newHash256, buf.Bytes())
	tmpHashData := sha256.Sha256final(newHash256)

	sha256.Sha256init(newHash256)
	sha256.Sha256update(newHash256, tmpHashData)
	finalHashData := sha256.Sha256final(newHash256)

	hashData = make([]uint64, 4)

	var v = make([]byte, 0)
	var hashIdx, hashLen int
	for _, value := range finalHashData {
		hashTmp := value
		v = append(v, hashTmp)
		hashIdx++

		if hashIdx%8 == 0 {
			hashData[hashLen] = binary.LittleEndian.Uint64(v)
			hashLen++
			hashIdx = 0
			v = make([]byte, 0)
		}
	}
	return
}
