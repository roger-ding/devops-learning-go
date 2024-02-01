# devops-learning-go

## Run SSH scenario
First generate private and public keys
```
go run ssh/keygen/main.go
```
It should output a mykey.pem and mykey.pub, copy into another set for server and name it server.pem and server.pub

Using server and client code
1. Start up ssh server on original terminal
```
go run ssh/server/main.go
``` 
2. Open a second terminal tab and execute client 
```
go run ssh/client/main.go
```
3. Observe output from both terminal tabs

Using only server code and separate terminal invocation
1. Start up ssh server on original terminal
```
go run ssh/server/main.go
``` 
2. Open a second terminal tab and execute
```
ssh localhost -p 2022 -i mykey.pem 'whoami'
```
3. Observe output from both terminal tabs

## Run DNS resolver scenario
1. On original terminal
```
go run dns/main.go
```
2. Open up second terminal tab and perform some dns commands
```
$> dig @127.0.0.1 google.com
$> dig @127.0.0.1 nba.com   
```
3. Observe output in original terminal

### Testing DNS resolver with go test
```
go test -v ./dns/pkg
```