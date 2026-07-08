package repository

type QuotationRepository interface {
	CreateQuotation()
	UpdateQuotation()
	DeleteQuotation()
	GetQuotation()
	GetAllQuotation()
}
