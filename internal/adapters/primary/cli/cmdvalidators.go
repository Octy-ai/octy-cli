package cli

import (
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
	for _, config := range configs {
		switch config.AlgorithmName {
		case "rec":
			err = v.ValidateStruct(&config.Configurations,
				v.Field(&config.Configurations.ItemIDStopList, v.Required.Error("invalid configuration yaml file provided. Missing required key: 'configurations.itemIDStopList'")),
				v.Field(&config.Configurations.ProfileFeatures, v.Required.Error("invalid configuration yaml file provided. Missing required key: 'configurations.profileFeatures'")),
			)
		case "churn":
			err = v.ValidateStruct(&config.Configurations,
				v.Field(&config.Configurations.ProfileFeatures, v.Required.Error("invalid configuration yaml file provided. Missing required key: 'configurations.profileFeatures'")),
			)
		}
	}
	if err != nil {
		return err
	}
	return nil
}
