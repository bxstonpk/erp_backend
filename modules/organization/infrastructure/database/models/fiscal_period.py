from database.imports import *

class FiscalPeriod(Base):
    __tablename__ = "fiscal_periods"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)

    year: Mapped[int] = mapped_column(Integer, index=True)
    period: Mapped[int] = mapped_column(Integer)  # 1-12

    start_date: Mapped[datetime] = mapped_column(Date)
    end_date: Mapped[datetime] = mapped_column(Date)

    status: Mapped[str] = mapped_column(String(10), default="OPEN", index=True)  # OPEN, CLOSED, LOCKED
    closed_by: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"), nullable=True)
    closed_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)

    company = relationship("Company", back_populates="fiscal_periods")
