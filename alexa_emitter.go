package fsmalexa

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"

	"github.com/go-carrot/fsm-emitable"
)

// https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#ssml-supported
type AlexaEmitter struct {
	ResponseWriter io.Writer
	hasSpeech      bool
	speechBuffer   bytes.Buffer
}

func (a *AlexaEmitter) Emit(input interface{}) error {
	switch v := input.(type) {

	case string:
		a.speechBuffer.WriteString("<s>")
		a.speechBuffer.WriteString(v)
		a.speechBuffer.WriteString("</s>")
		a.hasSpeech = true
		return nil

	case emitable.Sleep:
		a.speechBuffer.WriteString("<break time=\"")
		a.speechBuffer.WriteString(strconv.Itoa(v.LengthMillis))
		a.speechBuffer.WriteString("ms\"/>")
		return nil

	case emitable.QuickReply:
		// Write message
		a.speechBuffer.WriteString("<p>")
		a.speechBuffer.WriteString(v.Message)
		a.speechBuffer.WriteString("</p>")

		// Options
		a.speechBuffer.WriteString("<p>You can ")
		for i, reply := range v.Replies {
			a.speechBuffer.WriteString(reply)
			if i+2 < len(v.Replies) && len(v.Replies) > 2 {
				a.speechBuffer.WriteString(", ")
			} else if i+1 < len(v.Replies) {
				if len(v.Replies) > 2 {
					a.speechBuffer.WriteString(", or ")
				} else {
					a.speechBuffer.WriteString(" or ")
				}
			}
		}
		a.speechBuffer.WriteString("</p>")
		return nil

	case emitable.Typing:
		// Intentionally do nothing
		return nil

	case emitable.Audio:
		// TODO
		return nil

	case emitable.Video:
		// TODO
		return nil

	case emitable.File:
		// TODO
		return nil

	case emitable.Image:
		// TODO
		return nil
	}
	return errors.New("AlexaEmitter cannot handle " + reflect.TypeOf(input).String())
}

func (a *AlexaEmitter) Flush() error {
	response := &ResponseBody{
		Version: "1.0",
		Response: &Response{
			ShouldEndSession: true,
		},
	}

	// Handle Speech
	if a.hasSpeech {
		ssml := "<speak>" + a.speechBuffer.String() + "</speak>"
		response.Response.OutputSpeech = &OutputSpeech{
			Type: "SSML",
			SSML: ssml,
		}
	}

	// Handle Response
	b, err := json.Marshal(response)
	if err != nil {
		return err
	}
	fmt.Fprint(a.ResponseWriter, string(b))
	return nil
}
