package main

import (
	"fmt"
)

func main() {
	for text := range fanin(gen("Jan", 5), gen("AM", 10)) {
		fmt.Println(text)
	}
	fmt.Println("I do not hear anything, I stop listening")
}

func gen(message string, number int) chan string {
	output := make(chan string)
	//Send message 10x to output channel, then close output channel
	go func() {
		for i := 0; i < number; i++ {
			output <- message
		}
		close(output)
	}()
	return output
}

func fanin(inputs ...chan string) chan string {
	output := make(chan string)
	done := make(chan bool)

	go func() {

		// Start for each input channel a goroutine to process received message
		for _, input := range inputs {

			//Receive and process message from input channel and place result on output channel
			//When input channel is closed send message that the goroutine has finished
			go func(c chan string) {
				for message := range c {
					output <- fmt.Sprintf("I'am listening to %s", message)
				}
				done <- true
			}(input)
		}

		//Wait until all goroutines are finished
		for i := 0; i < len(inputs); i++ {
			<-done
		}

		//Then close output channel
		close(output)
	}()

	return output
}
