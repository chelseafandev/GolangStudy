package mycrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"hash"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const AES_128_KEY_SIZE = 16
const AES_256_KEY_SIZE = 32

const HASH_FUNCTION_SHA_1 = "sha1"
const HASH_FUNCTION_SHA_256 = "sha256"
const HASH_FUNCTION_SHA_512 = "sha512"

type MyCrypto struct {
	password  string
	hash      string
	salt      string
	iteration int
	keysize   int
	iv        string

	key []byte
}

func NewMyCrypto(password string, hash string, salt string, iteration int, keysize int, iv string) *MyCrypto {
	mc := MyCrypto{password: password, hash: hash, salt: salt, iteration: iteration, keysize: keysize, iv: iv}

	fmt.Println("------------ set up `material` ------------")
	fmt.Println("password 		 : ", mc.password)
	fmt.Println("hash 		 	 : ", mc.hash)
	fmt.Println("iteration		 : ", mc.iteration)
	fmt.Println("salt 			 : ", mc.salt)
	fmt.Println("keysize 		 : ", mc.keysize)
	fmt.Println("iv 			 : ", mc.iv)
	fmt.Println("-------------------------------------------")

	return &mc
}

func (mc *MyCrypto) Encrypt_aes_with_pbkdf2(plain_text string) (bool, string) {
	password_str_to_byte_slice := []byte(mc.password)
	salt_str_to_byte_slice := []byte(mc.salt)

	//
	var hash func() hash.Hash
	if strings.Compare(mc.hash, HASH_FUNCTION_SHA_1) == 0 {
		hash = sha1.New
	} else if strings.Compare(mc.hash, HASH_FUNCTION_SHA_256) == 0 {
		hash = sha256.New
	} else if strings.Compare(mc.hash, HASH_FUNCTION_SHA_512) == 0 {
		hash = sha512.New
	}

	//
	mc.key = pbkdf2.Key(password_str_to_byte_slice, salt_str_to_byte_slice, mc.iteration, mc.keysize, hash)

	//
	block, err := aes.NewCipher(mc.key)
	if err != nil {
		fmt.Println("aes.NewCipher fail : %s", err.Error())
		return false, ""
	}

	//
	block_size := block.BlockSize()
	input_str_to_btye_slice := []byte(plain_text)
	input_byte_slice_with_padding := mc.pkcs5_padding(input_str_to_btye_slice, block_size)

	//
	for len(mc.iv) < block_size {
		mc.iv += "0"
	}
	iv_str_to_byte_slice := []byte(mc.iv)
	block_mode := cipher.NewCBCEncrypter(block, iv_str_to_byte_slice)

	//
	encrypted_text := make([]byte, len(input_byte_slice_with_padding))
	block_mode.CryptBlocks(encrypted_text, input_byte_slice_with_padding)

	return true, base64.StdEncoding.EncodeToString(encrypted_text)
}

func (mc *MyCrypto) Decrypt_aes_with_pbkdf2(encrypted_text string) (bool, string) {

	decoded_input, err := base64.StdEncoding.DecodeString(encrypted_text)
	if err != nil {
		fmt.Println("base64.StdEncoding.DecodeString fail : %s", err.Error())
	}

	block, err := aes.NewCipher(mc.key)
	if err != nil {
		fmt.Println("aes.NewCipher fail : %s", err.Error())
		return false, ""
	}

	iv_str_to_byte_slice := []byte(mc.iv)
	block_mode := cipher.NewCBCDecrypter(block, iv_str_to_byte_slice)

	decrypted_text := make([]byte, len(decoded_input))
	block_mode.CryptBlocks(decrypted_text, decoded_input)
	decrypted_text = mc.pkcs5_unpadding(decrypted_text)

	return true, string(decrypted_text)
}

func (mc *MyCrypto) pkcs5_padding(data []byte, block_size int) []byte {
	padding := block_size - len(data)%block_size
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func (mc *MyCrypto) pkcs5_unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func (mc *MyCrypto) pkcs7_padding(data []byte, block_size int) ([]byte, error) {
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

func (mc *MyCrypto) pkcs7_unpadding(data []byte) ([]byte, error) {
	padlen := int(data[len(data)-1])
	pad := data[len(data)-padlen:]
	for i := 0; i < padlen; i++ {
		if pad[i] != byte(padlen) {
			return nil, fmt.Errorf("invalid padding")
		}
	}
	return data[:len(data)-padlen], nil
}
