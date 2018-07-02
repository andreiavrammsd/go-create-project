package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const repositoryURL = "git@github.com:andreiavrammsd/go-create-project.git"
const templateTmpDir = "/tmp/go-create-project-template"
const templateTmpDirProject = templateTmpDir + "/template"
const licenseFilename = "LICENSE"
const licenseCopyrightYearPlaceholder = "<year>"
const licenseCopyrightNamePlaceholder = "<name>"

var git = NewGit()

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Project name not specified")
		return
	}

	projectName := strings.Trim(os.Args[1], " ")
	if len(projectName) == 0 {
		fmt.Println("Empty project name specified")
		return
	}

	goPath := os.Getenv("GOPATH")
	if len(goPath) == 0 {
		fmt.Println("GOPATH not defined")
		return
	}

	_, err := os.Stat(goPath)
	if err != nil {
		fmt.Printf("Directory defined at GOPATH does not exist: %s\n", goPath)
		return
	}

	if !git.Exists() {
		fmt.Println("Git not installed or not set up globally")
		return
	}

	projectPath := filepath.Clean(goPath + "/src/" + projectName)
	if _, err := os.Stat(projectPath); err == nil {
		fmt.Printf("Directory already exists: %s\n", projectPath)
		return
	}

	_, _ = git.Clone(repositoryURL, templateTmpDir)

	err = CopyDir(templateTmpDirProject, projectPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var out []byte

	out, err = gitSetup(projectPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))

	err = cleanup(templateTmpDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Created new project at: %s\n", projectPath)
}

func gitSetup(dst string) (out []byte, err error) {
	if err = os.Chdir(dst); err != nil {
		return
	}

	if out, err = git.Init(); err != nil {
		return
	}

	read, err := ioutil.ReadFile(licenseFilename)
	if err != nil {
		return
	}

	year := strconv.FormatInt(int64(time.Now().Year()), 10)
	updatedContent := strings.Replace(string(read), licenseCopyrightYearPlaceholder, year, 1)

	name := git.GetConfig("--global", "user.name")
	updatedContent = strings.Replace(updatedContent, licenseCopyrightNamePlaceholder, string(name), 1)

	err = ioutil.WriteFile(licenseFilename, []byte(updatedContent), 0)
	if err != nil {
		return
	}

	if out, err = git.Add(".gitignore"); err != nil {
		return
	}

	if out, err = git.Add("LICENSE"); err != nil {
		return
	}

	if out, err = git.Commit("Initial commit."); err != nil {
		return
	}

	return
}

func cleanup(dir string) error {
	return os.RemoveAll(dir)
}
