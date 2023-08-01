.PHONY:	client server
all:	client server
client:
	cd client && go install ${LDFLAGS} gitbuilder-client.go
server:
	cd server && go install ${LDFLAGS} gitbuilder-server.go `ls -1 *.go|grep -v gitbuilder-server.go`
