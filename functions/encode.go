package functions

import (
	"encoding/base64"
	"encoding/hex"
	"net/url"
)

var Encoders = map[string]EncodeFunc{
	"base64": Base64Encode,
	"hex":    HexEncode,
	"url":    URLEncode,
}

func Base64Encode(input string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	return encoded, nil
}

func HexEncode(input string) (string, error) {
	encoded := hex.EncodeToString([]byte(input))
	return encoded, nil
}

func URLEncode(input string) (string, error) {
	encoded := url.QueryEscape(input)
	return encoded, nil
}
