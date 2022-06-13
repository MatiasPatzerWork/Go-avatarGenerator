package encoder

import (
	"crypto/sha1"
)

// Sha1Encoder is the struct that has the Sha1Encoder implemented in the EncodeInformation function.
// It receives a string and returns a slice of bytes (len = 20) and an error. By default this is nil.
type Sha1Encoder struct{}

// We add the EncodeInformation method to the Sha1Encoder struct.
func (e *Sha1Encoder) EncodeInformation(strInformation string) (encodedInformation []byte, err error) {
	// https://gobyexample.com/sha1-hashes
	h := sha1.New()                 // initiate the sha1 encode func
	h.Write([]byte(strInformation)) // write expects bytes. We convert the string to bytes
	// the .Sum() allows to append the slice of bytes to an existing slice of bytes
	return h.Sum(nil), nil
}

// NewSha1Encoder is a constructor for the Sha1Encoder struct.
func NewSha1Encoder() *Sha1Encoder {
	return &Sha1Encoder{}
}
