from database.imports import *

class WarehouseLocation(Base):
    __tablename__ = "warehouse_locations"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    warehouse_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("warehouses.id"), index=True)

    code: Mapped[str] = mapped_column(String(30), index=True)
    name: Mapped[str] = mapped_column(String(100))

    warehouse = relationship("Warehouse", back_populates="locations")
