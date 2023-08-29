# Golang-learn
Introduction of Go Language

## Description
This is not a project, but just my goal to learn more about Golang. Also, this content is for beginners like me, there are Golang basics.
I hope this content helps you guys to understand about goLang

## Installation
Follow these steps to install Go on your system:
#### Download Go:
Visit the official Go website at https://golang.org/dl/. Here, you will find the latest version of Go available for download. 
Choose the appropriate download link for your operating system (Windows, macOS, or Linux). Download the installer package.
#### Install Go on Windows:
- Double-click the downloaded installer package (.msi file).
- Follow the on-screen instructions. You can generally accept the default settings.
- Once the installation is complete, Go will be installed in the default location (usually C:\Go).
#### Install Go on macOS:
- Open the downloaded disk image (.pkg file).
- Follow the installation instructions provided in the installer window.
- Go will be installed to `/usr/local/go` by default.
#### Install Go on Linux:
- Open the downloaded disk image (.pkg file).
- Follow the installation instructions provided in the installer window.
- Go will be installed to `/usr/local/go` by default.
```sh
sudo tar -C /usr/local -xzf go1.x.x.linux-amd64.tar.gz
```
- Add Go's bin directory to your PATH environment variable. Edit your shell profile configuration file (e.g., ~/.bashrc, ~/.zshrc, ~/.profile) and add the following line:
```sh
export PATH=$PATH:/usr/local/go/bin
```
- Save the file and then run the following command to apply the changes to your current session:
```sh
source ~/.bashrc  # or source ~/.zshrc, depending on your shell
```
#### Verify Installation:
To verify that Go has been installed successfully, open a terminal and run the following command:
```sh
go version
```
You should see the installed Go version displayed in the output.

## Setting Up Your Workspace (GOPATH)
Go requires a specific directory structure called the GOPATH to manage your projects and dependencies. 
Starting from Go 1.11, you can work outside the GOPATH, but it's still commonly used.
1. Create Your Workspace:
- Choose a directory where you want to keep your Go projects and set it as your workspace.
- Create three subdirectories inside your workspace: bin, pkg, and src.
2. Set Environment Variables (Optional):
  
  You might want to set the following environment variables in your shell profile configuration file:
```sh
export GOPATH=/path/to/your/workspace
export PATH=$PATH:$GOPATH/bin
```
This will allow you to run Go executables globally and organize your workspace effectively.

## Get Started
Now that Go is installed, you're ready to start writing and running Go programs. 
You can use any text editor or integrated development environment (IDE) to write Go code.

To create and run a simple Go program, follow these steps:

1. Create a new directory inside your src directory (within your workspace) for your project.

2. Inside your project directory, create a new file with a .go extension (e.g., main.go).

3. Write your Go code in the file. For example, you can write a simple "Hello, World!" program:
```sh
package main

import "fmt"

func main() {
   fmt.Println("Hello, World!")
}

```
4. Open a terminal, navigate to your project directory, and run the following command to build and run your program:
```sh
go run main.go

```
Congratulations! You've successfully installed Go and written and executed your first Go program.

## Resources
- Official Go Website: https://golang.org/
- Go Documentation: https://golang.org/doc/
- Go Tour (Interactive Tutorial): https://tour.golang.org/welcome/1
Remember that this guide provides a basic overview of Go installation and setup.
As you delve deeper into Go programming, you'll explore more advanced topics and best practices.

Happy coding with Go! 🚀


