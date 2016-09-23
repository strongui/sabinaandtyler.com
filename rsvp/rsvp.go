package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Rsvp struct {
	Name      string `json:"name,omitempty"`
	Event     Event  `json:"event,omitempty"`
	Attending bool   `json:"attending"`
	Count     int    `json:"count"`
	Comments  string `json:"comments,omitempty"`
}

type RsvpSendOutput struct {
	Success bool `json:"success"`
}

func NewRsvp() (rsvp *Rsvp) {
	return new(Rsvp)
}

func (rsvp *Rsvp) Send() (result RsvpSendOutput, err error) {

	if len(strings.TrimSpace(rsvp.Name)) == 0 {
		err = errors.New("name field missing")
		result = RsvpSendOutput{false}
		return
	}

	svc := ses.New(session.New())

	params := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String("Nate Morse <nmorse@gmail.com>"),
				aws.String("Molly Chase <molly.chase@gmail.com>"),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(rsvp.MessageBody()), // Required
				},
			},
			Subject: &ses.Content{
				Data: aws.String(rsvp.Subject()),
			},
		},
		Source: aws.String("Wedding RSVP <wedding@mollyandnate.com>"),
		ReplyToAddresses: []*string{
			aws.String("Nate Morse <nmorse@gmail.com>"),
		},
	}

	_, e := svc.SendEmail(params)
	if e != nil {
		// failure
		err = e
		result = RsvpSendOutput{false}
		return
	}

	// success
	result = RsvpSendOutput{true}
	return
}

func (rsvp *Rsvp) Subject() string {

	prefix := "Wedding RSVP - "
	if rsvp.Attending {
		if rsvp.Event == ceremonyAndReception {
			return prefix + "Ceremony & Reception"
		} else {
			return prefix + "Reception Only"
		}
	}

	return prefix + "Regrets"
}

func (rsvp *Rsvp) MessageBody() string {

	prefix := "Yay! Molly and Nate! You have an RSVP to your wedding!\n\n"
	attending := "YES"

	if rsvp.Event == receptionOnly {
		prefix = "Looks like someone is ready to party!\n\n"
	}

	if !rsvp.Attending {
		prefix = "Oh no! Someone can't make it to your wedding!\n\n"
		attending = "NO"
	}

	details := fmt.Sprintf(
		"Name: %s\nAttending: %v\nEvent: %s\nNumber of Guests: %v\nComments: %v",
		rsvp.Name,
		attending,
		rsvp.Event.FriendlyName(),
		rsvp.Count,
		rsvp.Comments)

	return prefix + details
}
