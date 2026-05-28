from database.imports import *

class Customer(Base):
    __tablename__ = "customers"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))

    code: Mapped[str] = mapped_column(String, unique=True, index=True)
    name: Mapped[str] = mapped_column(String, index=True)
    legal_name: Mapped[Optional[str]] = mapped_column(String, index=True)
    tax_id: Mapped[Optional[str]] = mapped_column(String, index=True)

    email: Mapped[Optional[str]] = mapped_column(String, index=True)
    phone: Mapped[Optional[str]] = mapped_column(String, index=True)
    address_id: Mapped[Optional[UUID]] = mapped_column(UUID, ForeignKey("addresses.id"), nullable=True)

    currency_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("currencies.id"))
    credit_limit: Mapped[float] = mapped_column(Numeric(18, 2), default=0)

    status: Mapped[str] = mapped_column(String, index=True) # active, inactive

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    sales_orders = relationship("SalesOrder", back_populates="customer")
    branches = relationship("CustomerBranch", back_populates="customer")