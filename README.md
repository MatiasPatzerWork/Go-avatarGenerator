# GOLANG - IDENTICON GENERATOR
![Identicon_generator-portadaGitHub](https://user-images.githubusercontent.com/99605067/173466250-a9ebc805-4ca4-47cc-8246-b1e377cbf49a.png)

Hi! This is a little proyect from the Golang back-end course. 

This package gives you the ability to encode information into a identicon with some algorithm and image manipulation.

This is an example of use:

```golang
package main

import (
	"fmt"

	gen "github.com/MatiasPatzerWork/Go-avatarGenerator/avatar"
	// give itaname to make it easier to use
)

func main() {
	// initialize the generator called avGenerator with the function DefaultAvatar Generation
	avGenerator := gen.DefaultAvatarGeneration()
	// this is the information to be encoded
	email := "matias.patzer@gmail.com"
	// this function will do everything for you;)
	avGenerator.GenerateAndSaveAvatar(gen.Information{Email: email})
	// just to keep track of the execution
	fmt.Println("done using the package")
}

// Remember to doa"go get"to the avatar@latest to get the latest version of the package
// go get github.com/MatiasPatzerWork/Go-avatarGenerator/avatar@latest
// NOTE:you MUST add an image as base for the avatar generator inside your main folder.
// It does not haveaspecified size but it MUST be called blankAvatar.jpg.
// The ouput image"newAvatar.jpg"is given to you inside your main folder.
// if there is no blankAvatar.png image it loggs an error but it doesn't panic.


```


NOTE: Please read the comments below the main function. They will help you to use the package. Also, every function and struct has it's correspondent documentation. 

Your folder structure should look something like this:

![image](https://user-images.githubusercontent.com/99605067/173467471-7083ba4c-824f-4780-8c93-0c54c34b22d0.png)


Have fun!

Matías Gabriel Patzer
