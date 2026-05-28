from database.imports import *

class Employee(Base):
    __tablename__ = "employees"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)

    code: Mapped[str] = mapped_column(String, unique=True, index=True)
    first_name: Mapped[str] = mapped_column(String, index=True)
    last_name: Mapped[str] = mapped_column(String, index=True)

    email: Mapped[Optional[str]] = mapped_column(String, unique=True, index=True)
    phone: Mapped[Optional[str]] = mapped_column(String, index=True)
    position: Mapped[Optional[str]] = mapped_column(String, index=True)

    status: Mapped[str] = mapped_column(String, index=True) # active, inactive, terminated

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    company_branches = relationship("CompanyBranch", back_populates="manager")
    departments = relationship("Department", back_populates="manager")
