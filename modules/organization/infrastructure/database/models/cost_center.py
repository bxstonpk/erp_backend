from database.imports import *

class CostCenter(Base):
    __tablename__ = "cost_centers"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))

    parent_cost_center_id: Mapped[Optional[UUID]] = mapped_column(UUID, ForeignKey("cost_centers.id"), nullable=True)

    code: Mapped[str] = mapped_column(String, index=True)
    name: Mapped[str] = mapped_column(String, index=True)

    cost_center_type: Mapped[Optional[str]] = mapped_column(String, nullable=True)

    status: Mapped[Optional[str]] = mapped_column(String, nullable=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    company = relationship("Company", back_populates="cost_centers")
    parent_cost_center = relationship("CostCenter", remote_side=[id], back_populates="child_cost_centers")
    child_cost_centers = relationship("CostCenter", back_populates="parent_cost_center")
    project_sites = relationship("ProjectSite", back_populates="cost_center")