from database.imports import *

class InvoiceLine(Base):
    __tablename__ = "invoice_lines"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    invoice_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("invoices.id"), index=True)
    line_no: Mapped[int] = mapped_column(Integer)
    product_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("products.id"), nullable=True)
    description: Mapped[str] = mapped_column(String(500))
    uom_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("units_of_measure.id"), nullable=True)

    qty: Mapped[float] = mapped_column(Numeric(18, 4), default=1)
    unit_price: Mapped[float] = mapped_column(Numeric(18, 4))
    discount_pct: Mapped[float] = mapped_column(Numeric(5, 2), default=0)
    line_amount: Mapped[float] = mapped_column(Numeric(18, 2))
    tax_type: Mapped[str] = mapped_column(String(10), default="VAT7")  # NO_TAX, VAT7, EXEMPT
    tax_amount: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    account_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("accounts.id"), nullable=True)
    ref_line_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), nullable=True)  # SO/PO line populated from
    ref_module: Mapped[Optional[str]] = mapped_column(String(20), nullable=True)  # SO / PO
    billed_qty: Mapped[Optional[float]] = mapped_column(Numeric(18, 4), nullable=True)  # qty already billed (partial billing)

    invoice = relationship("Invoice", back_populates="lines")
