package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/gitbuilder"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"os"
)

var (
	echoClient pb.GitBuilderClient
)

func main() {
	flag.Parse()

	echoClient = pb.GetGitBuilderClient()

	// a context with authentication
	ctx := authremote.Context()

	empty := &common.Void{}
	response, err := echoClient.Ping(ctx, empty)
	utils.Bail("Failed to ping server", err)
	fmt.Printf("Response to ping: %v\n", response)

	fmt.Printf("Done.\n")
	os.Exit(0)
}
