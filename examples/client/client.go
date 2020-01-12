// Copyright 2013 Alexandre Fiori
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Event Socket client that connects to FreeSWITCH to originate a new call.
package main

import (
	"fmt"
	"log"

	"github.com/axengine/go-eventsocket/eventsocket"
)

const dest = "sofia/internal/1000%127.0.0.1"
const dialplan = "&socket(localhost:9090 async)"

func main() {
	c, err := eventsocket.Dial("39.97.177.209:8021", "ClueCon")
	if err != nil {
		log.Fatal(err)
	}
	c.Send("events json ALL")

	ev, err := c.SendEvent(eventsocket.MSG{
		"profile":        "internal",
		"event-string":   "check-sync",
		"user":           "100000C1000",
		"host":           "172.17.132.168",
		"content-type":   "application/simple-message-summary",
		"content-length": "4",
	}, "NOTIFY", "dasd")

	if err != nil {
		panic(err)
	}
	ev.PrettyPrint()

	for {
		ev, err := c.ReadEvent()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\nNew event")
		ev.PrettyPrint()
		//if ev.Get("Answer-State") == "hangup" {
		//	break
		//}
	}
	c.Close()
}
