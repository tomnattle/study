package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

var iv []byte

func init() {

	blockSize := 16
	if b, e := aes.NewCipher([]byte("AES_CBC_PCK_ALG")); e == nil {
		blockSize = b.BlockSize()
	}
	iv = make([]byte, blockSize)
	for i := 0; i < blockSize; i++ {
		iv[i] = 0
	}

	fmt.Println("_init", iv)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesDecrypt(pass, key string) (_pass string, err error) {
	passBytes, err := base64.StdEncoding.DecodeString(pass)
	if err != nil {
		return
	}

	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	fmt.Println("...", iv)
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(passBytes))
	blockMode.CryptBlocks(origData, passBytes)

	return string(PKCS5UnPadding(origData)), nil
}

func TestEnc(t *testing.T) {
	pass := "Uy3QFmIPvBIdwdlAdewJ5dlgkGF3NPpC3I/ptd9aTgW4DzceDIN2poatCfYu+VQsmp5bQHeSQEIw3Td3vi2C1pkRks3JLjdYURjcc9rrk1YstzqXjMlL7pNjlDp6RzKyLW1Mhn1bwQEkXXW8PCB88dFFj1UzORVPAnu7jQEF8ENuljE+/mIBV/QTB0cRmZt22xWHJhFNVh4T3srFQ4C/GAq98Tf0xwsB9sUNb69nWJlBa4Ma6dyfSzWsrCXUmM9t6KjEHlSa44Jj4W7pNwfAqWERIteqioGpAcKUL6x5FHWPzhVIaEbEKxT8/28i07ueQTSZZ+48FkV+DoW/TAWCRs7RBY8ib1c2EKfzqiatZyU="
	key := "GWSGRykqLqavn2vZ9Hv0rg=="

	tpass, err := AesDecrypt(pass, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解密后:%s\n", tpass)
}

func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
