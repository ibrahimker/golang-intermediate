package main

import (
	"io"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	const SSH_ADDRESS = "0.0.0.0:22"
	const SSH_USERNAME = "user"
	const SSH_PASSWORD = "password"

	sshConfig := &ssh.ClientConfig{
		User:            SSH_USERNAME,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(SSH_PASSWORD),
		},
	}

	// dial ssh
	client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
	if client != nil {
		defer client.Close()
	}
	if err != nil {
		log.Fatal("Failed to dial. " + err.Error())
	}

	// create session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session. " + err.Error())
	}
	//session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// test session
	stdinBuf, _ := session.StdinPipe()
	if err := session.Shell(); err != nil {
		panic(err)
	}
	stdinBuf.Write([]byte("echo hello\n"))
	stdinBuf.Write([]byte("ls -l ~/\n"))
	//err = session.Run("ls -l ~/")
	//if err != nil {
	//	log.Fatal("Command execution error. " + err.Error())
	//}

	// test StdinPipe
	//var stdout, stderr bytes.Buffer
	//session.Stdout = &stdout
	//session.Stderr = &stderr
	//
	//stdin, err := session.StdinPipe()
	//if err != nil {
	//	log.Fatal("Error getting stdin pipe. " + err.Error())
	//}
	//
	//err = session.Start("/bin/bash")
	//if err != nil {
	//	log.Fatal("Error starting bash. " + err.Error())
	//}
	//
	//commands := []string{
	//	"cd /where/is/the/path",
	//	"cd src/myproject",
	//	"ls",
	//	"exit",
	//}
	//for _, cmd := range commands {
	//	if _, err = fmt.Fprintln(stdin, cmd); err != nil {
	//		log.Fatal(err)
	//	}
	//}
	//
	//err = session.Wait()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//outputErr := stderr.String()
	//fmt.Println("============== ERROR")
	//fmt.Println(strings.TrimSpace(outputErr))
	//
	//outputString := stdout.String()
	//fmt.Println("============== OUTPUT")
	//fmt.Println(strings.TrimSpace(outputString))

	// test sftp
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal("Failed create client sftp client. " + err.Error())
	}

	//err = session.Run("touch ~/test-file.txt")
	//if err != nil {
	//	log.Fatal("Command execution error. " + err.Error())
	//}
	//session.Close()
	fDestination, err := sftpClient.Create("/home/user/test-file.txt")
	if err != nil {
		log.Fatal("Failed to create destination file. " + err.Error())
	}

	fSource, err := os.Open("/home/ibam/Documents/code/golang-intermediate/session-7/test-file.txt")
	if err != nil {
		log.Fatal("Failed to read source file. " + err.Error())
	}

	_, err = io.Copy(fDestination, fSource)
	if err != nil {
		log.Fatal("Failed copy source file into destination file. " + err.Error())
	}

	log.Println("File copied.")

	// create new session
}
