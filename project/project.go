package project

import (
	"path"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Definition holds values read from user
type Definition struct {
	Name, PackageName string
}

// GetDefinition returns user input read from the console
func GetDefinition() Definition {
	var project Definition

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("First we need to setup few things")

	fmt.Print("Enter project name: ")
	project.Name, _ = reader.ReadString('\n')

	fmt.Print("Enter desired package name: ")
	project.PackageName, _ = reader.ReadString('\n')

	project.Name = cleanWhitespace(project.Name)
	project.Name = sanitize(project.Name)

	project.PackageName = cleanWhitespace(project.PackageName)
	project.PackageName = strings.ToLower(project.PackageName)

	ensurePackageName(project.PackageName)

	return project
}

// GetFolder returns folder required for given packageName
func (d Definition) GetFolder() string {
	folders := strings.Split(d.PackageName, ".")
	return path.Join(folders...)
}

func cleanWhitespace(input string) string {
	return strings.ReplaceAll(input, "\n", "")
}

func sanitize(input string) string {
	input = strings.ReplaceAll(input, " ", "_")
	regex, _ := regexp.Compile("[^a-zA-Z0-9_]")
	input = regex.ReplaceAllString(input, "")
	return input
}

func ensurePackageName(input string) {
	regex, _ := regexp.Compile("[^a-z.]")
	if regex.MatchString(input) {
		panic("The package name is invalid. Fix this character: '" + regex.FindString(input) + "'")
	}
}
