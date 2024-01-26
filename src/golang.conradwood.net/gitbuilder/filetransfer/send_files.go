package filetransfer

import (
	"fmt"
	//	pb "golang.conradwood.net/apis/gitbuilder"
	"golang.conradwood.net/go-easyops/utils"
	"os"
	"strings"
)

func NewSender(opaque interface{}, f func(opaque interface{}, filename string, data []byte) error) *Sender {
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
	s.send_filenames = filenames
	return s.send()
}

type Sender struct {
	opaque         interface{}
	sendfkt        func(opaque interface{}, filename string, data []byte) error
	sourcedir      string
	send_filenames []string
	filenames      []string
}

func (s *Sender) send() error {
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

func (s *Sender) stream_file(filename string) error {
	if s.sendfkt == nil {
		return fmt.Errorf("no send function set")
	}
	fd, err := os.Open(s.sourcedir + "/" + filename)
	if err != nil {
		return err
	}
	buf := make([]byte, 2048)
	for {
		n, err := fd.Read(buf)
		if n > 0 {
			err = s.sendfkt(s.opaque, filename, buf[:n])
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
