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
		v.Field(&a.Configurations, v.Required.Error("invalid configuration yaml file provided. Missing required value for key: 'configurations'")),
	)
	if err != nil {
		return err
	}

	// validate embedded struct
	configs := a.Configurations
	err = v.ValidateStruct(&configs,
		v.Field(&configs.ContactName, v.Required.Error("invalid configuration yaml file provided. Missing required value for key: 'contactName'")),
		v.Field(&configs.ContactSurname, v.Required.Error("invalid configuration yaml file provided. Missing required value for key: 'contactSurame'")),
		v.Field(&configs.ContactEmailAddress, v.Required.Error("invalid configuration yaml file provided. Missing required value for key: 'contactEmail'"), is.Email),
		v.Field(&configs.WebhookURL, v.Required.Error("invalid configuration yaml file provided. Missing required value for key: 'webhookURL'"), is.URL),
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
		v.Field(&a.Configurations, v.Required.Error("invalid configuration yaml file provided. Missing required value for key: 'configurations'")),
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
				v.Field(&config.Configurations.ItemIDStopList, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'configurations[%v].itemIDStopList'", idx))),
				v.Field(&config.Configurations.ProfileFeatures, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'configurations[%v].profileFeatures'", idx))),
			)
		case "churn":
			err = v.ValidateStruct(&config.Configurations,
				v.Field(&config.Configurations.ProfileFeatures, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'configurations[%v].profileFeatures'", idx))),
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
		v.Field(&e.EventTypes, v.Required.Error("invalid configuration yaml file provided. Missing required value for key: 'eventTypeDefinitions'")),
	)
	if err != nil {
		return err
	}

	// validate embedded structs
	eventTypes := e.EventTypes
	for idx, eventType := range eventTypes {
		err = v.ValidateStruct(&eventType,
			v.Field(&eventType.EventType, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'eventTypeDefinitions[%v].eventType'", idx))),
			v.Field(&eventType.EventProperties, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'eventTypeDefinitions[%v].eventProperties'", idx))),
		)
		if err != nil {
			return err
		}
	}

	return nil

}

// Segments

func (s Segments) Validate() error {
	var err error
	//validate parent struct
	err = v.ValidateStruct(&s,
		v.Field(&s.Segments, v.Required.Error("invalid configuration yaml file provided. Missing required value for key: 'segmentDefinitions'")),
	)
	if err != nil {
		return err
	}

	segments := s.Segments
	for idx, segment := range segments {
		err = v.ValidateStruct(&segment,
			v.Field(&segment.SegmentName, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'segmentDefinitions[%v].segmentName'", idx))),
			v.Field(&segment.SegmentType, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'segmentDefinitions[%v].segmentType'", idx))),
			v.Field(&segment.EventSequence, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'segmentDefinitions[%v].eventSequence'", idx))),
		)
		if err != nil {
			return err
		}
		for eidx, e := range segment.EventSequence {
			err = v.ValidateStruct(&e,
				v.Field(&e.EventType, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'segmentDefinitions[%v].eventSequence[%v].eventType'", idx, eidx))),
				v.Field(&e.ActionInaction, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'segmentDefinitions[%v].eventSequence[%v].actionInaction'", idx, eidx))),
			)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

// Templates
func (t Templates) Validate() error {
	var err error
	//validate parent struct
	err = v.ValidateStruct(&t,
		v.Field(&t.Templates, v.Required.Error("invalid configuration yaml file provided. Missing required value for key: 'templates'")),
	)
	if err != nil {
		return err
	}

	templates := t.Templates
	for idx, template := range templates {
		err = v.ValidateStruct(&template,
			v.Field(&template.FriendlyName, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'templates[%v].friendlyName'", idx))),
			v.Field(&template.TemplateType, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'templates[%v].templateType'", idx))),
			v.Field(&template.Title, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'templates[%v].title'", idx))),
			v.Field(&template.Content, v.Required.Error(fmt.Sprintf("invalid configuration yaml file provided. Missing required value for key: 'templates[%v].content'", idx))),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: set validation limits on creation of objects. Max 500 per command (where applicable)
