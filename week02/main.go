package main

import (
	"context"
	"fmt"
)

type printer interface {
	print(ctx context.Context, in <-chan struct{}, out chan<- struct{}, end chan<- struct{})
}

func main() {
	lc := make(chan struct{})
	nc := make(chan struct{})
	end := make(chan struct{})
	child, cancel := context.WithCancel(context.Background())
	var p printer
	p = &number{}
	go p.print(child, nc, lc, end)
	p = &letter{}
	go p.print(child, lc, nc, end)
	nc <- struct{}{}
	for i := 2; i > 0; i-- {
		<-end
	}
	cancel()
}

type number struct {
}

func (b *number) print(ctx context.Context, in <-chan struct{}, out chan<- struct{}, end chan<- struct{}) {
	limit := 30
	for i := 0; i <= limit; i += 2 {
		<-in
		fmt.Printf("%d", i)
		if i+1 <= limit {
			fmt.Printf("%d", i+1)
		}
		out <- struct{}{}
	}
	end <- struct{}{}
	for {
		select {
		case <-ctx.Done():
			return
		case <-in:
			out <- struct{}{}
		}
	}
}

type letter struct {
}

func (l *letter) print(ctx context.Context, in <-chan struct{}, out chan<- struct{}, end chan<- struct{}) {
	limit := 'Z'
	for i := 'A'; i <= limit; i += 2 {
		<-in
		fmt.Printf("%c", i)
		if i+1 <= limit {
			fmt.Printf("%c", i+1)
		}
		out <- struct{}{}
	}
	end <- struct{}{}
	for {
		select {
		case <-ctx.Done():
			return
		case <-in:
			out <- struct{}{}
		}
	}
}
