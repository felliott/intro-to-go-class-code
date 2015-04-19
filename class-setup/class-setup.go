package main

import (
	"errors"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"

	"github.com/Bowery/prompt"
)

func main() {
	msg("")
	checkForGit()

	gopath, set := getGopath()
	mkdir(path.Join(gopath, "bin"))
	mkdir(path.Join(gopath, "pkg"))

	repos := []string{
		"https://github.com/autarch/intro-to-go-class-exercises.git",
		"https://github.com/stretchr/testify.git",
	}
	for _, r := range repos {
		cloneRepo(gopath, r)
	}

	if !set {
		msg("")
		msg("  You should set your GOPATH environment variable to " + gopath)
	}

	msg("")
}

func checkForGit() {
	if _, err := exec.Command("git", "--version").Output(); err != nil {
		exitWithErr(errors.New("  This program requires you to have git executable in your path but none was found."))
	}
}

func getGopath() (string, bool) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		u, err := user.Current()
		if err != nil {
			exitWithErr(err)
		}

		gopath := path.Join(u.HomeDir, "gotest")
		if askOrDie("Use " + gopath + " as your GOPATH") {
			return gopath, false
		}

		setGopathExit()
	}

	if askOrDie("GOPATH is set to " + gopath + ". Use that path for class exercises") {
		return gopath, true
	}

	setGopathExit()

	return "", false
}

func askOrDie(q string) bool {
	ok, err := prompt.Ask("  " + q)
	if err != nil {
		exitWithErr(err)
	}
	return ok
}

func setGopathExit() {
	msg("\n  Set your GOPATH to the desired path and run this program again.\n")
	os.Exit(0)
}

func mkdir(p string) {
	if dirExists(p) {
		msg("  " + p + " already exists")
		return
	}

	err := os.MkdirAll(p, 0755)
	if err != nil {
		exitWithErr(err)
	}
	msg("  Made a new directory: " + p)
}

func chdir(p string) {
	if err := os.Chdir(p); err != nil {
		exitWithErr(err)
	}
}

func cloneRepo(gopath, r string) {
	schemeless := strings.TrimPrefix(r, "https://")
	elt := strings.Split(schemeless, "/")

	srcPath := path.Join(gopath, "src", elt[0], elt[1])
	clonePath := path.Join(srcPath, strings.TrimSuffix(elt[2], ".git"))
	if dirExists(path.Join(clonePath)) {
		msg("  " + clonePath + " already exists")
		return
	}

	mkdir(srcPath)
	chdir(srcPath)
	clone(r)
}

func clone(r string) {
	msg("  Cloning " + r)
	if err := exec.Command("git", "clone", "-q", r).Run(); err != nil {
		exitWithErr(err)
	}
}

func dirExists(d string) bool {
	s, err := os.Stat(d)
	if err == nil {
		if s.IsDir() {
			return true
		} else {
			exitWithErr(errors.New(d + " already exists and is not a directory!"))
		}
	} else if err != nil && !os.IsNotExist(err) {
		exitWithErr(err)
	}

	return false
}

func msg(s string) {
	os.Stdout.WriteString(s + "\n")
}

func exitWithErr(err error) {
	os.Stderr.WriteString("\n" + err.Error() + "\n\n")
	os.Exit(1)
}
