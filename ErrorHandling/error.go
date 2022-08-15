package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"C:/Users/chukw/BatCave/Go/goPract/logger"

)

func init() {
	println("ffffff")

}

func main() {

	const grs = 10
	var d device;
	l := log.New(&d, "prefix", 0)

	for i := 0; i < grs; i++ {
		 go func (id int)  {
			for {
				l.Println(fmt.Sprintf("%d: log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)

	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		<- sigChan
		d.problem = !d.problem
	}

}

func webcall() error {
	return errors.New("errr")
}

type device struct {
		problem bool
	}

func (d *device) Write(p []byte) (n int, err error) {
	for d.problem {
		time.Sleep(time.Second)
	}

	fmt.Print(string(p))
	return len(p), nil

}


