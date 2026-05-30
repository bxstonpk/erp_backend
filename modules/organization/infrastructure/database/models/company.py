from database.imports import *

class Company(Base):
    __tablename__ = "companies"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)

    code: Mapped[str] = mapped_column(String(20), unique=True, index=True)
    name_th: Mapped[str] = mapped_column(String(255), index=True)
    name_en: Mapped[Optional[str]] = mapped_column(String(255), nullable=True)

    tax_id: Mapped[Optional[str]] = mapped_column(String(20), nullable=True, index=True)
    registered_no: Mapped[Optional[str]] = mapped_column(String(50), nullable=True)
    address: Mapped[Optional[str]] = mapped_column(Text, nullable=True)
    phone: Mapped[Optional[str]] = mapped_column(String(50), nullable=True)
    email: Mapped[Optional[str]] = mapped_column(String(100), nullable=True)
    logo_url: Mapped[Optional[str]] = mapped_column(String(500), nullable=True)

    fiscal_year_start: Mapped[int] = mapped_column(Integer, default=1)  # month 1-12
    base_currency: Mapped[str] = mapped_column(String(3), default="THB")

    is_active: Mapped[bool] = mapped_column(Boolean, default=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)

    branches = relationship("Branch", back_populates="company")
    departments = relationship("Department", back_populates="company")
    fiscal_periods = relationship("FiscalPeriod", back_populates="company")
