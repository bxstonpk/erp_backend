from database.imports import *

class Invoice(Base):
    __tablename__ = "invoices"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    branch_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("branches.id"))

    invoice_no: Mapped[str] = mapped_column(String(30), index=True)
    invoice_type: Mapped[str] = mapped_column(String(30), default="BILLING_NOTE")  # BILLING_NOTE, TAX_INVOICE, RECEIPT_TAX_INVOICE, PURCHASE_INVOICE, DEBIT_NOTE, CREDIT_NOTE
    tax_invoice_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("invoices.id"), nullable=True)
    tax_invoice_date: Mapped[Optional[datetime]] = mapped_column(Date, nullable=True)
    tax_invoice_no: Mapped[Optional[str]] = mapped_column(String(30), nullable=True)
    payment_received_date: Mapped[Optional[datetime]] = mapped_column(Date, nullable=True)
    tax_on_payment: Mapped[bool] = mapped_column(Boolean, default=False)  # True = issue tax invoice on payment (service)

    ref_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), nullable=True)  # SO or PO id
    ref_module: Mapped[Optional[str]] = mapped_column(String(20), nullable=True)  # SO, PO, BN, MANUAL
    partner_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("business_partners.id"))

    invoice_date: Mapped[datetime] = mapped_column(Date, index=True)
    due_date: Mapped[datetime] = mapped_column(Date)
    currency_code: Mapped[str] = mapped_column(String(3), default="THB")
    exchange_rate: Mapped[float] = mapped_column(Numeric(18, 6), default=1)

    subtotal: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    discount_amount: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    tax_amount: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    total_amount: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    paid_amount: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    outstanding: Mapped[Optional[float]] = mapped_column(Numeric(18, 2), nullable=True)  # computed: total_amount - paid_amount

    status: Mapped[str] = mapped_column(String(10), default="DRAFT", index=True)  # DRAFT, POSTED, PARTIAL, PAID, OVERDUE, CANCELLED, REVERSED
    journal_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("journals.id"), nullable=True)
    created_by: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)

    lines = relationship("InvoiceLine", back_populates="invoice")
