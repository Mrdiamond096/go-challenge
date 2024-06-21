package cipher

import (
	"io"
)

type Rot128Reader struct {
	reader io.Reader
}

func NewRot128Reader(r io.Reader) (*Rot128Reader, error) {
	return &Rot128Reader{reader: r}, nil
}

func (r *Rot128Reader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	if err != nil {
		return n, err
	}
	for i := range p[:n] {
		p[i] -= 128
	}
	return n, nil
}

type Rot128Writer struct {
	writer io.Writer
	buffer []byte // not thread-safe
}

func NewRot128Writer(w io.Writer) (*Rot128Writer, error) {
	return &Rot128Writer{
		writer: w,
		buffer: make([]byte, 4096),
	}, nil
}

func (w *Rot128Writer) Write(p []byte) (int, error) {
	n := copy(w.buffer, p)
	rot128(w.buffer[:n])
	return w.writer.Write(w.buffer[:n])
}

func rot128(buf []byte) {
	for idx := range buf {
		buf[idx] += 128
	}
}
