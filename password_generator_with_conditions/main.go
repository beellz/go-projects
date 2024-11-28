// to create a random password generator
// we need to have specific length
// numbers
// upper and lower case letter
// special character
// random output.

// if condition with output from cmd line

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "!@#$%^&*()" + "1234567890"

func String_charc(pass_length int, charset string) string {
	b := make([]byte, pass_length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return String_charc(length, charset)
}

func main() {

	var pass_length int

	fmt.Println("Welcome to password generator ")

	// get the password length from the args
    len_args := len(os.Args)

	if len_args > 1 {
		arg_length, _ := strconv.Atoi(os.Args[1])
		fmt.Println("you selected length is ", arg_length)
		fmt.Print("password is :")
		fmt.Print(String(arg_length))
	}else if len_args <= 1 {

		fmt.Print("How long the password length you want: ")
		fmt.Scanf("%d", &pass_length)
		fmt.Println("You selected length is ", pass_length)
		fmt.Print("password is :")
		fmt.Print(String(pass_length))
	
	}
	

	
}