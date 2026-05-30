from database.imports import *

class PurchaseOrder(Base):
    __tablename__ = "purchase_orders"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    branch_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("branches.id"))
    department_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("departments.id"), nullable=True)

    po_no: Mapped[str] = mapped_column(String(30), index=True)
    vendor_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("business_partners.id"))
    po_date: Mapped[datetime] = mapped_column(Date, index=True)
    expected_date: Mapped[Optional[datetime]] = mapped_column(Date, nullable=True)

    currency_code: Mapped[str] = mapped_column(String(3), default="THB")
    exchange_rate: Mapped[float] = mapped_column(Numeric(18, 6), default=1)
    status: Mapped[str] = mapped_column(String(10), default="DRAFT", index=True)  # DRAFT, CONFIRMED, PARTIAL, RECEIVED, CANCELLED

    subtotal: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    tax_amount: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    total_amount: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    notes: Mapped[Optional[str]] = mapped_column(Text, nullable=True)

    approved_by: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"), nullable=True)
    approved_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
    created_by: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)

    vendor = relationship("BusinessPartner")
    lines = relationship("PurchaseOrderLine", back_populates="purchase_order")
