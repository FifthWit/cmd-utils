package main

import (
	"cmd-utils/functions"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "decode":
		decodeCmd := flag.NewFlagSet("decode", flag.ExitOnError)
		base64Flag := decodeCmd.String("base64", "", "Base64 string to decode")
		hexFlag := decodeCmd.String("hex", "", "Hex string to decode")
		urlFlag := decodeCmd.String("url", "", "URL string to decode")
		decodeCmd.Parse(os.Args[2:])

		if *base64Flag != "" {
			result, err := functions.Base64Decode(*base64Flag)
			if err != nil {
				fmt.Println("Error decoding base64:", err)
				os.Exit(1)
			}
			fmt.Println(result)
		} else if *hexFlag != "" {
			result, err := functions.HexDecode(*hexFlag)
			if err != nil {
				fmt.Println("Error decoding hex:", err)
				os.Exit(1)
			}
			fmt.Println(result)
		} else if *urlFlag != "" {
			result, err := functions.URLDecode(*urlFlag)
			if err != nil {
				fmt.Println("Error decoding URL:", err)
				os.Exit(1)
			}
			fmt.Println(result)
		} else {
			fmt.Println("Please provide a base64 string to decode using -base64 flag")
		}

	case "encode":
		encodeCmd := flag.NewFlagSet("encode", flag.ExitOnError)
		base64Flag := encodeCmd.String("base64", "", "Input string to encode")
		hexFlag := encodeCmd.String("hex", "", "Input string to encode")
		urlFlag := encodeCmd.String("url", "", "Input string to encode")
		encodeCmd.Parse(os.Args[2:])
		if *base64Flag != "" {
			result, err := functions.Base64Encode(*base64Flag)
			if err != nil {
				fmt.Println("Error encoding base64:", err)
				os.Exit(1)
			}
			fmt.Println(result)
		} else if *hexFlag != "" {
			result, error := functions.HexEncode(*hexFlag)
			if error != nil {
				fmt.Println("Error encoding hex:", error)
				os.Exit(1)
			}
			fmt.Println(result)
		} else if *urlFlag != "" {
			result, error := functions.URLEncode(*urlFlag)
			if error != nil {
				fmt.Println("Error encoding URL:", error)
				os.Exit(1)
			}
			fmt.Println(result)
		} else {
			fmt.Println("Please provide an input string to encode")
		}

	default:
		fmt.Println("No Command provided, please provide a valid command or use -help")
		os.Exit(1)
	}
}
