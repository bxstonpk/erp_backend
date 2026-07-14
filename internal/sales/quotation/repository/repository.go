package repository

import (
	"erp/backend/internal/sales/quotation/entity"
)

type QuotationRepository interface {
	CreateQuotation(quotation *entity.Quotation) error
	UpdateQuotation(quotation *entity.Quotation) error
	DeleteQuotation(id string) error
	GetQuotation(id string) (*entity.Quotation, error)
	GetAllQuotation() ([]*entity.Quotation, error)
}
