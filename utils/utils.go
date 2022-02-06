package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func GenerateRandomPassword(length int) string {
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_.-"
	chl := len(chars) - 1
	pass := []byte{}
	for i := 0; i < length; i++ {
		pass = append(pass, chars[rand.Intn(chl)])
	}
	return string(pass)
}

func GenerateRandomEmail() string {
	domains := []string{"gmail.com", "hotmail.com", "outlook.com", "yahoo.com"}
	domain := domains[rand.Intn(len(domains)-1)]
	email := GenerateWord()

	if rand.Intn(3) == 1 {
		separator := RandomChar("-_.")
		email += separator + GenerateWord()
	}

	if rand.Intn(3) == 1 {
		number := rand.Intn(2000)
		email += strconv.Itoa(number)
	}

	return fmt.Sprintf("%v@%v", email, domain)
}

func GenerateEmail(name, lastname string) string {
	domains := []string{"gmail.com", "hotmail.com", "outlook.com", "yahoo.com"}
	email := ""

	if rand.Intn(2) == 1 {
		email += strings.ToLower(name)
	} else {
		email += name
	}

	if rand.Intn(4) != 1 {
		email += RandomChar("-_.")
	}

	if rand.Intn(2) == 1 {
		email += strings.ToLower(lastname)
	} else {
		email += lastname
	}

	if rand.Intn(4) == 1 {
		number := rand.Intn(2000)
		email += strconv.Itoa(number)
	}

	email += "@" + domains[rand.Intn(len(domains)-1)]

	return email
}

func fileToList(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func SelectCommonPassword() string {
	passwords := fileToList("./wordlists/commonPasswords.txt")
	return passwords[rand.Intn(len(passwords)-1)]
}

func GenerateName() string {
	names := fileToList("./wordlists/realNames.txt")
	return names[rand.Intn(len(names)-1)]
}

func GenerateWord() string {
	words := fileToList("./wordlists/words.txt")
	return words[rand.Intn(len(words)-1)]
}

func GenerateLastname() string {
	lastnames := fileToList("./wordlists/realLastnames.txt")
	return lastnames[rand.Intn(len(lastnames)-1)]
}

func RandomChar(s string) string {
	return string(s[rand.Intn(len(s))])
}
