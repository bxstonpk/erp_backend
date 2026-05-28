from database.imports import *

class OrganizationGroup(Base):
    __tablename__ = "organization_groups"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    code: Mapped[str] = mapped_column(String, unique=True, index=True)
    name: Mapped[str] = mapped_column(String, index=True)
    legal_name: Mapped[str] = mapped_column(String, index=True)
    tax_id: Mapped[str] = mapped_column(String, unique=True, index=True) # Tax format # #### #####
    description: Mapped[Optional[str]] = mapped_column(String)

    base_currency_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("currencies.id"))
    country_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("countries.id"))
    timezone_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("timezones.id"))

    status: Mapped[str] = mapped_column(String, index=True) # active, inactive, pending

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    base_currency = relationship("Currency", back_populates="organization_groups")
    country = relationship("Country", back_populates="organization_groups")
    timezone = relationship("Timezone", back_populates="organization_groups")
    companies = relationship("Company", back_populates="organization_group")