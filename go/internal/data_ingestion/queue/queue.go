package queue

import "log"

func HandleMsg(s string) {
	log.Printf("Ingesting data for %s", s)
}
