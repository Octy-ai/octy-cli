package models

import "encoding/json"

// ** Octy REST Response Models **

// ---

type OctyCreateItemsResp struct {
	RequestMeta    RequestMeta          `json:"request_meta"`
	Items          []Item               `json:"items"`
	FailedToCreate []FailedToCreateItem `json:"failed_to_create"`
}

type FailedToCreateItem struct {
	ItemID       string `json:"item_id"`
	ErrorMessage string `json:"error_message"`
}

type Item struct {
	ItemID          string `json:"item_id"`
	ItemCategory    string `json:"item_category"`
	ItemName        string `json:"item_name"`
	ItemDescription string `json:"item_description"`
	ItemPrice       int64  `json:"item_price"`
}

func UnmarshalOctyCreateItemsResp(data []byte) (OctyCreateItemsResp, error) {
	var r OctyCreateItemsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---
