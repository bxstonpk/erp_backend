from database.imports import *

class InventoryBalance(Base):
    __tablename__ = "inventory_balances"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    warehouse_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("warehouses.id"), index=True)
    location_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("warehouse_locations.id"), nullable=True)
    product_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("products.id"), index=True)

    qty_on_hand: Mapped[float] = mapped_column(Numeric(18, 4), default=0)
    qty_reserved: Mapped[float] = mapped_column(Numeric(18, 4), default=0)
    qty_available: Mapped[Optional[float]] = mapped_column(Numeric(18, 4), nullable=True)  # computed: qty_on_hand - qty_reserved
    avg_cost: Mapped[float] = mapped_column(Numeric(18, 4), default=0)
    last_updated: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
