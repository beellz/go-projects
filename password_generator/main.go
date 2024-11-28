// to create a random password generator
// we need to have specific length
// numbers
// upper and lower case letter
// special character
// random output.

package main

import (
	"fmt"
	"math/rand"
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
	fmt.Print("How long the password length you want: ")
	fmt.Scanf("%d", &pass_length)
	fmt.Println("you selected length is ", pass_length)

	fmt.Print("password is :")
	fmt.Print(String(pass_length))

	// number

	// for i := 0; i < pass_length; i++ {
	// 	// fmt.Printf("GeeksforGeeks\n")
	// 	res1 := rand.Intn(10)
	// 	// fmt.Println("hello world")
	// 	fmt.Printf("%v", res1)
	// }

	// upper and lower case letters

}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// const charset = "abcdefghijklmnopqrstuvwxyz" +
// 	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// var seededRand *rand.Rand = rand.New(
// 	rand.NewSource(time.Now().UnixNano()))

// func StringWithCharset(length int, charset string) string {
// 	b := make([]byte, length)
// 	for i := range b {
// 		b[i] = charset[seededRand.Intn(len(charset))]
// 	}
// 	return string(b)
// }

// func String(length int) string {
// 	return StringWithCharset(length, charset)
// }

// func main() {
// 	fmt.Println(String(5))
// }
