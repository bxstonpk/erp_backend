from database.imports import *

class GoodsReceiptLine(Base):
    __tablename__ = "goods_receipt_lines"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    grn_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("goods_receipts.id"), index=True)
    po_line_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("purchase_order_lines.id"), nullable=True)
    product_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("products.id"))
    warehouse_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("warehouses.id"))
    location_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("warehouse_locations.id"), nullable=True)
    uom_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("units_of_measure.id"))

    qty_received: Mapped[float] = mapped_column(Numeric(18, 4))
    unit_cost: Mapped[float] = mapped_column(Numeric(18, 4))

    goods_receipt = relationship("GoodsReceipt", back_populates="lines")
