from database.imports import *

class Product(Base):
    __tablename__ = "products"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))

    uom_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("units_of_measure.id"))
    tax_id: Mapped[Optional[UUID]] = mapped_column(UUID, ForeignKey("taxes.id"), nullable=True)

    code: Mapped[str] = mapped_column(String, unique=True, index=True)
    name: Mapped[str] = mapped_column(String, index=True)
    description: Mapped[Optional[str]] = mapped_column(String)
    product_type: Mapped[str] = mapped_column(String, index=True) # good, service

    purchase_price: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    sale_price: Mapped[float] = mapped_column(Numeric(18, 2), default=0)

    is_active: Mapped[bool] = mapped_column(Boolean, default=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    stock_levels = relationship("StockLevel", back_populates="product")
    stock_moves = relationship("StockMove", back_populates="product")
