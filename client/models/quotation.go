package models

type Quotation struct {
    Name string `json:"name"`
    Code string `json:"code"`
    CodeIn string `json:"codein"`
    High string `json:"high"`
    Low string `json:"low"`
	Bid string `json:"bid"`
}
