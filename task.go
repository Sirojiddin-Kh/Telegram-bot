package main

import (

	"fmt"
)

type message struct {
	messages string
	priority string
	createdAt time
}

func main() {

	type messageSlice []message

	boxmessage := []messageSlice {

		[]message {
			message {
				messages : "Hello",
				priority : "Low",
			},
			message {
				messages : "Hi",
				priority : "medium",
			},
		},
		[]message {
			message {
				messages : "Hello",
				priority : "Low",
			},
			message {
				messages : "Hi",
				priority : "medium",
			},
		},

	}


	

	fmt.Println(boxmessage)
	
}