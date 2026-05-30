from database.imports import *

class WhtReportPeriod(Base):
    __tablename__ = "wht_report_periods"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    branch_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("branches.id"))

    form_type: Mapped[str] = mapped_column(String(5))  # ND1, ND3, ND53
    report_year: Mapped[int] = mapped_column(Integer, index=True)
    report_month: Mapped[int] = mapped_column(Integer)
    due_date: Mapped[datetime] = mapped_column(Date)

    total_income: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    total_wht: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    penalty_amount: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    surcharge_amount: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    status: Mapped[str] = mapped_column(String(10), default="DRAFT", index=True)  # DRAFT, SUBMITTED, AMENDED

    submitted_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
    submitted_by: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"), nullable=True)
    reference_no: Mapped[Optional[str]] = mapped_column(String(50), nullable=True)
    payment_date: Mapped[Optional[datetime]] = mapped_column(Date, nullable=True)
    payment_amount: Mapped[Optional[float]] = mapped_column(Numeric(18, 2), nullable=True)
    notes: Mapped[Optional[str]] = mapped_column(Text, nullable=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)

    transactions = relationship("WhtTransaction", back_populates="report_period")
