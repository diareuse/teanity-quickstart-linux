package main

import (
	"flag"
	"os"
	"path"
	"strings"

	"./files"
	"./githelper"
	"./project"
)

const repoURL = "https://github.com/skoumalcz/teanity-app.git"
const defaultPackage = "com.skoumal.teanity.app"

func main() {
	branch := flag.String("branch", "1.2", "Describes branch used to initialize this project")
	flag.Parse()

	println("Starting \"New Project Wizard\" for Teanity v" + *branch)

	project := project.GetDefinition()
	git := githelper.By(repoURL, *branch, project.Name)

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
