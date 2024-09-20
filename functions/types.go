package functions

type EncodeFunc func(string) (string, error)
type DecodeFunc func(string) (string, error)
type RandomFunc func(int) (string, error)
