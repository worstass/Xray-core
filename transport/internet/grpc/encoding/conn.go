package encoding

import (
	"context"
	"net"

	"github.com/xtls/xray-core/common/buf"
	"github.com/xtls/xray-core/common/net/cnc"
)

type HunkConn interface {
	Send(*MultiHunk) error
	Recv() (*MultiHunk, error)
	SendMsg(m interface{}) error
	RecvMsg(m interface{}) error
}

type StreamCloser interface {
	CloseSend() error
}

type HunkReaderWriter struct {
	hc     HunkConn
	cancel context.CancelFunc

	buf   [][]byte
	index int
}

func NewHunkReadWriter(hc HunkConn, cancel context.CancelFunc) *HunkReaderWriter {
	return &HunkReaderWriter{hc, cancel, nil, 0}
}

func NewHunkConn(hc HunkConn, cancel context.CancelFunc) net.Conn {
	wrc := NewHunkReadWriter(hc, cancel)
	return cnc.NewConnection(
		cnc.ConnectionInputMulti(wrc),
		cnc.ConnectionOutputMulti(wrc),
		cnc.ConnectionOnClose(wrc),
	)
}

func (h *HunkReaderWriter) forceFetch() error {
	hunk, err := h.hc.Recv()
	if err != nil {
		return newError("failed to fetch hunk from gRPC tunnel").Base(err)
	}

	h.buf = hunk.Data
	h.index = 0

	return nil
}

func (h *HunkReaderWriter) ReadMultiBuffer() (buf.MultiBuffer, error) {
	if err := h.forceFetch(); err != nil {
		return nil, err
	}

	var mb = make(buf.MultiBuffer, 0, len(h.buf))
	for _, b := range h.buf {
		if cap(b) >= buf.Size {
			h.index = len(h.buf)
			mb = append(mb, buf.NewExisted(b))
			h.index++
			continue
		}

		nb := buf.New()
		nb.Extend(int32(len(b)))
		copy(nb.Bytes(), b)

		mb = append(mb, nb)
	}
	return mb, nil
}

func (h *HunkReaderWriter) WriteMultiBuffer(mb buf.MultiBuffer) error {
	defer buf.ReleaseMulti(mb)
	hunk := &MultiHunk{Data: make([][]byte, len(mb))}
	for _, b := range mb {
		hunk.Data = append(hunk.Data, b.Bytes())
	}

	err := h.hc.Send(hunk)
	if err != nil {
		return err
	}
	return nil
}

func (h *HunkReaderWriter) Close() error {
	if h.cancel != nil {
		h.cancel()
	}
	if sc, match := h.hc.(StreamCloser); match {
		return sc.CloseSend()
	}

	return nil
}
