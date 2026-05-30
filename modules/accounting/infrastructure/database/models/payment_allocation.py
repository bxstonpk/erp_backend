from database.imports import *

class PaymentAllocation(Base):
    __tablename__ = "payment_allocations"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    payment_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("payments.id"), index=True)
    invoice_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("invoices.id"), index=True)
    allocated_amount: Mapped[float] = mapped_column(Numeric(18, 2))
    allocated_date: Mapped[datetime] = mapped_column(Date)

    payment = relationship("Payment", back_populates="allocations")
