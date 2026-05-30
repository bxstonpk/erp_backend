from database.imports import *

class WhtTransaction(Base):
    __tablename__ = "wht_transactions"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    branch_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("branches.id"))
    fiscal_period_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("fiscal_periods.id"))

    wht_direction: Mapped[Optional[str]] = mapped_column(String(10), nullable=True)  # WITHHELD (we withhold), SUFFERED (withheld from us)
    ref_module: Mapped[str] = mapped_column(String(30))  # PAY=Payment, INV=Invoice
    ref_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True))
    ref_doc_no: Mapped[str] = mapped_column(String(50))
    payment_date: Mapped[datetime] = mapped_column(Date, index=True)

    partner_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("business_partners.id"))
    partner_name: Mapped[str] = mapped_column(String(255))  # snapshot
    partner_tax_id: Mapped[str] = mapped_column(String(20))  # snapshot
    partner_type: Mapped[str] = mapped_column(String(10))  # INDIVIDUAL, JURISTIC
    wht_type_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("wht_types.id"))
    wht_description: Mapped[Optional[str]] = mapped_column(String(255), nullable=True)

    income_amount: Mapped[float] = mapped_column(Numeric(18, 2))  # assessable income (before withholding)
    wht_rate: Mapped[float] = mapped_column(Numeric(5, 2))  # snapshot rate
    wht_amount: Mapped[float] = mapped_column(Numeric(18, 2))
    net_paid_amount: Mapped[Optional[float]] = mapped_column(Numeric(18, 2), nullable=True)  # computed: income_amount - wht_amount
    currency_code: Mapped[str] = mapped_column(String(3), default="THB")
    exchange_rate: Mapped[float] = mapped_column(Numeric(18, 6), default=1)
    income_amount_thb: Mapped[float] = mapped_column(Numeric(18, 2))
    wht_amount_thb: Mapped[float] = mapped_column(Numeric(18, 2))

    journal_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("journals.id"), nullable=True)
    wht_account_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("accounts.id"), nullable=True)
    certificate_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("wht_certificates.id"), nullable=True)
    is_remitted: Mapped[bool] = mapped_column(Boolean, default=False)
    report_period_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("wht_report_periods.id"), nullable=True)

    created_by: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)

    wht_type = relationship("WhtType", back_populates="transactions")
    report_period = relationship("WhtReportPeriod", back_populates="transactions")
