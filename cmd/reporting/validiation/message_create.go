package validation

import (
	"twitch_chat_analysis/cmd/reporting/model"

	validator "gopkg.in/asaskevich/govalidator.v9"
)

func ValidateCreateMessageRequest(message model.Message) error {
	if _, err := validator.ValidateStruct(message); err != nil {
		return err
	}

	return nil
}
