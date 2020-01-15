package record

import (
	"io"
)

func NewAppendWriterAtOffset(w io.ReadWriteSeeker, offset int64) (*Writer, error) {
	startingBlock := offset / blockSize
	buf := make([]byte, blockSize)

	if _, err := w.Seek(startingBlock*blockSize, io.SeekStart); err != nil {
		return nil, err
	}
	if _, err := w.Read(buf); err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	if _, err := w.Seek(startingBlock*blockSize, io.SeekStart); err != nil {
		return nil, err
	}

	flusher, ok := w.(flusher)
	if !ok {
		flusher = nil
	}

	nw := &Writer{
		w:           w,
		f:           flusher,
		j:           int(offset % blockSize),
		blockNumber: startingBlock,
	}

	for i := range buf {
		nw.buf[i] = buf[i]
	}

	return nw, nil

}
