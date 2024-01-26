package filetransfer

import (
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
	"os"
	"path/filepath"
)

type Transferrer struct {
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

func New(dir string) (*Transferrer, error) {
	t := &Transferrer{dir: dir}
	err := utils.RecreateSafely(dir)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Transferrer) Receive(data FileTransferPacket) error {
	fmt.Printf("Writing data to file \"%s\"\n", data.GetFilename())

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
func (t *Transferrer) Close() {
	if t.curfile != nil {
		t.curfile.Close()
		t.curfile = nil
	}
}
func (t *Transferrer) openFile(filename string) (*filehandle, error) {
	res := &filehandle{filename: filename}
	fname := t.dir + "/" + filename
	dir := filepath.Dir(fname)
	fmt.Printf("Creating dir \"%s\"\n", dir)
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
