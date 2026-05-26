from database.imports import *

class FiscalPeriod(Base):
    __tablename__ = "fiscal_periods"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))
    
    year: Mapped[int] = mapped_column(Integer, index=True)
    period_number: Mapped[int] = mapped_column(Integer, index=True)

    start_date: Mapped[datetime] = mapped_column(DateTime(timezone=True))
    end_date: Mapped[datetime] = mapped_column(DateTime(timezone=True))

    status: Mapped[str] = mapped_column(String, index=True)

    is_closed: Mapped[bool] = mapped_column(Boolean, default=False)
    closed_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=datetime.utc)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=datetime.utc, onupdate=datetime.utc)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    company = relationship("Company", back_populates="fiscal_periods")
    projects = relationship("Project", back_populates="fiscal_period")