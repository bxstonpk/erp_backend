from database.imports import *

class Payment(Base):
    __tablename__ = "payments"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    branch_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("branches.id"))

    payment_no: Mapped[str] = mapped_column(String(30), index=True)
    payment_type: Mapped[str] = mapped_column(String(10))  # RECEIPT=AR, PAYMENT=AP
    partner_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("business_partners.id"))
    payment_date: Mapped[datetime] = mapped_column(Date, index=True)
    payment_method: Mapped[str] = mapped_column(String(20), default="TRANSFER")  # CASH, TRANSFER, CHEQUE, CREDIT_CARD, OTHER
    bank_account: Mapped[Optional[str]] = mapped_column(String(100), nullable=True)
    cheque_no: Mapped[Optional[str]] = mapped_column(String(50), nullable=True)

    currency_code: Mapped[str] = mapped_column(String(3), default="THB")
    exchange_rate: Mapped[float] = mapped_column(Numeric(18, 6), default=1)
    amount: Mapped[float] = mapped_column(Numeric(18, 2))
    withholding_tax: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    wht_transaction_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("wht_transactions.id"), nullable=True)
    wht_certificate_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("wht_certificates.id"), nullable=True)
    net_amount: Mapped[float] = mapped_column(Numeric(18, 2))

    status: Mapped[str] = mapped_column(String(10), default="DRAFT", index=True)  # DRAFT, POSTED, CANCELLED
    journal_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("journals.id"), nullable=True)
    created_by: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)

    allocations = relationship("PaymentAllocation", back_populates="payment")
