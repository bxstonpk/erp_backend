from database.imports import *

class UnitOfMeasure(Base):
    __tablename__ = "units_of_measure"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    code: Mapped[str] = mapped_column(String, unique=True, index=True) # e.g. PCS, KG, BOX
    name: Mapped[str] = mapped_column(String, index=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)
