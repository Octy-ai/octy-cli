package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

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

// ParseErrors: parse returned errors into slice of errors
func ParseErrors(errResp OctyErrorResp) []error {
	var errs []error
	for _, e := range errResp.Error.Errors {
		var errMsg string
		if e.Message == "" {
			errMsg = errResp.Error.Reason
		} else {
			errMsg = e.Message
		}
		errs = append(errs, fmt.Errorf("apierror[%s]: %v", strconv.Itoa(int(errResp.Error.Code)), errMsg))
	}
	return errs
}
