from database.imports import *

class BusinessUnit(Base):
    __tablename__ = "business_units"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))

    code: Mapped[str] = mapped_column(String, unique=True, index=True)
    name: Mapped[str] = mapped_column(String, index=True)

    description: Mapped[Optional[str]] = mapped_column(String)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=datetime.utc)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=datetime.utc, onupdate=datetime.utc)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    company = relationship("Company", back_populates="business_units")
    projects = relationship("Project", back_populates="business_unit")