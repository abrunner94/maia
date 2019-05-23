package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

const (
	configFile = "../config.json"
)

type Maia struct {
	Host     string   `json:"host"`
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
		//fmt.Println(m)
		ssh := createSSHClient(m)
		createConnection(*ssh)
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

func createSSHClient(m Maia) *Connection {
	// No password, fall back to .pem file
	var sshConfig ssh.ClientConfig

	if m.Password == "" && m.Key != "" {
		sshConfig = ssh.ClientConfig{
			User: m.User,
			Auth: []ssh.AuthMethod{
				publicKeyFile(m.Key),
			},
		}
	} else if m.Password != "" && m.Key == "" {
		sshConfig = ssh.ClientConfig{
			User: m.User,
			Auth: []ssh.AuthMethod{
				ssh.Password(m.Password),
			},
		}
	} else {
		sshConfig = ssh.ClientConfig{
			User: m.User,
		}
	}

	maiaConn := Connection{
		Config:    m,
		SSHClient: sshConfig,
	}
	fmt.Println("Created SSH client!")
	return &maiaConn
}

func createConnection(c Connection) {
	port := "22"
	addr := c.Config.Host + ":" + port
	config := c.SSHClient
	_, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		log.Fatalf("Could not create connection: %v", err)
	}

	/*sess, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Could not create new session: %v", err)
	}*/
	fmt.Println("Created connection!")
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
