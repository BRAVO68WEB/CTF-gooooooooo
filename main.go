package main

import (
	"crypto/aes"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	key := "thisis32bitlongpassphraseimusing"
	userPtr := flag.String("username", "user", "username of the decoder")
	passPtr := flag.String("password", "IamNoob", "password of the decoder")

	flag.Parse()

	// if userPtr and passPtr contain defaut values, print "hello-world"
	if *userPtr == "user" && *passPtr == "IamNoob" {
		fmt.Println("gooooooooo - The File decryptor !!")
		fmt.Println("   Flags -")
		fmt.Println("     -username       Pass your username here")
		fmt.Println("     -password       Pass your password here")
		fmt.Println("   Example command -")
		fmt.Println("     ./gooooooooo -username='' -password=''")
	} else if *userPtr == "admin" && *passPtr == "IamTheBestHacker" {
		fmt.Println("Access Granted, Welcome " + *userPtr + " the decoder")

		// Read the file "data.txt"
		f, err := os.Open("data.txt")

		if err != nil {
			fmt.Println(err)
			return
		}

		// Read the file "data.txt" as string and then convert from hex it to string
		// and print it
		for {
			var str string
			_, err := fmt.Fscan(f, &str)
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
				return
			}

			// Decrypt from AES to string
			depStr := DecryptAES([]byte(key), str)

			var finalFlag = string(depStr)
			// Print the flag to a file "flag.txt"
			file, err := os.Create("flag.txt")

			if err != nil {
				fmt.Println(err)
				return
			}

			l, err := file.WriteString("flag{" + finalFlag + "}")

			if err != nil {
				fmt.Println(err)
				file.Close()
				return
			}

			fmt.Println(l, "bytes written successfully")
		}
	} else {
		fmt.Println("Nice try, " + *userPtr + "  but you are not the decoder")
	}
}

func DecryptAES(key []byte, ct string) string {

	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	return string(pt[:])
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
