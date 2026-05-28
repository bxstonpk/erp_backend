from database.imports import *

class Company(Base):
    __tablename__ = "companies"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    organization_group_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("organization_groups.id"))

    code: Mapped[str] = mapped_column(String, unique=True, index=True)
    name: Mapped[str] = mapped_column(String, index=True)
    legal_name: Mapped[str] = mapped_column(String, index=True)
    tax_id: Mapped[str] = mapped_column(String, unique=True, index=True)
    description: Mapped[Optional[str]] = mapped_column(String)

    registration_number: Mapped[str] = mapped_column(String, unique=True, index=True)

    email: Mapped[str] = mapped_column(String, unique=True, index=True)
    phone: Mapped[Optional[str]] = mapped_column(String, index=True)
    website: Mapped[Optional[str]] = mapped_column(String, index=True)

    logo_url: Mapped[Optional[str]] = mapped_column(String)

    base_currency_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("currencies.id"))
    country_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("countries.id"))
    timezone_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("timezones.id"))

    fiscal_year_start_month: Mapped[str] = mapped_column(String, index=True) # January, February, etc.

    status: Mapped[str] = mapped_column(String, index=True) # active, inactive, pending

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    organization_group = relationship("OrganizationGroup", back_populates="companies")
    base_currency = relationship("Currency", back_populates="companies")
    country = relationship("Country", back_populates="companies")
    timezone = relationship("Timezone", back_populates="companies")
    company_branches = relationship("CompanyBranch", back_populates="company")
    business_units = relationship("BusinessUnit", back_populates="company")
    departments = relationship("Department", back_populates="company")
    cost_centers = relationship("CostCenter", back_populates="company")
    fiscal_periods = relationship("FiscalPeriod", back_populates="company")
    project_sites = relationship("ProjectSite", back_populates="company")
    addresses = relationship("Address", back_populates="company")