from database.imports import *

class SalesOrderLine(Base):
    __tablename__ = "sales_order_lines"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    so_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("sales_orders.id"), index=True)
    line_no: Mapped[int] = mapped_column(Integer)
    product_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("products.id"))
    description: Mapped[Optional[str]] = mapped_column(String(500), nullable=True)
    uom_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("units_of_measure.id"))

    qty_ordered: Mapped[float] = mapped_column(Numeric(18, 4))
    qty_shipped: Mapped[float] = mapped_column(Numeric(18, 4), default=0)
    unit_price: Mapped[float] = mapped_column(Numeric(18, 4))
    discount_pct: Mapped[float] = mapped_column(Numeric(5, 2), default=0)
    line_amount: Mapped[float] = mapped_column(Numeric(18, 2))
    tax_type: Mapped[str] = mapped_column(String(10), default="VAT7")  # NO_TAX, VAT7, EXEMPT
    tax_amount: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    warehouse_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("warehouses.id"), nullable=True)

    sales_order = relationship("SalesOrder", back_populates="lines")
