package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dirname, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }

	files, err := ioutil.ReadDir(dirname + "/.config/wallpaper/walls/")
	if err != nil {
		log.Fatal(err)
	}

	random_index := rand.Intn(len(files))
	fmt.Println(files[random_index].Name())
}
