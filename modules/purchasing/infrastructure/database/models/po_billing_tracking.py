from database.imports import *

class PoBillingTracking(Base):
    __tablename__ = "po_billing_tracking"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    po_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("purchase_orders.id"), index=True)
    po_line_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("purchase_order_lines.id"))
    invoice_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("invoices.id"))
    invoice_line_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("invoice_lines.id"))

    billed_qty: Mapped[float] = mapped_column(Numeric(18, 4))
    billed_amount: Mapped[float] = mapped_column(Numeric(18, 2))
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
