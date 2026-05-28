from database.imports import *

class Address(Base):
    __tablename__ = "addresses"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)

    line1: Mapped[str] = mapped_column(String, index=True)
    line2: Mapped[str] = mapped_column(String, index=True, nullable=True)

    subdistrict: Mapped[str] = mapped_column(String, index=True)
    district: Mapped[str] = mapped_column(String, index=True)
    province: Mapped[str] = mapped_column(String, index=True)
    postal_code: Mapped[str] = mapped_column(String, index=True)

    country_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("countries.id"))

    latitude: Mapped[float] = mapped_column(Float, nullable=True)
    longitude: Mapped[float] = mapped_column(Float, nullable=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))
    company = relationship("Company", back_populates="addresses")
    country = relationship("Country", back_populates="addresses")
    company_branches = relationship("CompanyBranch", back_populates="address")
    project_sites = relationship("ProjectSite", back_populates="address")