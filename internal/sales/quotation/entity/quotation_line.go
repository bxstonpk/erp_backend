package entity

type QuotationLine struct {
	Id          string
	QuotationId string
	VatTypeId   string
	ProjectId   string
	ItemId      string
	Quantity    float64
	UnitId      string
	UnitPrice   float64
	Discount    float64
	Total       float64
}
