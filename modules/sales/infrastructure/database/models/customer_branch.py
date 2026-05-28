from database.imports import *

class CustomerBranch(Base):
    __tablename__ = "customer_branches"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    customer_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("customers.id"))
    name: Mapped[str] = mapped_column(String, index=True)
    address_id: Mapped[Optional[UUID]] = mapped_column(UUID, ForeignKey("addresses.id"), nullable=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    customer = relationship("Customer", back_populates="branches")