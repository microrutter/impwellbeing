package constants

import (
	"github.com/impwellbeing/services/postgres"
	"github.com/nu7hatch/gouuid"
)

var (
	QuestionList = []postgres.Question{
		{Uuid: returnUUID(), Type: "onetofour", Title: "How overloaded with pressure at work are you?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "How overwhelmed do you feel?"},

		{Uuid: returnUUID(), Type: "yes/no", Title: "Do you struggle to switch off from your demands at work?"},

		{Uuid: returnUUID(), Type: "yes/no", Title: "Do you feel in control to make changes at work?"},

		{Uuid: returnUUID(), Type: "yes/no", Title: "Do you feel a strong sense of purpose at work?"},

		{Uuid: returnUUID(), Type: "yes/no", Title: "Do you feel supported at work?"},

		{Uuid: returnUUID(), Type: "yes/no", Title: "Do you feel positive about work?"},

		{Uuid: returnUUID(), Type: "yes/no", Title: "Do you feel exhausted?"},

		{Uuid: returnUUID(), Type: "yes/no", Title: "Can you speak openly about Mental Health and well-being?"},
	}
)

func returnUUID() *uuid.UUID {
	uuid, _ := uuid.NewV4()
	return uuid
}