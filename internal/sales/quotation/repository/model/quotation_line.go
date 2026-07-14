package model

type QuotationLine struct {
	QuotationLineUUID string  `db:"uuid" gorm:"primaryKey" json:"quotation_line_uuid"`
	QuotationUUID     string  `db:"quotation_uuid" json:"quotation_uuid"`
	VatTypeUUID       string  `db:"vat_type_uuid" json:"vat_type_uuid"`
	ProjectUUID       string  `db:"project_uuid" json:"project_uuid"`
	ItemUUID          string  `db:"item_uuid" json:"item_uuid"`
	Quantity          float64 `db:"quantity" json:"quantity"`
	UnitUUID          string  `db:"unit_uuid" json:"unit_uuid"`
	UnitPrice         float64 `db:"unit_price" json:"unit_price"`
	DiscountAmount    float64 `db:"discount_amount" json:"discount_amount"`
	TotalAmount       float64 `db:"total_amount" json:"total_amount"`
}
