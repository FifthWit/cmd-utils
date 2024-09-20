package main

import (
	"cmd-utils/functions"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <encode|decode|random> -<type> <input>")
		os.Exit(1)
	}

	action := os.Args[1]
	typeFlag := os.Args[2]
	input := os.Args[3]

	switch action {
	case "encode":
		if encoder, exists := functions.Encoders[typeFlag[1:]]; exists {
			result, err := encoder(input)
			if err != nil {
				fmt.Println("Error encoding:", err)
				os.Exit(1)
			}
			fmt.Println(result)
		} else {
			fmt.Println("Unknown encoding type:", typeFlag)
			os.Exit(1)
		}
	case "decode":
		if decoder, exists := functions.Decoders[typeFlag[1:]]; exists {
			result, err := decoder(input)
			if err != nil {
				fmt.Println("Error decoding:", err)
				os.Exit(1)
			}
			fmt.Println(result)
		} else {
			fmt.Println("Unknown decoding type:", typeFlag)
			os.Exit(1)
		}
	case "random":
		if typeFlag != "-length" {
			fmt.Println("Unknown flag:", typeFlag)
			os.Exit(1)
		}
		length, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid length:", input)
			os.Exit(1)
		}
		result := functions.RandomString(length)
		fmt.Println(result)
	default:
		fmt.Println("Unknown action:", action)
		os.Exit(1)
	}
}
