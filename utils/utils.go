package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func GenerateRandomPassword() string {
	var pass string
	var passTypes = []string{
		"name.separator.word.number",
		"number.separator.name.separator.word",
		"name.separator.date",
		"name.date.word.number",
	}
	passType := rand.Intn(len(passTypes))

	parts := strings.Split(passTypes[passType], ".")

	for _, part := range parts {
		if part == "name" {
			name := GenerateName()
			if rand.Intn(2) == 1 {
				pass += strings.ToLower(name)
			} else {
				pass += name
			}
		}

		if part == "separator" && rand.Intn(2) == 1 {
			pass += RandomChar("-_.")
		}

		if part == "word" {
			pass += GenerateWord()
		}

		if part == "date" {
			pass += GenerateDate()
		}

		if part == "number" {
			pass += strconv.Itoa(rand.Intn(2000))
		}

		if part == "lastname" {
			lastname := GenerateLastname()
			if rand.Intn(2) == 1 {
				pass += strings.ToLower(lastname)
			} else {
				pass += lastname
			}
		}
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

func GenerateDate() string {
	day := strconv.Itoa(rand.Intn(29) + 1)
	if len(day) < 2 {
		day = "0" + day
	}
	month := strconv.Itoa(rand.Intn(11) + 1)
	if len(month) < 2 {
		month = "0" + month
	}
	year := strconv.Itoa(rand.Intn(2016-1970) + 1970)
	if rand.Intn(2) == 1 {
		year = year[2:]
	}

	return fmt.Sprintf("%v%v%v", day, month, year)
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
