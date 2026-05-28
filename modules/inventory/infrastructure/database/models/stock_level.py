from database.imports import *

class StockLevel(Base):
    __tablename__ = "stock_levels"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    product_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("products.id"))
    warehouse_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("warehouses.id"))

    quantity: Mapped[float] = mapped_column(Numeric(18, 2), default=0)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    product = relationship("Product", back_populates="stock_levels")
    warehouse = relationship("Warehouse", back_populates="stock_levels")
