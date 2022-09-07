package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"sync"
)

func main() {
	n := flag.Int("num", 4, "0的位数")
	thread := flag.Int("thread", 10, "线程数")
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < *thread; i++ {
		go gen(*n)
	}
	wg.Wait()
}
func gen(num int) {
	for true {
		p, a := generate()
		s := "0x"
		for i := 0; i < num; i++ {
			s = s + "0"
		}
		//fmt.Println(s, a[0:num+2])
		if s == a[0:num+2] {
			fmt.Println(a, p)
		}
	}
}
func generate() (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyString := hexutil.Encode(privateKeyBytes)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return privateKeyString, address

}
