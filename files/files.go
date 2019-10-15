package files

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

// ProjectFile encapsulates root folder of the project
type ProjectFile struct {
	RootDir string
}

// By creates project file to call functions upon
func By(root string) ProjectFile {
	return ProjectFile{root}
}

// CleanUp removes directories which are not suitable or transitive to other projects
func (p ProjectFile) CleanUp() {
	ignoredDirs := []string{ /*nothing to ignore yet*/ }

	for _, dir := range ignoredDirs {
		err := os.RemoveAll(path.Join(p.RootDir, dir))
		if err != nil {
			log.Println("Debug.Error", err)
		}
	}
}

// FileIterable defines function used for iterating through files
type FileIterable func(file os.FileInfo)

// ForEach invokes for each file in directory of root ProjectFile
func (p ProjectFile) ForEach(onEach FileIterable) {
	dirs, err := ioutil.ReadDir(p.RootDir)
	if err != nil {
		log.Println("Debug.Error", err)
		return
	}
	for _, dir := range dirs {
		onEach(dir)
	}
}

// MoveDir renames folder to new name
func (p ProjectFile) MoveDir(old, new string) {
	oldPath := path.Join(p.RootDir, old)
	newPath := path.Join(p.RootDir, new)
	Move(&oldPath, &newPath)
}

// ReplaceInFiles removes old are puts new in its place
func (p ProjectFile) ReplaceInFiles(old, new string) {
	Replace(&p.RootDir, &old, &new)
}
