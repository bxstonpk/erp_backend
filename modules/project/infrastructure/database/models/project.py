from database.imports import *

class Project(Base):
    __tablename__ = "projects"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)

    business_unit_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("business_units.id"))
    fiscal_period_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("fiscal_periods.id"))

    code: Mapped[str] = mapped_column(String, unique=True, index=True)
    name: Mapped[str] = mapped_column(String, index=True)
    description: Mapped[Optional[str]] = mapped_column(String)

    start_date: Mapped[datetime] = mapped_column(DateTime(timezone=True))
    end_date: Mapped[datetime] = mapped_column(DateTime(timezone=True))

    status: Mapped[str] = mapped_column(String, index=True) # active, inactive, completed, on_hold

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    business_unit = relationship("BusinessUnit", back_populates="projects")
    fiscal_period = relationship("FiscalPeriod", back_populates="projects")
    project_sites = relationship("ProjectSite", back_populates="project")
