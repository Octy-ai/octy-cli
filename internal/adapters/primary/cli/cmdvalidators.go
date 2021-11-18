package cli

import (
	"fmt"

	v "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Account configurations

func (a AccountConf) Validate() error {

	//validate parent struct
	err := v.ValidateStruct(&a,
		v.Field(&a.Configurations, v.Required.Error("invalid configuration yaml file provided. Missing required key: 'configs'")),
	)
	if err != nil {
		return err
	}

	// validate embedded struct
	configs := a.Configurations
	err = v.ValidateStruct(&configs,
		v.Field(&configs.ContactName, v.Required.Error("invalid configuration yaml file provided. Missing required key: 'contactName'")),
		v.Field(&configs.ContactSurname, v.Required.Error("invalid configuration yaml file provided. Missing required key: 'contactSurame'")),
		v.Field(&configs.ContactEmailAddress, v.Required.Error("invalid configuration yaml file provided. Missing required key: 'contactEmail'"), is.Email),
		v.Field(&configs.WebhookURL, v.Required.Error("invalid configuration yaml file provided. Missing required key: 'webhookURL'"), is.URL),
	)
	if err != nil {
		return err
	}
	return nil
}

// Algorithm configurations

func (a AlgoConf) Validate() error {

	var err error

	//validate parent struct
	err = v.ValidateStruct(&a,
		v.Field(&a.Configurations, v.Required.Error("invalid configuration yaml file provided. Missing required key: 'configurations'")),
	)
	if err != nil {
		return err
	}

	// validate embedded structs
	configs := a.Configurations
	for idx, config := range configs {
		switch config.AlgorithmName {
		case "rec":
			err = v.ValidateStruct(&config.Configurations,
				v.Field(&config.Configurations.ItemIDStopList, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required key: 'configurations[%v].itemIDStopList'", idx))),
				v.Field(&config.Configurations.ProfileFeatures, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required key: 'configurations[%v].profileFeatures'", idx))),
			)
		case "churn":
			err = v.ValidateStruct(&config.Configurations,
				v.Field(&config.Configurations.ProfileFeatures, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required key: 'configurations[%v].profileFeatures'", idx))),
			)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

// Event types
func (e EventTypes) Validate() error {
	var err error

	//validate parent struct
	err = v.ValidateStruct(&e,
		v.Field(&e.EventTypes, v.Required.Error("invalid configuration yaml file provided. Missing required key: 'eventTypeDefinitions'")),
	)
	if err != nil {
		return err
	}

	// validate embedded structs
	eventTypes := e.EventTypes
	for idx, eventType := range eventTypes {
		err = v.ValidateStruct(&eventType,
			v.Field(&eventType.EventType, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required key: 'eventTypeDefinitions[%v].eventType'", idx))),
			v.Field(&eventType.EventProperties, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required key: 'eventTypeDefinitions[%v].eventProperties'", idx))),
		)
		if err != nil {
			return err
		}
	}

	return nil

}
