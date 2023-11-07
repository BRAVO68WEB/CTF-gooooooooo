package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
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

			// Convert from hex to string
			s, err := hex.DecodeString(str)
			if err != nil {
				log.Fatal(err)
			}

			// Decode the string from base64 to string
			decoded, err := base64.StdEncoding.DecodeString(string(s))
			if err != nil {
				log.Fatal(err)
			}

			var finalFlag = string(decoded)

			// Print the flag to a file "flag.txt"
			file, err := os.Create("flag.txt")

			if err != nil {
				fmt.Println(err)
				return
			}

			l, err := file.WriteString(finalFlag)

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
