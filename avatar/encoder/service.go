package encoder

import (
	"crypto/sha1"
)

// Encode information should implemente the hash codification of the given string
// it receives an string and returns encoded string and an error.
type Sha1Encoder struct{}

// We add the EncodeInformation method to the Sha1Encoder struct
func (e *Sha1Encoder) EncodeInformation(strInformation string) (encodedInformation []byte, err error) {
	// https://gobyexample.com/sha1-hashes
	h := sha1.New()                 // initiate the sha1 encode func
	h.Write([]byte(strInformation)) // write expects bytes. We convert the string to bytes
	// the .Sum() allows to append the slice of bytes to an existing slice of bytes
	return h.Sum(nil), nil
}

func NewSha1Encoder() *Sha1Encoder {
	return &Sha1Encoder{}
}

// NOTAS PARA MÍ - Wini no leas esto
/*
Sha1Encoder es una estructura que implementa la interface cryptoEncoder. Esto es,
que una estructura tiene una función cuya firma es coincidente con la de la interface.
Esto es un "contrato" que permite aislar el codigo de la funcionalidad que tiene EncodeInformation
En la lógica principal, simplemente me interesa que el servicio tenga un encoder
que implemente una interfaz cryptoEncoder y que le paso un str y retorna un array de bytes
y un error.
*/
