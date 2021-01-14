package cloudinitfile

import "io/ioutil"

type MacFile struct {
	buf []byte
}

func NewMac(fpath string) (*MacFile, error) {
	buf, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return &MacFile{
		buf: buf,
	}, nil
}

func (m *MacFile) Contents() ([]byte, error) {
	return m.buf, nil
}

func (m *MacFile) ResetInstanceData() {
	// No-op, not needed on Mac.
}

func (m *MacFile) AddKipFile(content, path, permissions string) {
	// No-op, not supported on Mac.
}
