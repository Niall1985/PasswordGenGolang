package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//creating a random generator with a seed based on the corrent time
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//creating the character sets
	cl := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sl := "abcdefghijklmnopqrstuvwxyz"
	nos := "0123456789"
	sc := "!@#$%^&*()-_=+[]{}|;:,.<>?/"

	chars := cl + sl + nos + sc //combining all the character sets into 1 string

	//below is the finction which selects random characters
	get_Random_Character := func(s string) byte {
		return s[r.Intn(len(s))]
	}

	// function to generate a random password based on the length specified, in this case 10
	generatePassword := func(length int) string {
		password := make([]byte, length) //slice to store the password characters
		for i := range password {
			password[i] = get_Random_Character(chars)
		}
		return string(password) //return the generated password in string format
	}
	len_of_password := 10
	random_password := generatePassword(len_of_password)
	fmt.Printf("Generated Password: %s\n", random_password)
}
