package model

type QuotationDetail struct {
	QuotationDetailUUID string `db:"uuid" gorm:"primaryKey" json:"quotation_detail_uuid"`
	QuotationLineUUID   string `db:"quotation_line_uuid" json:"quotation_line_uuid"`
	Description         string `db:"description" json:"description"`
	ShowDisplay         bool   `db:"show_display" json:"show_display"`
}
