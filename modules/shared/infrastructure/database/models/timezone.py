from database.imports import *

class Timezone(Base):
    __tablename__ = "timezones"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    code: Mapped[str] = mapped_column(String, unique=True, index=True) # IANA name e.g. Asia/Bangkok
    name: Mapped[str] = mapped_column(String, index=True)
    utc_offset: Mapped[str] = mapped_column(String) # e.g. +07:00

    status: Mapped[str] = mapped_column(String, index=True) # active, inactive

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    organization_groups = relationship("OrganizationGroup", back_populates="timezone")
    companies = relationship("Company", back_populates="timezone")
