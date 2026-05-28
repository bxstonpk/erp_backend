from database.imports import *

class PurchaseOrder(Base):
    __tablename__ = "purchase_orders"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))
    vendor_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("vendors.id"))

    order_number: Mapped[str] = mapped_column(String, unique=True, index=True)
    order_date: Mapped[datetime] = mapped_column(DateTime(timezone=True), index=True)

    currency_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("currencies.id"))
    total_amount: Mapped[float] = mapped_column(Numeric(18, 2), default=0)

    status: Mapped[str] = mapped_column(String, index=True) # draft, confirmed, received, cancelled

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    vendor = relationship("Vendor", back_populates="purchase_orders")
    lines = relationship("PurchaseOrderLine", back_populates="purchase_order")
