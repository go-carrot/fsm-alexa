package fsmalexa

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
)

// https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#ssml-supported
type AlexaEmitter struct {
	ResponseWriter io.Writer
	speechBuffer   bytes.Buffer
}

func (a *AlexaEmitter) Emit(input interface{}) error {
	switch v := input.(type) {
	case string:
		a.speechBuffer.WriteString("<s>")
		a.speechBuffer.WriteString(v)
		a.speechBuffer.WriteString("</s>")
		return nil
	}
	return errors.New("AlexaEmitter cannot handle " + reflect.TypeOf(input).String())
}

func (a *AlexaEmitter) Flush() error {
	// TODO, prepare response, properly
	response := &ResponseBody{
		Version: "1.0",
		Response: &Response{
			OutputSpeech: &OutputSpeech{
				Type: "PlainText",
				Text: "Hello world!  This is a sample.",
			},
			ShouldEndSession: true,
		},
	}

	// Handle Response
	b, err := json.Marshal(response)
	if err != nil {
		return err
	}
	fmt.Fprint(a.ResponseWriter, string(b))
	return nil
}
