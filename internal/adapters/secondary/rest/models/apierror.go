package models

import "encoding/json"

type OctyErrorResp struct {
	RequestMeta RequestMeta        `json:"request_meta"`
	Error       OctyErrorRespError `json:"error"`
}

type OctyErrorRespError struct {
	Code   int64          `json:"code"`
	Reason string         `json:"reason"`
	Errors []ErrorElement `json:"errors"`
}

type ErrorElement struct {
	Message      string `json:"message"`
	ExtendedHelp string `json:"extended_help"`
}

func UnmarshalOctyErrorResp(data []byte) (OctyErrorResp, error) {
	var r OctyErrorResp
	err := json.Unmarshal(data, &r)
	return r, err
}
