from database.imports import *

class CompanyBranch(Base):
    __tablename__ = "company_branches"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))

    branch_type: Mapped[str] = mapped_column(String, index=True) # head office, branch office, warehouse, etc.
    tax_branch_number: Mapped[str] = mapped_column(String, unique=True, index=True)

    phone: Mapped[Optional[str]] = mapped_column(String, index=True)
    email: Mapped[Optional[str]] = mapped_column(String, unique=True, index=True)
    address_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("addresses.id"))
    manager_employee_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("employees.id"))
    is_headquarters: Mapped[str] = mapped_column(String, index=True) # true or false

    status: Mapped[str] = mapped_column(String, index=True) # active, inactive, pending

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=datetime.utc)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=datetime.utc, onupdate=datetime.utc)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    company = relationship("Company", back_populates="company_branches")
    address = relationship("Address", back_populates="company_branches")
    manager = relationship("Employee", back_populates="company_branches")