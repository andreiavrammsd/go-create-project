package main

import "os/exec"

// Git handles main Git commands
type Git struct {
	cmd string
}

// Exists checks if Git is installed
func (git *Git) Exists() bool {
	out, _ := exec.Command(git.cmd).Output()
	return len(out) > 0
}

// Init creates new repository
func (git *Git) Init() ([]byte, error) {
	return exec.Command(git.cmd, "init").Output()
}

// GetConfig retrieves Git configuration
func (git *Git) GetConfig(file, key string) string {
	value, _ := exec.Command(git.cmd, "config", file, key).Output()
	return string(value)
}

// Add starts tracking file into repository
func (git *Git) Add(file string) ([]byte, error) {
	return exec.Command(git.cmd, "add", file).Output()
}

// Commit the added files
func (git *Git) Commit(message string) ([]byte, error) {
	return exec.Command(git.cmd, "commit", "-m", message).Output()
}

// Clone repository from url to specified dir
func (git *Git) Clone(url, dir string) ([]byte, error) {
	return exec.Command(git.cmd, "clone", url, dir).Output()
}

// NewGit creates new Git management instance
func NewGit() *Git {
	return &Git{
		cmd: "git",
	}
}
