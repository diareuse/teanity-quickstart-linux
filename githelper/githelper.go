package githelper

import (
	"os"
	"os/exec"
	"path"
)

// Git defines structure used to hold required values
type Git struct {
	URL, Branch, Name string
}

// By creates new Git object
func By(url, branch, name string) Git {
	return Git{URL: url, Branch: branch, Name: name}
}

// Clone clones the repository and checkouts the correct branch
func (git Git) Clone() {
	exec.Command("git", "clone", "--recurse-submodules", git.URL, git.Name).Run()
	exec.Command("git", "checkout", "-b", git.Branch).Run()
}

// Init creates an empty repository
func (git Git) Init() {
	os.RemoveAll(path.Join(git.Name, ".git"))
	exec.Command("git", "init", git.Name)
}
