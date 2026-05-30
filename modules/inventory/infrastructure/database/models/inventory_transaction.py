from database.imports import *

class InventoryTransaction(Base):
    __tablename__ = "inventory_transactions"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    warehouse_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("warehouses.id"), index=True)
    location_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("warehouse_locations.id"), nullable=True)
    product_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("products.id"), index=True)

    txn_type: Mapped[str] = mapped_column(String(20))  # RECEIVE, ISSUE, ADJUST, TRANSFER_IN, TRANSFER_OUT, RETURN_IN, RETURN_OUT
    ref_module: Mapped[Optional[str]] = mapped_column(String(30), nullable=True)  # PO, SO, TRANSFER, ADJUST
    ref_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), nullable=True)
    ref_line_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), nullable=True)
    txn_date: Mapped[datetime] = mapped_column(Date, index=True)

    qty: Mapped[float] = mapped_column(Numeric(18, 4))
    unit_cost: Mapped[float] = mapped_column(Numeric(18, 4), default=0)
    total_cost: Mapped[Optional[float]] = mapped_column(Numeric(18, 4), nullable=True)  # computed: qty * unit_cost
    journal_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("journals.id"), nullable=True)

    created_by: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
