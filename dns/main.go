package main

import (
  "fmt"
  "log"
  "net"

  "github.com/roger-ding/devops-learning-go/dns/pkg"
)

func main() {
  p, err := net.ListenPacket("udp", ":53")
  if err != nil {
    log.Fatal(err)
  }
  defer p.Close()

  for {
    buf := make([]byte, 512)
    n, addr, err := p.ReadFrom(buf)
    if err != nil {
      fmt.Printf("Connection error [%s]: %s\n", addr.String(), err)
      continue
    }
    go pkg.HandlePacket(p, addr, buf[:n])
  }
}