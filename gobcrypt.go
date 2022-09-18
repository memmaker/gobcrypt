package main

import (
	"bufio"
	"golang.org/x/crypto/bcrypt"
	"os"
)

// GenHash generates a bcrypt hashed password string
func GenHash(password []byte) string {
	pw, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		return ""
	}

	return string(pw)
}

func main() {
	// get command line arguments
	args := os.Args[1:]
	if len(args) == 0 {
		stat, err := os.Stdin.Stat()
		if err != nil {
			printUsage()
		}
		if stat.Size() > 0 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				os.Stdout.WriteString(GenHash(scanner.Bytes()) + "\n")
			}
			return
		}
		printUsage()
	}

	plainPassword := []byte(args[0])
	hash := GenHash(plainPassword)
	println(hash)
}

func printUsage() {
	os.Stdout.WriteString("Usage: gobcrypt password\n")
	os.Stdout.WriteString("       gobcrypt < password\n")
	os.Exit(1)
}
