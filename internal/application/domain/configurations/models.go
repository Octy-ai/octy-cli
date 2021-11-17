package configurations

// octy configuration domain models

// OctyAccountConfigurations : model representing Octy account configurations
type OctyAccountConfigurations struct {
	ContactName         string `json:"contact_name" yaml:"contactName"`
	ContactSurname      string `json:"contact_surname" yaml:"contactSurname"`
	ContactEmailAddress string `json:"contact_email_address" yaml:"contactEmail"`
	WebhookURL          string `json:"webhook_url" yaml:"webhookURL"`
	AuthenticatedIDKey  string `json:"authenticated_id_key" yaml:"authenticatedIDKey"`
}

// OctyAlgorithmConfiguration : model representing Octy churn prediction & recommendation algorithm configurations
type OctyAlgorithmConfiguration struct {
	AlgorithmName  string         `json:"algorithm_name" yaml:"algorithmName"`
	Configurations Configurations `json:"configurations" yaml:"configurations"`
}

type Configurations struct {
	RecommendInteractedItems bool     `json:"recommend_interacted_items,omitempty" yaml:"recommendInteractedItems,omitempty"`
	ItemIDStopList           []string `json:"item_id_stop_list,omitempty" yaml:"itemIDStopList,omitempty"`
	ProfileFeatures          []string `json:"profile_features" yaml:"profileFeatures"`
}

//NewAcc : returns a pointer to a new OctyAccountConfigurations instance
func NewAcc() *OctyAccountConfigurations {
	return &OctyAccountConfigurations{}
}

//NewAlg : returns a pointer to a new slice of OctyAlgorithmConfiguration instances
func NewAlg() *[]OctyAlgorithmConfiguration {
	return &[]OctyAlgorithmConfiguration{}
}
