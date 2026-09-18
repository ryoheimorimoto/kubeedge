package fifo

import (
	"fmt"
	"sync"

	"k8s.io/klog/v2"

	"github.com/kubeedge/beehive/pkg/core/model"
	"github.com/kubeedge/kubeedge/pkg/viaduct/pkg/comm"
)

type MessageFifo struct {
	fifo      chan model.Message
	done      chan struct{}
	closeOnce sync.Once
}

// set the fifo capacity to MessageFiFoSizeMax
func NewMessageFifo() *MessageFifo {
	return &MessageFifo{
		fifo: make(chan model.Message, comm.MessageFiFoSizeMax),
		done: make(chan struct{}),
	}
}

// Put put the message into fifo
func (f *MessageFifo) Put(msg *model.Message) {
	select {
	case f.fifo <- *msg:
	default:
		// discard the old message
		<-f.fifo
		// push into fifo
		f.fifo <- *msg
		klog.Warning("too many message received, fifo overflow")
	}
}

// Get get message from fifo
// this api is blocked when the fifo is empty
func (f *MessageFifo) Get(msg *model.Message) error {
	// Drain what is already buffered before reporting the close, so a
	// teardown does not discard messages that arrived before it.
	select {
	case *msg = <-f.fifo:
		return nil
	default:
	}

	select {
	case *msg = <-f.fifo:
		return nil
	case <-f.done:
		// A message can be delivered just as the close happens, which makes
		// both cases ready; the select then picks at random, so check the
		// buffer once more before reporting the close.
		select {
		case *msg = <-f.fifo:
			return nil
		default:
			return fmt.Errorf("the fifo is broken")
		}
	}
}

// Close releases the callers blocked in Get. It deliberately leaves the
// message channel open: Put runs on the connection's read loop, so closing
// the channel here would panic with "send on closed channel" whenever a
// message arrives while the connection is being torn down.
func (f *MessageFifo) Close() {
	f.closeOnce.Do(func() {
		close(f.done)
	})
}
