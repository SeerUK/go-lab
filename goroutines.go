package main

import (
	"fmt"
	"sync"
)

func getMessage(name string) string {
	return "Hello, " + name
}

func main() {
	wg := sync.WaitGroup{}
	names := []string{
		"Callum",
		"Elliot",
		"Ioannis",
	}

	messages := make(chan string, len(names))

	wg.Add(len(names))

	for i, _ := range names {
		name := names[i]
		go func() {
			messages <- getMessage(name)
			wg.Done()
		}()
	}

	wg.Wait()
	close(messages)

	for msg := range messages {
		fmt.Println(msg)
	}

	fmt.Println("Done!")
}
