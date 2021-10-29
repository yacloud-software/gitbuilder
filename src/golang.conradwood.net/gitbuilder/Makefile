.PHONY:	client server
all:	client server
client:
	cd client && go install gitbuilder-client.go
server:
	cd server && go install gitbuilder-server.go `ls -1 *.go|grep -v gitbuilder-server.go`
