package main

import (
	"encoding/json"
	"testing"
)

func TestRsvpMarshalJSONSuccess(t *testing.T) {

	for i, this := range []struct {
		input          Rsvp
		expectedOutput string
	}{
		{Rsvp{"Name1", ceremonyAndReception, true, 1, "comment 1"},
			`{"name":"Name1","event":"CeremonyAndReception","attending":true,"count":1,"comments":"comment 1"}`},

		{Rsvp{"Name2", receptionOnly, false, 2, "comment 2"},
			`{"name":"Name2","event":"ReceptionOnly","attending":false,"count":2,"comments":"comment 2"}`},

		{Rsvp{Name: "Name3", Attending: true, Count: 3, Comments: "comment 3"},
			`{"name":"Name3","attending":true,"count":3,"comments":"comment 3"}`},

		{Rsvp{Event: ceremonyAndReception, Attending: false, Count: 4, Comments: "comment 4"},
			`{"event":"CeremonyAndReception","attending":false,"count":4,"comments":"comment 4"}`},

		{Rsvp{},
			`{"attending":false,"count":0}`},
	} {

		b, err := json.Marshal(this.input)

		if err != nil {
			t.Errorf("[%d] unexpected error: %v", i, err)
		} else {
			actualOutput := string(b)
			if this.expectedOutput != actualOutput {
				t.Errorf("[%d] unexpected json result: '%s'.", i, actualOutput)
			}
		}
	}
}

// NOTE: Uncomment to actually test an email...
// func TestRsvpSend(t *testing.T) {

// 	rsvp := NewRsvp()
// 	rsvp.Name = "Nate Morse"
// 	rsvp.Event = ceremonyAndReception
// 	rsvp.Attending = true
// 	rsvp.Count = 2
// 	rsvp.Comments = "I love you!"

// 	r, err := rsvp.Send()

// 	if err != nil {
// 		t.Errorf("unexpexted error: %v", err)
// 	}

// 	if !r.Success {
// 		t.Error("send rsvp failed")
// 	}
// }

func TestRsvpMessageBody(t *testing.T) {

	for i, this := range []struct {
		input          Rsvp
		expectedOutput string
	}{
		{Rsvp{"Name1", ceremonyAndReception, true, 1, "comment 1"},
			`Yay! Molly and Nate! You have an RSVP to your wedding!

Name: Name1
Attending: YES
Event: Ceremony and Reception
Number of Guests: 1
Comments: comment 1`},

		{Rsvp{"Name2", ceremonyAndReception, false, 2, "comment 2"},
			`Oh no! Someone can't make it to your wedding!

Name: Name2
Attending: NO
Event: Ceremony and Reception
Number of Guests: 2
Comments: comment 2`},

		{Rsvp{"Name3", receptionOnly, true, 3, "comment 3"},
			`Looks like someone is ready to party!

Name: Name3
Attending: YES
Event: Reception Only
Number of Guests: 3
Comments: comment 3`},

		{Rsvp{"Name4", receptionOnly, false, 4, "comment 4"},
			`Oh no! Someone can't make it to your wedding!

Name: Name4
Attending: NO
Event: Reception Only
Number of Guests: 4
Comments: comment 4`},
	} {
		actualOutput := this.input.MessageBody()
		if this.expectedOutput != actualOutput {
			t.Errorf("[%d] unexpected output: '%v'.", i, actualOutput)
		}
	}
}

func TestRsvpSubject(t *testing.T) {

	for i, this := range []struct {
		input          Rsvp
		expectedOutput string
	}{
		{Rsvp{Attending: true, Event: ceremonyAndReception},
			"Wedding RSVP - Ceremony & Reception"},

		{Rsvp{Attending: false, Event: ceremonyAndReception},
			"Wedding RSVP - Regrets"},

		{Rsvp{Attending: true, Event: receptionOnly},
			"Wedding RSVP - Reception Only"},

		{Rsvp{Attending: false, Event: receptionOnly},
			"Wedding RSVP - Regrets"},
	} {
		actualOutput := this.input.Subject()
		if this.expectedOutput != actualOutput {
			t.Errorf("[%d] unexpected output: '%v'.", i, actualOutput)
		}
	}
}

func TestRsvpUnmarshalJSONSuccess(t *testing.T) {

	for i, this := range []struct {
		input          string
		expectedOutput Rsvp
	}{
		{`{"name":"Name1","event":"CeremonyAndReception","attending":true,"count":1,"comments":"comment 1"}`,
			Rsvp{"Name1", ceremonyAndReception, true, 1, "comment 1"}},

		{`{"name":"Name2","event":"ReceptionOnly","attending":false,"count":2,"comments":"comment 2"}`,
			Rsvp{"Name2", receptionOnly, false, 2, "comment 2"}},

		{`{"name":"Name3","attending":true,"count":3,"comments":"comment 3"}`,
			Rsvp{Name: "Name3", Attending: true, Count: 3, Comments: "comment 3"}},

		{`{"event":"CeremonyAndReception","attending":false,"count":4}`,
			Rsvp{Event: ceremonyAndReception, Attending: false, Count: 4}},

		{`{"attending":false}`,
			Rsvp{}},
	} {

		actualOutput := NewRsvp()
		err := json.Unmarshal([]byte(this.input), &actualOutput)

		if err != nil {
			t.Errorf("[%d] unexpected error: %v", i, err)
		} else {
			if this.expectedOutput != *actualOutput {
				t.Errorf("[%d] unexpected json result: '%v'.", i, actualOutput)
			}
		}
	}
}
