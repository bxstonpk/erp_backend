from database.imports import *

class TaxInvoiceQueue(Base):
    __tablename__ = "tax_invoice_queue"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    billing_note_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("invoices.id"))  # invoice that is a BILLING_NOTE
    payment_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("payments.id"))  # payment that triggered issuing

    status: Mapped[str] = mapped_column(String(10), default="PENDING", index=True)  # PENDING, PROCESSING, DONE, FAILED
    tax_invoice_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("invoices.id"), nullable=True)
    error_msg: Mapped[Optional[str]] = mapped_column(Text, nullable=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    processed_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
