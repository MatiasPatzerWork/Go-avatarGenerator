package avatar

import (
	"fmt"

	"github.com/MatiasPatzerWork/Go-avatarGenerator/avatar/encoder"
	"github.com/MatiasPatzerWork/Go-avatarGenerator/avatar/images"
)

// cryptoEncoder is  who can encode information.
type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation []byte, err error)
}

// imageGenerator is someone who can make images.
type imageGenerator interface {
	BuildAndSaveImage(encodedInformation []byte) error
}

// Information contains the fields for the information to encode.
type Information struct {
	// here go all information that yo want to encode
	Email string
}

// GeneratorOne is an struct that has an default encoder and image generator.
type GeneratorOne struct {
	encoder   cryptoEncoder
	generator imageGenerator
}

// DefaultAvatarGeneration returns a default avatar generator (*GeneratorOne) with a Sha1Encoder
// and a default avatar drawer. You must add a "blankAvatar.png" file inside your main folder.
func DefaultAvatarGeneration() *GeneratorOne {
	return &GeneratorOne{
		encoder:   encoder.NewSha1Encoder(),
		generator: images.NewDrawer(images.NewColorCreatorFromBytes()),
	}
}

// GenerateAndSaveAvatar takes the information of the type e.g: information.Information(Email:"your email goes here").
// It generates the identicon for the information and outputs a "newAvatar.jpg" file.
// You can change the type of information you want to encode by adding it as a field to the information struct and
// change the argument of s.encoder.EncodeInformation()
func (s *GeneratorOne) GenerateAndSaveAvatar(information Information) error {
	//here will be all logic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("an error ocurred: ", r)
		}
	}()
	encodedInformation, err := s.encoder.EncodeInformation(information.Email)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	err = s.generator.BuildAndSaveImage(encodedInformation)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return nil
}
