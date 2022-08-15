package logger

import (
	"fmt"
	"io"
	"sync"
)

//
type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

//New ...
func New(w io.Writer, cap int) *Logger {
	l := Logger {
		ch: make(chan string, cap),
	}

	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		for v:= range l.ch{
			fmt.Fprint(w,v)
		}
	}()

	return &l
}


//
func (l *Logger) Close() {
	close(l.ch)
	l.wg.Wait()
}

//
func (l *Logger) Println(v string) {
	select {
	case l.ch <- v:
	default:
		fmt.Println("DROP")
	}
}