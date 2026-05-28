from database.imports import *

class Department(Base):
    __tablename__ = "departments"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))

    parent_department_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("departments.id"))
    code: Mapped[str] = mapped_column(String, unique=True, index=True)
    name: Mapped[str] = mapped_column(String, index=True)

    manager_employee_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("employees.id"))

    status: Mapped[str] = mapped_column(String, index=True) # active, inactive, pending

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    company = relationship("Company", back_populates="departments")
    parent_department = relationship("Department", remote_side=[id], back_populates="child_departments")
    child_departments = relationship("Department", back_populates="parent_department")
    manager = relationship("Employee", back_populates="departments")