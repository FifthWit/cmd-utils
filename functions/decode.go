// functions/decode.go
package functions

import (
	"encoding/base64"
	"encoding/hex"
	"net/url"
)

var Decoders = map[string]DecodeFunc{
	"base64": Base64Decode,
	"hex":    HexDecode,
	"url":    URLDecode,
}

func Base64Decode(input string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func HexDecode(input string) (string, error) {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func URLDecode(input string) (string, error) {
	decoded, err := url.QueryUnescape(input)
	if err != nil {
		return "", err
	}
	return decoded, nil
}
