package models

import "encoding/json"

// ** Octy REST Response Models **

type OctyGetChurnPredictionReportResp struct {
	RequestMeta           RequestMeta           `json:"request_meta"`
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

func UnmarshalOctyGetChurnPredictionReportResp(data []byte) (OctyGetChurnPredictionReportResp, error) {
	var r OctyGetChurnPredictionReportResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---
