package main

import (
	"os"
	"fmt"
	"strings"
	"os/exec"
	"path/filepath"
	"io/ioutil"
	"strconv"
	"time"
)

const repositoryUrl = "git@github.com:andreiavrammsd/go-create-project.git"
const templateTmpDir = "/tmp/go-create-project-template"
const templateTmpDirProject = templateTmpDir + "/template"
const licenseFilename = "LICENSE"
const licenseCopyrightYearPlaceholder = "<year>"
const licenseCopyrightNamePlaceholder = "<name>"

var git = NewGit()

func main() {
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
	
	if len(os.Args) < 2 {
		fmt.Println("Project name not specified")
		return
	}
	projectName := strings.Trim(os.Args[1], " ")
	if len(projectName) == 0 {
		fmt.Println("Empty project name specified")
		return
	}

	projectPath := filepath.Clean(goPath + "/src/" + projectName)
	if _, err := os.Stat(projectPath); err == nil {
		fmt.Printf("Directory already exists: %s\n", projectPath)
		return
	}

	GetTemplate(repositoryUrl, templateTmpDir)

	err = CopyDir(templateTmpDirProject, projectPath)
	if err != nil {
		fmt.Println(err)
	}

	GitSetup(projectPath)
	Cleanup(templateTmpDir)

	fmt.Printf("Created new project at: %s\n", projectPath)
}

func GetTemplate(repositoryUrl, templateTmpDir string) {
	exec.Command("git", "clone", repositoryUrl, templateTmpDir).Run()
}

func GitSetup(dst string) {
	os.Chdir(dst)

	git.Init()
	
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

	git.Add(".gitignore")
	git.Add("LICENSE")
	git.Commit("Initial commit.")
	//exec.Command("git", "add", ".gitignore", "LICENSE").Run()
	//exec.Command("git", "commit", "-m", "Initial commit.").Run()
}

func Cleanup (dir string) {
	os.RemoveAll(dir)
}
