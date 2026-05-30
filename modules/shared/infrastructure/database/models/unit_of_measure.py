from database.imports import *

class UnitOfMeasure(Base):
    __tablename__ = "units_of_measure"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)

    code: Mapped[str] = mapped_column(String(20), index=True)  # e.g. PCS, KG, Job
    name_th: Mapped[str] = mapped_column(String(100))
