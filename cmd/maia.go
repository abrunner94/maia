package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

const (
	configFile = "../config.json"
)

type Maia struct {
	Host     string   `json:"host"`
	Port     string   `json:"port"`
	User     string   `json:"user"`
	Password string   `json:"password"`
	Key      string   `json:"key"`
	Command  []string `json:"command"`
}

type Connection struct {
	Config    Maia
	SSHClient ssh.ClientConfig
}

func main() {
	config := readConfig()

	// Loop over the credentials for SSH connection
	for _, m := range config {
		conn := createClient(m)
		ssh := createConnection(*conn)
		cmds := m.Command

		// Create session and run commands
		createSession(ssh, cmds)
	}
}

// Read the config.json file and
func readConfig() []Maia {
	var ms []Maia

	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Could not read config.json: %v", err)
	}

	var tmpMaia []Maia
	err = json.Unmarshal([]byte(file), &tmpMaia)
	if err != nil {
		log.Fatalf("Could not unmarshal file: %v", err)
	}

	ms = append(ms, tmpMaia...)
	return ms
}

func createClient(m Maia) *Connection {
	// No password, fall back to .pem file
	var sshConfig ssh.ClientConfig

	if m.Password == "" && m.Key != "" {
		sshConfig = ssh.ClientConfig{
			User: m.User,
			Auth: []ssh.AuthMethod{
				publicKeyFile(m.Key),
			},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}
	} else if m.Password != "" && m.Key == "" {
		sshConfig = ssh.ClientConfig{
			User: m.User,
			Auth: []ssh.AuthMethod{
				ssh.Password(m.Password),
			},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}
	} else {
		sshConfig = ssh.ClientConfig{
			User: m.User,
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}
	}

	maiaConn := Connection{
		Config:    m,
		SSHClient: sshConfig,
	}
	fmt.Println("Created SSH client!")
	return &maiaConn
}

func createConnection(c Connection) *ssh.Client {
	port := c.Config.Port
	host := c.Config.Host
	addr := host + ":" + port
	config := c.SSHClient

	ssh, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		log.Fatalf("Could not create connection: %v", err)
	}
	fmt.Println("Created connection!")
	return ssh
}

func createSession(client *ssh.Client, cmds []string) {
	// Run all the commands as in the config file
	for _, cmd := range cmds {
		sess, err := client.NewSession()
		if err != nil {
			log.Fatalf("Could not create new session: %v", err)
		}
		var b bytes.Buffer
		sess.Stdout = &b
		if err := sess.Run(cmd); err != nil {
			log.Fatalf("Could not execute command: %v", err)
		}
		fmt.Println(b.String())
		defer sess.Close()
	}
}

func publicKeyFile(file string) ssh.AuthMethod {
	buff, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Could not read key file: %v", err)
	}

	key, err := ssh.ParsePrivateKey(buff)
	if err != nil {
		log.Fatalf("Could not parse key file: %v", err)
	}
	return ssh.PublicKeys(key)
}
