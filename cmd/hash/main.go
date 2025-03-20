package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

const sol = "someText"

func getHashSha(text string) []byte {
	text += sol
	hash := sha256.New()
	hash.Write([]byte(text))
	return hash.Sum(nil)
}
func checkHash(text string, hash []byte) bool {
	textHash := getHashSha(text)
	return bytes.Equal(hash, textHash)
}

func getHashMd5(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}
func checkHashMd5(fileName string, hash []byte) (bool, error) {
	fileHash, err := getHashMd5(fileName)
	if err != nil {
		return false, err
	}
	return bytes.Equal(hash, fileHash), nil
}

func main() {

	fmt.Println(" * * * SHA256 * * *")
	psw := "qwerty"
	fmt.Printf("pasword : %s\n", psw)
	hash := getHashSha(psw)
	fmt.Printf("pasword hash: %x\n", hash)
	fmt.Print("enter password: ")
	var newPsw string
	fmt.Scanln(&newPsw)
	fmt.Printf("check pasword: %t\n", checkHash(newPsw, hash))

	fmt.Println(" * * * MD5 * * *")
	fileName := "1.txt"
	fmt.Printf("file name : %s\n", fileName)
	newHash, err := getHashMd5(fileName)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("file hash: %x\n", newHash)
	fmt.Println("change file and press Enter")
	fmt.Scanln()
	check, err := checkHashMd5(fileName, newHash)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("check file hash: %t\n", check)

}
