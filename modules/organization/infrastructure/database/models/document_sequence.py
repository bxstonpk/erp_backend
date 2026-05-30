from database.imports import *

class DocumentSequence(Base):
    __tablename__ = "document_sequences"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    branch_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("branches.id"), nullable=True)

    doc_type: Mapped[str] = mapped_column(String(30), index=True)  # SO, PO, INV, JV, GRN, DO, REC, PAY ...
    prefix: Mapped[Optional[str]] = mapped_column(String(10), nullable=True)
    year_format: Mapped[str] = mapped_column(String(10), default="YY")  # NONE, YY, YYYY
    month_format: Mapped[str] = mapped_column(String(10), default="MM")  # NONE, MM
    sep_char: Mapped[Optional[str]] = mapped_column(String(5), default="-")
    padding: Mapped[int] = mapped_column(Integer, default=5)
    current_number: Mapped[int] = mapped_column(Integer, default=0)
    reset_cycle: Mapped[str] = mapped_column(String(10), default="YEARLY")  # NEVER, YEARLY, MONTHLY
    last_reset_date: Mapped[Optional[datetime]] = mapped_column(Date, nullable=True)
