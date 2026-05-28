from database.imports import *

class Currency(Base):
    __tablename__ = "currencies"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    code: Mapped[str] = mapped_column(String, unique=True, index=True) # ISO 4217 e.g. THB, USD, EUR
    name: Mapped[str] = mapped_column(String, index=True)
    symbol: Mapped[Optional[str]] = mapped_column(String)
    decimal_places: Mapped[int] = mapped_column(Integer, default=2)

    status: Mapped[str] = mapped_column(String, index=True) # active, inactive

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    organization_groups = relationship("OrganizationGroup", back_populates="base_currency")
    companies = relationship("Company", back_populates="base_currency")
