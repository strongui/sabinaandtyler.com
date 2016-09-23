package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Event string

const (
	ceremonyAndReception Event = "CeremonyAndReception"
	receptionOnly        Event = "ReceptionOnly"
)

func (event Event) FriendlyName() string {
	if event == ceremonyAndReception {
		return "Ceremony and Reception"
	}
	return "Reception Only"
}

func (event *Event) MarshalJSON() (b []byte, err error) {
	b = []byte(fmt.Sprintf("\"%s\"", *event))
	return
}

func (event *Event) UnmarshalJSON(b []byte) (err error) {

	str, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	e := Event(str)
	switch e {
	case ceremonyAndReception, receptionOnly:
		*event = e
		return
	}

	msg := fmt.Sprintf("invalid event type: '%s'", e)
	err = errors.New(msg)
	return
}
