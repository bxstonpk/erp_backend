package entity

type Quotation struct {
	Id                 string
	CompanyId          string
	BranchId           string
	ProjectId          string
	CustomerId         string
	CustomerBranchId   string
	BillToId           string
	QuotationNumber    string
	QuotationDate      string
	QuotationDueDate   string
	QuotationStatus    string
	QuotationType      string
	Remarks            string
	VatTypeId          string
	DiscountAmount     float64
	DiscountPercentage float64
	TotalAmount        float64
	MakerId            string
	CheckerId          string
	ApproverId         string

	QuotationLines   []QuotationLine
	QuotationDetails []QuotationDetail
}
