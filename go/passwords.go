package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func IsPasswordMatchWithHash(password string, hash string) bool {
	fmt.Println("Hash as bytes: ", []byte(hash))
	fmt.Println("Password as bytes: ", []byte(password))

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "myamazingpassword"
	hash, _ := HashPassword(password)

	fmt.Println("Password: ", password)
	fmt.Println("Hash: ", hash)

	match := IsPasswordMatchWithHash(password, hash)
	fmt.Println("Is Match: ", match)
}
