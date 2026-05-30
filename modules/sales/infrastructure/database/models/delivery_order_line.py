from database.imports import *

class DeliveryOrderLine(Base):
    __tablename__ = "delivery_order_lines"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    do_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("delivery_orders.id"), index=True)
    so_line_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("sales_order_lines.id"), nullable=True)
    product_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("products.id"))
    warehouse_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("warehouses.id"))
    location_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("warehouse_locations.id"), nullable=True)
    uom_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("units_of_measure.id"))

    qty_shipped: Mapped[float] = mapped_column(Numeric(18, 4))
    unit_cost: Mapped[float] = mapped_column(Numeric(18, 4), default=0)

    delivery_order = relationship("DeliveryOrder", back_populates="lines")
