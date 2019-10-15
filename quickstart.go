package main

import (
	"strings"
	"os"
	"path"

	"./files"
	"./githelper"
	"./project"
)

const repoURL = "https://github.com/diareuse/teanity-app.git"
const branch = "one.zero"
const defaultPackage = "com.skoumal.teanity.app"

func main() {
	project := project.GetDefinition()
	git := githelper.By(repoURL, branch, project.Name)

	git.Clone()
	git.Init()

	pathDefault := strings.ReplaceAll(defaultPackage, ".", "/")
	pathPrefix := path.Join("src", "main", "java")
	pathBefore := path.Join(pathPrefix, pathDefault)
	pathAfter := path.Join(pathPrefix, project.GetFolder())

	files := files.By(project.Name)
	files.CleanUp()
	files.ForEach(func(file os.FileInfo) {
		if file.IsDir() {
			pathOld := path.Join(file.Name(), pathBefore)
			pathNew := path.Join(file.Name(), pathAfter)
			files.MoveDir(pathOld, pathNew)
		}
	})
	files.ReplaceInFiles(defaultPackage, project.PackageName)
}
