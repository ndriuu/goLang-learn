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


