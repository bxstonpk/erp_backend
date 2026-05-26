from database.imports import *

class Country(Base):
    __tablename__ = "countries"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    code: Mapped[str] = mapped_column(String, unique=True, index=True)
    name: Mapped[str] = mapped_column(String, index=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=datetime.utc)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=datetime.utc, onupdate=datetime.utc)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    companies = relationship("Company", back_populates="country")