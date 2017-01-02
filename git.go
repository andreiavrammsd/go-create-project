package main

import "os/exec"

type Git struct {
	cmd string
}

func (git *Git) Exists() bool {
	out, _ := exec.Command(git.cmd).Output()
	return len(out) > 0
}

func (git *Git) Init() ([]byte, error) {
	return exec.Command(git.cmd, "init").Output()
}

func (git *Git) GetConfig(file, key string) string {
	value, _ := exec.Command(git.cmd, "config", file, key).Output()
	return string(value)
}

func (git *Git) Add(file string) ([]byte, error) {
	return exec.Command(git.cmd, "add", file).Output()
}

func (git *Git) Commit(message string) ([]byte, error) {
	return exec.Command(git.cmd, "commit", "-m", message).Output()
}

func NewGit() *Git {
	return &Git{
		cmd: "git",
	}
}
