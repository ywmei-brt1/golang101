package main

import (
	"errors"
	"fmt"
)

var (
	cnt = 3
	ch  = make(chan error, cnt)
)

func task(ch chan error, name string) {
	var err error
	fmt.Printf("doing task %s\n", name)
	ch <- err
}

func scheduler() error {
	defer close(ch)
	go task(ch, "eat breakfast")
	go task(ch, "eat bread")
	go task(ch, "drink")
	for {
		if e := <-ch; e != nil {
			return errors.New("task crashed")
		}
		cnt -= 1
		if cnt == 0 {
			return nil
		}
	}
}

func main() {
	if err := scheduler(); err != nil {
		fmt.Printf("failed to executed tasks due to %+v", err)
	} else {
		fmt.Println("successfully executed tasks")
	}
}
