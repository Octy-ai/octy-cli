package churnpredictionreport

// octy churn predcition domain models

// OctyChurnPredictionReport : model representing Octy churn prediction report
type OctyChurnPredictionReport struct {
	ChurnPredictionReport ChurnPredictionReport `json:"churn_prediction_report"`
}

type ChurnPredictionReport struct {
	TrainingJobData TrainingJobData `json:"training_job_data"`
	ChurnData       ChurnData       `json:"churn_data"`
}

type ChurnData struct {
	CurrentChurnPercentage    float64                `json:"current_churn_percentage"`
	ChurnDirectionIndication  string                 `json:"churn_direction_indication"`
	ChurnPercentageDifference float64                `json:"churn_percentage_difference"`
	FeaturesOfImportance      []FeaturesOfImportance `json:"features_of_importance"`
}

type FeaturesOfImportance struct {
	FeatureName       string  `json:"feature_name"`
	FeatureImportance float64 `json:"feature_importance"`
}

type TrainingJobData struct {
	TrainingJobID   string  `json:"training_job_id"`
	ModelAccuracy   float64 `json:"model_accuracy"`
	TrainingJobDate string  `json:"training_job_date"`
}

// NewCPR : returns a pointer to a new OctyChurnPredictionReport instance
func NewCPR() *OctyChurnPredictionReport {
	return &OctyChurnPredictionReport{}
}
