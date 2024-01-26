package filetransfer

import (
	"fmt"
	"io"
	"path/filepath"
	//	pb "golang.conradwood.net/apis/gitbuilder"
	"golang.conradwood.net/go-easyops/utils"
	"os"
	"strings"
)

func NewSender(opaque interface{}, f func(opaque interface{}, filename string, fi os.FileInfo, data []byte) error) *Sender {
	if f == nil {
		panic("missing send function")
	}
	s := &Sender{
		opaque:  opaque,
		sendfkt: f,
	}
	return s
}

// send all files in dir to peer
func (s *Sender) SendFiles(sourcedir string) error {
	s.sourcedir = sourcedir
	return s.send()
}

// send some files in dir to peer
func (s *Sender) SendSomeFiles(sourcedir string, filenames []string) error {
	s.sourcedir = sourcedir
	s.filenames = filenames
	return s.send_some_files()
}

type Sender struct {
	opaque    interface{}
	sendfkt   func(opaque interface{}, filename string, fi os.FileInfo, data []byte) error
	sourcedir string
	filenames []string
	count     int
}

func (s *Sender) send() error {
	err := utils.DirWalk(s.sourcedir, func(root, rel string) error {
		if strings.HasSuffix(rel, "~") {
			return nil
		}
		filename := filepath.Base(rel)
		if strings.HasPrefix(filename, "#") {
			return nil
		}
		if strings.HasPrefix(rel, ".git/") {
			return nil
		}
		//		fmt.Printf("sending %s\n", rel)
		s.filenames = append(s.filenames, rel)
		return nil
	})
	if err != nil {
		return err
	}
	return s.send_some_files()
}
func (s *Sender) send_some_files() error {
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
	s.count = len(s.filenames)
	return nil
}
func (s *Sender) FilesSent() int {
	return s.count
}
func (s *Sender) stream_file(filename string) error {
	if s.sendfkt == nil {
		return fmt.Errorf("no send function set")
	}
	fileInfo, err := os.Stat(s.sourcedir + "/" + filename)
	if err != nil {
		return err
	}
	fd, err := os.Open(s.sourcedir + "/" + filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	buf := make([]byte, 2048)
	first := true // must send at least one packet, even if file is 0 bytes long
	for {
		n, err := fd.Read(buf)
		if first || n > 0 {
			err = s.sendfkt(s.opaque, filename, fileInfo, buf[:n])
			if err != nil {
				fmt.Printf("Send error: %s\n", utils.ErrorString(err))
				return err
			}
		}
		first = false
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	return nil
}
