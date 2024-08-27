package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
)

func pick_random(dirname string) string {
	files, err := os.ReadDir(dirname + "/.config/wallpaper/walls/")
	if err != nil {
		log.Fatal(err)
	}

	randomIndex := rand.Intn(len(files))
	filename := files[randomIndex].Name()
	return filename
}

func write_current(dirname string, filename string) {
	path := []byte(dirname + "/.config/wallpaper/walls/" + filename)
	outDir := dirname + "/.config/wallpaper/current"

	if err := os.WriteFile(outDir, path, 0666); err != nil {
		log.Fatal(err)
	}
}

func set_wallpaper(dirname string) {
	current := dirname + "/.config/wallpaper/current"

	content, err := os.ReadFile(current)
	if err != nil {
		fmt.Println("Err")
	}

	cmd := exec.Command("feh", "--bg-fill", string(content))
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	randomPtr := flag.Bool("random", false, "Random Wallpaper")
	setPtr := flag.Bool("set", false, "Set Wallpaper")
	verbosePtr := flag.Bool("verbose", false, "Verbose")
	givenFile := flag.String("filename", "", "filename")
	flag.Parse()

	if *randomPtr {
		filename := pick_random(dirname)
		write_current(dirname, filename)

		if *verbosePtr {
			fmt.Printf("Set '%s' as wallpaper.\n", filename)
		}
	} else if len(*givenFile) > 0 {
		if _, err := os.Stat(*givenFile); os.IsNotExist(err) {
			write_current(dirname, *givenFile)

			if *verbosePtr {
				fmt.Printf("Set '%s' as wallpaper.\n", *givenFile)
			}
		} else {
			fmt.Printf("File '%s' does not exist.\n", *givenFile)
		}

	}

	if *setPtr {
		set_wallpaper(dirname)

		if *verbosePtr {
			fmt.Println("Setting wallpaper ")
		}
	}
}
