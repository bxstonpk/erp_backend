from database.imports import *

class SalesOrderLine(Base):
    __tablename__ = "sales_order_lines"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    sales_order_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("sales_orders.id"))
    product_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("products.id"))
    tax_id: Mapped[Optional[UUID]] = mapped_column(UUID, ForeignKey("taxes.id"), nullable=True)

    description: Mapped[Optional[str]] = mapped_column(String)
    quantity: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    unit_price: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    line_total: Mapped[float] = mapped_column(Numeric(18, 2), default=0)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    sales_order = relationship("SalesOrder", back_populates="lines")
