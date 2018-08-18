package address

import (
	"fmt"
	"log"
	"testing"
)

func Test_XdagAddress2hash(t *testing.T) {
	var testAddress = []uint8("0ZMzmll399pSV45ELTp6TdY8SpMGt6u3")
	hash, err := XdagAddress2hash(testAddress)
	if err != nil {
		log.Fatalln("[xdag_address2hash]err is ", err)
	}
	fmt.Println(hash)
	hashStr := fmt.Sprintf("%016x%016x%016x%016x", hash[3], hash[2], hash[1], hash[0])
	if hashStr != "0000000000000000b7abb706934a3cd64d7a3a2d448e5752daf777599a3393d1" {
		log.Fatalln("[xdag_address2hash] Fail ", hashStr, " source: ", "0000000000000000b7abb706934a3cd64d7a3a2d448e5752daf777599a3393d1")
	}
}

func Test_XdagHash2address(t *testing.T) {
	var testHash = []uint64{15778211046238688209, 5582838654177269586, 13234873168827137238, 324235667}
	address := XdagHash2Address(testHash)
	if string(address) != "0ZMzmll399pSV45ELTp6TdY8SpMGt6u3" {
		log.Fatalln("[xdag_hash2address] Fail ", string(address), " source: 0ZMzmll399pSV45ELTp6TdY8SpMGt6u3")
	}
}
