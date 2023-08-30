package models

type Quotation_USDBRL struct {
	UsdBrl Quotation `json:"USDBRL"`
}

type Quotation struct {
    Id string `json:"-" gorm:"primaryKey"`
    Name string `json:"name"`
    Code string `json:"code"`
    CodeIn string `json:"codein"`
    High string `json:"high"`
    Low string `json:"low"`
	Bid string `json:"bid"`
}
