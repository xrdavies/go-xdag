package address

import (
	"encoding/binary"
	"errors"
	"log"
)

var bits2mime = []uint8("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

var mime2bits []uint8

// xdage init the address package
func init() {
	log.Println("xdagAddress init")

	mime2bits = make([]uint8, 0)
	for i := 0; i < 256; i++ {
		mime2bits = append(mime2bits, 0xFF)
	}

	for i := 0; i < 64; i++ {
		mime2bits[bits2mime[i]] = uint8(i)
	}
}

// XdagAddress2hash converts address to hash
func XdagAddress2hash(address []uint8) (hash []uint64, err error) {
	hash = make([]uint64, 4)

	var i, c, d, n, e uint
	var l, hashIdx, hashLen int
	var v = make([]byte, 0)

	for i = 0; i < 32; i++ {
		for {
			if l >= len(address) {
				return nil, errors.New("address is wrong")
			}
			c = uint(address[l])
			l++
			d = uint(mime2bits[c])

			if (d & 0xC0) <= 0 {
				break
			}
		}

		e <<= 6
		e |= d
		n += 6

		if n >= 8 {
			n -= 8
			hashTmp := byte(e >> n)
			v = append(v, hashTmp)
			hashIdx++
			if hashIdx%8 == 0 {
				hash[hashLen] = binary.LittleEndian.Uint64(v)
				hashLen++
				hashIdx = 0
				v = make([]byte, 0)
			}
		}
	}

	return
}

// XdagHash2Address converts hash to address
func XdagHash2Address(hash []uint64) (address []uint8) {
	var c, d, hashIdx uint
	var hashV = hash[0]
	address = make([]uint8, 0)

	for i := 0; i < 32; i++ {
		if hashV == 0 {
			hashIdx++
			hashV = hash[hashIdx]
		}
		if d < 6 {
			d += 8
			c <<= 8
			v := uint8(hashV)
			hashV >>= 8
			c |= uint(v)
		}
		d -= 6
		address = append(address, bits2mime[c>>d&0x3F])
	}

	return
}
