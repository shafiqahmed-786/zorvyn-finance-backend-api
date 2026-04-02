package dto

type CreateRecordRequest struct {
    Amount   float64 `json:"amount"`
    Type     string  `json:"type"`
    Category string  `json:"category"`
    Date     string  `json:"date"`
    Notes    string  `json:"notes"`
}