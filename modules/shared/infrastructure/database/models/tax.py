from database.imports import *

class Tax(Base):
    __tablename__ = "taxes"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    code: Mapped[str] = mapped_column(String, unique=True, index=True) # e.g. VAT7, WHT3
    name: Mapped[str] = mapped_column(String, index=True)
    rate: Mapped[float] = mapped_column(Numeric(5, 2)) # percentage e.g. 7.00

    status: Mapped[str] = mapped_column(String, index=True) # active, inactive

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)
