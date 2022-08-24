package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	mycrypto "GolangStudy/UnitTest/crypto/mycrypto"
)

func crypto_sample() {
	fmt.Println("Hello `Crypto` World!\n")

	password := "mycryptopassword"
	hash := mycrypto.HASH_FUNCTION_SHA_512
	salt := "mycryptosalt"
	iteration := 1000
	keysize := mycrypto.AES_128_KEY_SIZE
	iv := "mycryptoiv"

	mc := mycrypto.NewMyCrypto(password, hash, salt, iteration, keysize, iv)

	//
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type a string which you want to encrypt : ")
	scanner.Scan()
	plain_text := scanner.Text()

	//
	result, encrypted_text := mc.Encrypt_aes_with_pbkdf2(plain_text)
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
	result, decrypted_text := mc.Decrypt_aes_with_pbkdf2(encrypted_text)
	if result {
		fmt.Println("\ndecrypt_aes128_with_pbkdf2 success :)")
		fmt.Println("  plain_text     	  : ", plain_text)
		fmt.Println("  decrypted_text	  : ", decrypted_text)
	} else {
		fmt.Println("decrypt_aes128_with_pbkdf2 fail :(")
	}
}

func main() {
	crypto_sample()
}
