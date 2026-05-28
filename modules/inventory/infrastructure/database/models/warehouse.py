from database.imports import *

class Warehouse(Base):
    __tablename__ = "warehouses"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))

    code: Mapped[str] = mapped_column(String, unique=True, index=True)
    name: Mapped[str] = mapped_column(String, index=True)
    address_id: Mapped[Optional[UUID]] = mapped_column(UUID, ForeignKey("addresses.id"), nullable=True)

    status: Mapped[str] = mapped_column(String, index=True) # active, inactive

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    stock_levels = relationship("StockLevel", back_populates="warehouse")
    stock_moves = relationship("StockMove", back_populates="warehouse")
