package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

const AES_128_KEY_SIZE = 16
const AES_256_KEY_SIZE = 32

var password string
var salt string
var iteration int
var key_size int

var iv string
var key []byte

func encrypt_aes_with_pbkdf2(plain_text string) (bool, string) {
	password_str_to_byte_slice := []byte(password)
	salt_str_to_byte_slice := []byte(salt)

	// hash 종류 : sha1, sha256, sha512
	key = pbkdf2.Key(password_str_to_byte_slice, salt_str_to_byte_slice, iteration, key_size, sha1.New)

	//
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("aes.NewCipher fail : %s", err.Error())
		return false, ""
	}

	//
	block_size := block.BlockSize()
	input_str_to_btye_slice := []byte(plain_text)
	input_byte_slice_with_padding := pkcs5_padding(input_str_to_btye_slice, block_size)

	//
	for len(iv) < block_size {
		iv += "0"
	}
	iv_str_to_byte_slice := []byte(iv)
	block_mode := cipher.NewCBCEncrypter(block, iv_str_to_byte_slice)

	//
	encrypted_text := make([]byte, len(input_byte_slice_with_padding))
	block_mode.CryptBlocks(encrypted_text, input_byte_slice_with_padding)

	return true, base64.StdEncoding.EncodeToString(encrypted_text)
}

func decrypt_aes_with_pbkdf2(encrypted_text string) (bool, string) {

	decoded_input, err := base64.StdEncoding.DecodeString(encrypted_text)
	if err != nil {
		fmt.Println("base64.StdEncoding.DecodeString fail : %s", err.Error())
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("aes.NewCipher fail : %s", err.Error())
		return false, ""
	}

	iv_str_to_byte_slice := []byte(iv)
	block_mode := cipher.NewCBCDecrypter(block, iv_str_to_byte_slice)

	decrypted_text := make([]byte, len(decoded_input))
	block_mode.CryptBlocks(decrypted_text, decoded_input)
	decrypted_text = pkcs5_unpadding(decrypted_text)

	return true, string(decrypted_text)
}

func pkcs5_padding(data []byte, block_size int) []byte {
	padding := block_size - len(data)%block_size
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func pkcs5_unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func pkcs7_padding(data []byte, block_size int) ([]byte, error) {
	if block_size <= 0 {
		return nil, fmt.Errorf("invalid block size %d", block_size)
	}

	padlen := 1
	for ((len(data) + padlen) % block_size) != 0 {
		padlen = padlen + 1
	}

	pad := bytes.Repeat([]byte{byte(padlen)}, padlen)
	return append(data, pad...), nil
}

func pkcs7_unpadding(data []byte) ([]byte, error) {
	padlen := int(data[len(data)-1])
	pad := data[len(data)-padlen:]
	for i := 0; i < padlen; i++ {
		if pad[i] != byte(padlen) {
			return nil, fmt.Errorf("invalid padding")
		}
	}
	return data[:len(data)-padlen], nil
}

func main() {
	fmt.Println("Hello `Golang::Crypto` World!\n")

	// set up `material`
	password = "hellogolangcryptoworld"
	iteration = 1000
	salt = "mysalt"
	iv = "myiv"
	key_size = AES_128_KEY_SIZE // 단위는 byte

	fmt.Println("------------ set up `material` ------------")
	fmt.Println("password 		 	: ", password)
	fmt.Println("iteration			: ", iteration)
	fmt.Println("salt 			 	: ", salt)
	fmt.Println("key_size 		 	: ", key_size)
	fmt.Println("iv 			 	: ", iv)
	fmt.Println("-------------------------------------------")

	//
	plain_text := "chelseafandev"

	//
	result, encrypted_text := encrypt_aes_with_pbkdf2(plain_text)
	if result {
		fmt.Println("\nencrypt_aes128_with_pbkdf2 success :)")
		fmt.Println("  plain_text     	  : ", plain_text)
		fmt.Println("  encrypted(base64) 	  : ", encrypted_text)
		decoded, _ := base64.StdEncoding.DecodeString(encrypted_text)
		fmt.Println("  encrypted(hex)    	  : ", hex.EncodeToString(decoded))
	} else {
		fmt.Println("encrypt_aes128_with_pbkdf2 fail :(")
	}

	//
	result, decrypted_text := decrypt_aes_with_pbkdf2(encrypted_text)
	if result {
		fmt.Println("\ndecrypt_aes128_with_pbkdf2 success :)")
		fmt.Println("  plain_text     	  : ", plain_text)
		fmt.Println("  decrypted_text	  : ", decrypted_text)
	} else {
		fmt.Println("decrypt_aes128_with_pbkdf2 fail :(")
	}
}
