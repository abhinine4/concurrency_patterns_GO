package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

// recieves messages from a channel
func (s *Server) StartAndListen() {

	//naming the for loop to break out of it
free:
	for {
		select {
		case msg := <-s.msgch:
			fmt.Printf("recieved message from: %s payload %s\n", msg.From, msg.Payload)
		case <-s.quitch:
			fmt.Println("Server is doing a gracefull shutdown")
			// logic for gracefull shutdown
			break free
		default:

		}
	}
	fmt.Println("Server is shutdown")
}

// recieves payload and sens message to channel
func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "SomeRandomGuy",
		Payload: payload,
	}
	msgch <- msg
}

func graceFullShutdown(quitch chan struct{}) {
	close(quitch)
}

func main() {
	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}
	go s.StartAndListen()

	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgch, "with some payload to server")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		graceFullShutdown(s.quitch)
	}()

	select {} // creates a deadlock

}

// ****************************************** sample concurrency explanation

// func main() {
// 	now := time.Now()

// 	userId := 10
// 	respch := make(chan string, 128) // buffered channel

// 	wg := &sync.WaitGroup{}

// 	go fetchUserData(userId, respch, wg)
// 	wg.Add(1)
// 	go fetchUserRecommendation(userId, respch, wg)
// 	wg.Add(1)
// 	go fetchUserLikes(userId, respch, wg)
// 	wg.Add(1)

// 	wg.Wait()

// 	close(respch)

// 	for resp := range respch {
// 		fmt.Println(resp)
// 	}

// 	fmt.Println(time.Since(now))
// }

// func fetchUserData(userId int, respch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(80 * time.Millisecond)
// 	respch <- "user data"
// 	wg.Done()
// }

// func fetchUserRecommendation(userId int, respch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(120 * time.Millisecond)
// 	respch <- "user recommendations"
// 	wg.Done()
// }

// func fetchUserLikes(userId int, respch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(50 * time.Millisecond)
// 	respch <- "user likes"
// 	wg.Done()
// }
