package filetransfer

import (
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
	"os"
	"path/filepath"
)

type Receiver struct {
	dir     string
	curfile *filehandle
}
type filehandle struct {
	filename string
	fd       *os.File
}
type FileTransferPacket interface {
	GetFilename() string
	GetData() []byte
}

func NewReceiver(dir string) (*Receiver, error) {
	t := &Receiver{dir: dir}
	if !utils.FileExists(dir) {
		return nil, fmt.Errorf("cannot receive into non-existing directory \"%s\"", dir)
	}
	return t, nil
}

func (t *Receiver) Receive(data FileTransferPacket) error {
	//	fmt.Printf("Writing data to file \"%s\"\n", data.GetFilename())

	if t.curfile != nil && t.curfile.filename != data.GetFilename() {
		t.curfile.Close()
		t.curfile = nil
	}

	if t.curfile == nil {
		cf, err := t.openFile(data.GetFilename())
		if err != nil {
			return err
		}
		t.curfile = cf
	}

	err := t.curfile.Write(data.GetData())
	return err
}
func (t *Receiver) Close() {
	if t.curfile != nil {
		t.curfile.Close()
		t.curfile = nil
	}
}
func (t *Receiver) openFile(filename string) (*filehandle, error) {
	res := &filehandle{filename: filename}
	fname := t.dir + "/" + filename
	dir := filepath.Dir(fname)
	//fmt.Printf("Creating dir \"%s\"\n", dir)
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return nil, err
	}
	fd, err := os.Create(fname)
	if err != nil {
		return nil, err
	}
	res.fd = fd
	return res, nil
}
func (fh *filehandle) Write(data []byte) error {
	if fh.fd == nil {
		panic(fmt.Sprintf("attempt to write to closed fd for %s", fh.filename))
	}
	_, err := fh.fd.Write(data)
	return err
}
func (fh *filehandle) Close() {
	if fh.fd == nil {
		return
	}
	fh.fd.Close()
	fh.fd = nil
}
