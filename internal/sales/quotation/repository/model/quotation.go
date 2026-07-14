package model

type Quotation struct {
	QuotationUUID         string  `db:"uuid" gorm:"primaryKey" json:"quotation_uuid"`
	CompanyUUID           string  `db:"company_uuid" json:"company_uuid"`
	CompanyBranchUUID     string  `db:"company_branch_uuid" json:"company_branch_uuid"`
	ProjectUUID           string  `db:"project_uuid" json:"project_uuid"`
	CustomerUUID          string  `db:"customer_uuid" json:"customer_uuid"`
	CustomerBranchUUID    string  `db:"customer_branch_uuid" json:"customer_branch_uuid"`
	RecipientCustomerUUID string  `db:"recipient_customer_uuid" json:"recipient_customer_uuid"`
	QuotationNumber       string  `db:"quotation_number" gorm:"unique" json:"quotation_number"`
	QuotationDate         string  `db:"quotation_date" json:"quotation_date"`
	QuotationDueDate      string  `db:"quotation_due_date" json:"quotation_due_date"`
	QuotationStatus       string  `db:"quotation_status" json:"quotation_status"`
	QuotationType         string  `db:"quotation_type" json:"quotation_type"`
	QuotationRemarks      string  `db:"quotation_remarks" json:"quotation_remarks"`
	VatTypeUUID           string  `db:"vat_type_uuid" json:"vat_type_uuid"`
	DiscountAmount        float64 `db:"discount_amount" json:"discount_amount"`
	DiscountPercentage    float64 `db:"discount_percentage" json:"discount_percentage"`
	TotalAmount           float64 `db:"total_amount" json:"total_amount"`
	MakerUUID             string  `db:"maker_uuid" json:"maker_uuid"`
	CheckerUUID           string  `db:"checker_uuid" json:"checker_uuid"`
	ApproverUUID          string  `db:"approver_uuid" json:"approver_uuid"`
}
