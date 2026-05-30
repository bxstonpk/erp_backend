from database.imports import *

class GoodsReceipt(Base):
    __tablename__ = "goods_receipts"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    branch_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("branches.id"))

    grn_no: Mapped[str] = mapped_column(String(30), index=True)
    po_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("purchase_orders.id"), nullable=True)
    vendor_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("business_partners.id"))
    receipt_date: Mapped[datetime] = mapped_column(Date)
    status: Mapped[str] = mapped_column(String(10), default="DRAFT", index=True)  # DRAFT, POSTED, CANCELLED
    journal_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("journals.id"), nullable=True)

    created_by: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)

    lines = relationship("GoodsReceiptLine", back_populates="goods_receipt")
