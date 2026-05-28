from database.imports import *
from datetime import datetime

class ProjectSite(Base):
    __tablename__ = "project_sites"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))

    branch_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("company_branches.id"))
    cost_center_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("cost_centers.id"))

    code: Mapped[str] = mapped_column(String, index=True)
    name: Mapped[str] = mapped_column(String, index=True)

    project_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("projects.id"))
    address_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("addresses.id"))

    start_date: Mapped[datetime] = mapped_column(DateTime(timezone=True))
    end_date: Mapped[datetime] = mapped_column(DateTime(timezone=True))

    status: Mapped[str] = mapped_column(String, index=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    company = relationship("Company", back_populates="project_sites")
    branch = relationship("CompanyBranch", back_populates="project_sites")
    project = relationship("Project", back_populates="project_sites")
    address = relationship("Address", back_populates="project_sites")
    cost_center = relationship("CostCenter", back_populates="project_sites")