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

// this interface is for the implementation of different types of save operations.
// ver si la uso
/*
type imageSaver interface {
	SaveImage() (err error)
}
*/
/*
// Service contains functionalities related to avatar generation.
type Service struct {
	encoder   cryptoEncoder
	generator imageGenerator
	//saver     imageSaver
}
*/

// Information contains information(?)
type Information struct {
	// here go all information that yo want to encode
	Email string
}

type GeneratorOne struct {
	// juampi te robo los nombres de las interfaces
	encoder   cryptoEncoder
	generator imageGenerator
}

// DefaultAvatarGeneration returns a default avatar generator with a Sha1Encoder
// and a default avatar drawer. You must add an "blankAvatar.png" file inside your main folder.
func DefaultAvatarGeneration() *GeneratorOne {
	return &GeneratorOne{
		encoder:   encoder.NewSha1Encoder(),
		generator: images.NewDrawer(images.NewColorCreatorFromBytes()),
	}
}

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

// CONSULTAR is
/*
Quiero saber como sería el uso del paquete. Entender como se usaría asi poder
ver por qué tienen que tener distintas funcionalidades.
2DA CONSULTA
Como construir un instanciador de los structs vacío y otro que permita variar los parámetros
Es decir, como hizo wini en la consulta, que yo pueda darle otro encoder y otra generación
de imagenes.

Suponiendo un funcionamiento del paquete:
gen1 := DefaultAvatarGeneration() → me genera un GeneratorOne

// s es la información a hashear
s := "matias.patzer@gmail.com"

encodedInformation, encodeErr := gen1.EncodeInformation(s)
if encodeErr != nil {
	fmt.Errorf(encodeErr)
}
// acá yo tendría mi string hasheado. Por como lo tengo implementado es un
// array de 20 bytes

errImg := gen1.BuildAndSaveImage(s)
// me genera y guarda la imagen
if errImg != nil {
	fmt.Errorf(errImg)
}

*/

// LA OTRA CONSULTA ESTÁ EN SERVICE.GO DE IMAGES
