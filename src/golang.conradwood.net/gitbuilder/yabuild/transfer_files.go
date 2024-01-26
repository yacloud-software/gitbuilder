package main

import (
	"fmt"
	pb "golang.conradwood.net/apis/gitbuilder"
	"golang.conradwood.net/go-easyops/utils"
	"os"
	"strings"
)

type transfer_destination interface {
	Send(m *pb.BuildLocalFiles) error
}

// send all files in dir to server
func transfer_files(td transfer_destination, sourcedir string) error {
	t := &sender{td: td, sourcedir: sourcedir}
	return t.send()
}

type sender struct {
	td        transfer_destination
	sourcedir string
	filenames []string
}

func (s *sender) send() error {
	err := utils.DirWalk(s.sourcedir, func(root, rel string) error {
		if strings.HasSuffix(rel, "~") {
			return nil
		}
		if strings.HasPrefix(rel, "#") {
			return nil
		}
		if strings.HasPrefix(rel, ".git/") {
			return nil
		}
		//		fmt.Printf("sending %s\n", rel)
		s.filenames = append(s.filenames, rel)
		return nil
	})
	p := utils.ProgressReporter{}
	p.SetTotal(uint64(len(s.filenames)))
	for _, f := range s.filenames {
		p.Add(1)
		p.Print()
		err := s.stream_file(f)
		if err != nil {
			return err
		}
	}
	return err
}

func (s *sender) stream_file(filename string) error {
	fd, err := os.Open(s.sourcedir + "/" + filename)
	if err != nil {
		return err
	}
	buf := make([]byte, 2048)
	for {
		n, err := fd.Read(buf)
		if n > 0 {
			br := &pb.BuildLocalFiles{
				FileTransfer: &pb.FileTransferPart{Filename: filename, Data: buf[:n]},
			}
			err = s.td.Send(br)
			if err != nil {
				fmt.Printf("Send error: %s\n", utils.ErrorString(err))
				return err
			}
		}
		if err != nil {
			break
		}
	}
	defer fd.Close()
	return nil
}
