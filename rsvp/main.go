package main

import (
	"encoding/json"

	"github.com/jasonmoo/lambda_proc"
)

func main() {
	lambda_proc.Run(func(context *lambda_proc.Context, eventJSON json.RawMessage) (interface{}, error) {
		rsvp := NewRsvp()
		if err := json.Unmarshal(eventJSON, &rsvp); err != nil {
			return nil, err
		}
		return rsvp.Send()
	})
}
