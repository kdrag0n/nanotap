package main

/*
 * Original code Copyright (c) 2017 Olivier Poitrey
 *
 * MIT License
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

import (
	"context"
	"io"
	"sync"

	"code.cloudfoundry.org/go-diodes"
)

var bufPool = &sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 500)
	},
}

type DiodeAlerter func(missed int)

// DiodeWriter is a io.Writer wrapper that uses a diode to make Write lock-free,
// non-blocking and thread safe.
type DiodeWriter struct {
	writer io.Writer
	diode  *diodes.ManyToOne
	waiter *diodes.Waiter
	cancel context.CancelFunc
	done   chan struct{}
}

// NewDiodeWriter creates a writer wrapping w with a many-to-one diode in order to
// never block log producers and drop events if the writer can't keep up with
// the flow of data.
//
// Use a DiodeWriter when
//
//     wr := NewDiodeWriter(w, 1000, func(missed int) {
//         log.Printf("Dropped %d messages", missed)
//     })
//     log := zerolog.New(wr)
func NewDiodeWriter(w io.Writer, size int, f DiodeAlerter) DiodeWriter {
	ctx, cancel := context.WithCancel(context.Background())
	d := diodes.NewManyToOne(size, diodes.AlertFunc(f))
	dw := DiodeWriter{
		writer: w,
		diode:  d,
		waiter: diodes.NewWaiter(d,
			diodes.WithWaiterContext(ctx)),
		cancel: cancel,
		done:   make(chan struct{}),
	}
	go dw.poll()
	return dw
}

func (dw DiodeWriter) Write(p []byte) (n int, err error) {
	// p is pooled in zerolog so we can't hold it past this call, hence the
	// copy.
	p = append(bufPool.Get().([]byte), p...)
	dw.waiter.Set(diodes.GenericDataType(&p))
	return len(p), nil
}

// Close releases the diode poller and call Close on the wrapped writer if
// io.Closer is implemented.
func (dw DiodeWriter) Close() error {
	dw.cancel()
	<-dw.done
	if w, ok := dw.writer.(io.Closer); ok {
		return w.Close()
	}
	return nil
}

func (dw DiodeWriter) poll() {
	defer close(dw.done)
	for {
		d := dw.waiter.Next()
		if d == nil {
			return
		}
		p := *(*[]byte)(d)
		dw.writer.Write(p)
		bufPool.Put(p[:0])
	}
}
