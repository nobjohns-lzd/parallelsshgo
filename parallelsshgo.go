package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"github.com/hypersleep/easyssh"
	"github.com/howeyc/gopass"
)

func main() {

	input := new(sshInput)
	input.getConfig()
	input.execute()
}

type sshInput struct {
	user     string
	password string
	servers  []string
	key      string
	port     string
	command  string
}

func (sshConfig *sshInput) getConfig() *sshInput {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Username: ")
	userRaw, _ := reader.ReadString('\n')
	sshConfig.user = strings.TrimSpace(userRaw)

	fmt.Printf("Password(optional): ")
	passwordByte, err := gopass.GetPasswd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sshConfig.password = string(passwordByte)

	fmt.Printf("Server names, separated by semi-colon: ")
	serversRaw, _ := reader.ReadString('\n')
	sshConfig.servers = strings.Split(serversRaw, ";")

	// fmt.Printf("SSH Key(location based on home directory. Ex: /.ssh/id_rsa): ")
	// keyRaw, _ := reader.ReadString('\n')
	// sshConfig.key = strings.TrimSpace(keyRaw)
	sshConfig.key = "/.ssh/id_rsa"

	fmt.Printf("Port: ")
	portRaw, _ := reader.ReadString('\n')
	sshConfig.port = strings.TrimSpace(portRaw)

	fmt.Printf("Command to execute: ")
	commandRaw, _ := reader.ReadString('\n')
	sshConfig.command = strings.TrimSpace(commandRaw)

	return sshConfig
}

func (sshConfig *sshInput) execute() {
	// TODO: This code need to replace by goroutine
	ssh := &easyssh.MakeConfig{
		User:   sshConfig.user,
		Key:    sshConfig.key,
		Port:   sshConfig.port,
	}
	for _, server := range sshConfig.servers {
		ssh.Server = strings.TrimSpace(server)
		response, err := ssh.Run(sshConfig.command)
		if err != nil {
			panic("Command execution failed: " + err.Error())
		} else {
			fmt.Println(server + " >\n" + "Command: " + sshConfig.command + "\n" + "Output:\n" + response)
		}
	}
}