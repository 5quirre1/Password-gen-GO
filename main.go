package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)


var (
	lowerCharSet          = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet          = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet string = "!@#$%^&*()-"
	numberCharSet         = "123567890"
	minSpecialChar        = 2
	minUpperChar          = 2
	minNumberChar         = 2
	passwordLength        = 10
)

func main() {
	totalCharLenWithoutLowerChar := minUpperChar + minSpecialChar + minNumberChar

	if totalCharLenWithoutLowerChar >= passwordLength {
		fmt.Println("give the length of the password plz")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("How many passwords do ya want? - ")
	scanner.Scan()

	numberOfPasswords, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Plz give a number g")
		os.Exit(1)
	}
	rand.Seed(time.Now().Unix())

	for i := 0; i < numberOfPasswords; i++ {
		password := generatePassword()
		fmt.Printf("Password %v is %v \n", i+1, password)
	}
	

}

func generatePassword() string {

	password := ""


	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password = password + string(specialCharSet[random])
	}

	for i := 0; i < minUpperChar; i++ {
		random := rand.Intn(len(upperCharSet))
		password = password + string(upperCharSet[random])
	}

	for i := 0; i < minNumberChar; i++ {
		random := rand.Intn(len(numberCharSet))
		password = password + string(numberCharSet[random])
	}

	totalCharLenWithoutLowerChar := minUpperChar + minSpecialChar + minNumberChar

	remainingCharLen := passwordLength - totalCharLenWithoutLowerChar


	for i := 0; i < remainingCharLen; i++ {
		random := rand.Intn(len(lowerCharSet))
		password = password + string(lowerCharSet[random])
	}



	passwordRune := []rune(password)
	rand.Shuffle(len(passwordRune), func(i, j int) {
		passwordRune[i], passwordRune[j] = passwordRune[j], passwordRune[i]
	})

	password = string(passwordRune)
	return password
}package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)


var (
	lowerCharSet          = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet          = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet string = "!@#$%^&*()-"
	numberCharSet         = "123567890"
	minSpecialChar        = 2
	minUpperChar          = 2
	minNumberChar         = 2
	passwordLength        = 10
)

func main() {
	totalCharLenWithoutLowerChar := minUpperChar + minSpecialChar + minNumberChar

	if totalCharLenWithoutLowerChar >= passwordLength {
		fmt.Println("give the length of the password plz")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("How many passwords do ya want? - ")
	scanner.Scan()

	numberOfPasswords, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Plz give a number g")
		os.Exit(1)
	}
	rand.Seed(time.Now().Unix())

	for i := 0; i < numberOfPasswords; i++ {
		password := generatePassword()
		fmt.Printf("Password %v is %v \n", i+1, password)
	}
	

}

func generatePassword() string {

	password := ""


	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password = password + string(specialCharSet[random])
	}

	for i := 0; i < minUpperChar; i++ {
		random := rand.Intn(len(upperCharSet))
		password = password + string(upperCharSet[random])
	}

	for i := 0; i < minNumberChar; i++ {
		random := rand.Intn(len(numberCharSet))
		password = password + string(numberCharSet[random])
	}

	totalCharLenWithoutLowerChar := minUpperChar + minSpecialChar + minNumberChar

	remainingCharLen := passwordLength - totalCharLenWithoutLowerChar


	for i := 0; i < remainingCharLen; i++ {
		random := rand.Intn(len(lowerCharSet))
		password = password + string(lowerCharSet[random])
	}



	passwordRune := []rune(password)
	rand.Shuffle(len(passwordRune), func(i, j int) {
		passwordRune[i], passwordRune[j] = passwordRune[j], passwordRune[i]
	})

	password = string(passwordRune)
	return password
}
