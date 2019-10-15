package files

import (
	"strings"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Move files and folder to new destination. In case "old" file is a directory, the function
//reflects the directory to the "new" location. It keeps the structure the same.
func Move(old, new *string) {
	stat, err := os.Stat(*old)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Println("Debug.Error", "Cannot read file ", err)
		return
	}
	if stat.IsDir() {
		if err := os.MkdirAll(*new, os.ModePerm); err != nil {
			log.Println("Debug.Error", "Cannot create directory", *new)
			return
		}
		dirs, err := ioutil.ReadDir(*old)
		if err != nil {
			log.Println("Debug.Error", "Cannot read directory ", err)
			return
		}
		for _, name := range dirs {
			oldPath := path.Join(*old, name.Name())
			newPath := path.Join(*new, name.Name())
			Move(&oldPath, &newPath)
		}
	} else {
		os.Rename(*old, *new)
	}
}

// Replace replaces "old" with "new" in "file" recursively
func Replace(file, old, new *string) {
	replace(file, old, new)
}

func replace(file, old, new *string) {
	stat, err := os.Stat(*file)
	if err != nil {
		log.Println("Debug.Error", "Cannot read file ", err)
		return
	}
	if stat.IsDir() {
		dirs, err := ioutil.ReadDir(*file)
		if err != nil {
			log.Println("Debug.Error", "Cannot read directory ", err)
			return
		}
		for _, name := range dirs {
			newPath := path.Join(*file, name.Name())
			replace(&newPath, old, new)
		}
	} else {
		read, err := ioutil.ReadFile(*file)
		if err != nil {
			log.Println("Debug.Error", "Cannot read file", err)
		}

		newContents := strings.ReplaceAll(string(read), *old, *new)

		err = ioutil.WriteFile(*file, []byte(newContents), 0)
		if err != nil {
			log.Println("Debug.Error", "Cannot write file", err)
		}
	}
}
