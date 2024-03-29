package dto

import "time"

type RequestCurrencyDTO struct {
	FullName    string `xml:"fullname"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Quant       int    `xml:"quant"`
	Index       string `xml:"index"`
}

type CreateCurrencDTO struct {
	Title string
	Code  string
	Value string
	Date  time.Time
}

type ResponseCurrencyDTO struct {
	ID    uint   `json:"id"`
	Code  string `json:"code"`
	Value string `json:"value"`
	Date  string `json:"date"`
}
