package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	connect := &twiml.VoiceConnect{}

	conversation := &twiml.VoiceConversation{}
	conversation.ServiceInstanceSid = "ISxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	conversation.Record = "record-from-answer"

	connect.InnerElements = []twiml.Element{conversation}

	twimlResult, err := twiml.Voice([]twiml.Element{connect})
	if err == nil {
		fmt.Println(twimlResult)
	} else {
		fmt.Println(err)
	}
}
