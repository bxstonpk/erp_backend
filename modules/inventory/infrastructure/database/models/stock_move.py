from database.imports import *

class StockMove(Base):
    __tablename__ = "stock_moves"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    product_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("products.id"))
    warehouse_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("warehouses.id"))

    move_type: Mapped[str] = mapped_column(String, index=True) # in, out, transfer, adjustment
    quantity: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    move_date: Mapped[datetime] = mapped_column(DateTime(timezone=True), index=True)
    reference: Mapped[Optional[str]] = mapped_column(String, index=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    product = relationship("Product", back_populates="stock_moves")
    warehouse = relationship("Warehouse", back_populates="stock_moves")
