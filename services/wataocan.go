package services

import (
	"fmt"

	"github.com/jerrywang1981/watson/assistant"
)

type TaoCanInput struct {
	Age  int `form:"age"`
	Flow int `form:"flow"`
	Call int `form:"call"`
	Cost int `form:"cost"`
}

func processInput(input *TaoCanInput) {
	switch {
	case input.Age < 20:
		input.Age = 20
	}
}

func prepareSentence(input *TaoCanInput) string {
	return fmt.Sprintf(
		"The user is %d years old, and flow is %d, and call is %d, and average cost is %d",
		input.Age, input.Flow, input.Call, input.Cost)
}

func Send(input *TaoCanInput) *assistant.WAResult {
	processInput(input)
	msg := prepareSentence(input)
	return assistant.Bot.Send(msg)
}
