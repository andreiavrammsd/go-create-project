# Create GO project

## Basic project template

### Creates directory with basic template

#### Dependencies

* GOPATH environment variable must be set
* Git

#### Setup

* Clone this repository to your GO src directory
* [install.sh](install.sh)

#### Usage

* gop project-name

#### Workflow

* Retrieves newest version from repository
* Creates project src directory with template
* Updates license: adds current year and the name you configured for git (git config --global user.name)
* If previous step was successful, creates initial commit with license and gitignore
