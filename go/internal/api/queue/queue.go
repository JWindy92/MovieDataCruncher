package queue

import "log"

func HandleMsg(s string) {
	log.Printf("API Recieved a msg: %s", s)
}
