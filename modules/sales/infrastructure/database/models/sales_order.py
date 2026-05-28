from database.imports import *

class SalesOrder(Base):
    __tablename__ = "sales_orders"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))
    customer_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("customers.id"))

    order_number: Mapped[str] = mapped_column(String, unique=True, index=True)
    order_date: Mapped[datetime] = mapped_column(DateTime(timezone=True), index=True)

    tax_id: Mapped[Optional[UUID]] = mapped_column(UUID, ForeignKey("taxes.id"), nullable=True)

    currency_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("currencies.id"))
    total_amount: Mapped[float] = mapped_column(Numeric(18, 2), default=0)

    status: Mapped[str] = mapped_column(String, index=True) # draft, confirmed, invoiced, cancelled

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    customer = relationship("Customer", back_populates="sales_orders")
    lines = relationship("SalesOrderLine", back_populates="sales_order")
    tax = relationship("Tax", back_populates="sales_orders")