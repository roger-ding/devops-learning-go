package main

import (
  "fmt"
  "os"

  "github.com/roger-ding/devops-learning-go/ssh"
)

func main() {
  var (
    privateKey []byte
    publicKey  []byte
    err        error
  )
  
  if privateKey, publicKey, err = ssh.GenerateKeys(); err != nil {
    fmt.Printf("Error: %s\n", err)
    os.Exit(1)
  }
	
  if err = os.WriteFile("mykey.pem", privateKey, 0600); err != nil {
    fmt.Printf("Error: %s\n", err)
    os.Exit(1)
  }
  
  if err = os.WriteFile("mykey.pub", publicKey, 0644); err != nil {
    fmt.Printf("Error: %s\n", err)
    os.Exit(1)
  }
}
