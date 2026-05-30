from database.imports import *

class Journal(Base):
    __tablename__ = "journals"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    branch_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("branches.id"))
    fiscal_period_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("fiscal_periods.id"))

    journal_no: Mapped[str] = mapped_column(String(30), index=True)
    journal_type: Mapped[str] = mapped_column(String(20))  # MANUAL, SALES, PURCHASE, PAYMENT, RECEIPT, CLOSING, ADJUSTMENT
    ref_module: Mapped[Optional[str]] = mapped_column(String(30), nullable=True)  # SO, PO, AP, AR ...
    ref_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), nullable=True)
    transaction_date: Mapped[datetime] = mapped_column(Date, index=True)
    description: Mapped[Optional[str]] = mapped_column(String(500), nullable=True)
    status: Mapped[str] = mapped_column(String(10), default="DRAFT", index=True)  # DRAFT, POSTED, REVERSED

    created_by: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"))
    posted_by: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"), nullable=True)
    posted_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)

    lines = relationship("JournalLine", back_populates="journal")
