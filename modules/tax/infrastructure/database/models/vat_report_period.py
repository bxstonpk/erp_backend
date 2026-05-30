from database.imports import *

class VatReportPeriod(Base):
    __tablename__ = "vat_report_periods"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    branch_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("branches.id"))  # branch filing ภ.พ.30

    report_year: Mapped[int] = mapped_column(Integer, index=True)
    report_month: Mapped[int] = mapped_column(Integer)  # 1-12
    due_date: Mapped[datetime] = mapped_column(Date)

    total_output_base: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    total_output_vat: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    total_input_base: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    total_input_vat: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    vat_payable: Mapped[Optional[float]] = mapped_column(Numeric(18, 2), nullable=True)  # computed: output_vat - input_vat

    status: Mapped[str] = mapped_column(String(10), default="DRAFT", index=True)  # DRAFT, SUBMITTED, AMENDED
    submitted_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
    submitted_by: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"), nullable=True)
    reference_no: Mapped[Optional[str]] = mapped_column(String(50), nullable=True)
    payment_date: Mapped[Optional[datetime]] = mapped_column(Date, nullable=True)
    payment_amount: Mapped[Optional[float]] = mapped_column(Numeric(18, 2), nullable=True)
    penalty_amount: Mapped[Optional[float]] = mapped_column(Numeric(18, 2), default=0)
    surcharge_amount: Mapped[Optional[float]] = mapped_column(Numeric(18, 2), default=0)
    notes: Mapped[Optional[str]] = mapped_column(Text, nullable=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
