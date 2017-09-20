package gitroots

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Repos is a set of directories that are the top level directories
// of git projects.
type Repos map[string]struct{}

func walkDirectories(root string) (Repos, error) {
	repositories := make(Repos)
	walkFunc := func(path string, info os.FileInfo, err error) error {
		args := []string{"rev-parse", "--show-toplevel"}
		if !info.IsDir() {
			return nil
		}
		os.Chdir(path)
		out, err := exec.Command("git", args...).Output()
		if err != nil {
			return nil
		}
		repositories[strings.TrimSpace(string(out))] = struct{}{}
		return nil
	}
	err := filepath.Walk(root, walkFunc)
	return repositories, err
}

// Accept a string representing a filesystem path and return a
// Repos that represents every directory that is the root of a
// git project.
func Walk(root string) (Repos, error) {
	return walkDirectories(root)
}
