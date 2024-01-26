.PHONY:	client server yabuild
all:	client server yabuild
client:
	cd client && go install ${LDFLAGS} gitbuilder-client.go
server:
	cd server && go install ${LDFLAGS} gitbuilder-server.go `ls -1 *.go|grep -v gitbuilder-server.go`
yabuild:
	cd yabuild && go install ${LDFLAGS} yabuild.go `ls -1 *.go|grep -v yabuild.go`
