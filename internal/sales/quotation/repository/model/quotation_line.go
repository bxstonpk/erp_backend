package model

type QuotationLine struct {
	QuotationLineUUID string `db:"uuid" gorm:"primaryKey" json:"quotation_line_uuid"`
	QuotationUUID     string `db:"quotation_uuid" json:"quotation_uuid"`
	VatTypeUUID       string `db:"vat_type_uuid" json:"vat_type_uuid"`
	ProjectUUID       string `db:"project_uuid" json:"project_uuid"`
	ItemUUID          string `db:"item_uuid" json:"item_uuid"`
	Quantity          string `db:"quantity" json:"quantity"`
	UnitUUID          string `db:"unit_uuid" json:"unit_uuid"`
	UnitPrice         string `db:"unit_price" json:"unit_price"`
	DiscountAmount    string `db:"discount_amount" json:"discount_amount"`
	TotalAmount       string `db:"total_amount" json:"total_amount"`
}
