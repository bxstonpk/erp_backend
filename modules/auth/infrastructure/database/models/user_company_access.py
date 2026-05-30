from database.imports import *

class UserCompanyAccess(Base):
    __tablename__ = "user_company_access"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    user_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"), index=True)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    branch_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("branches.id"), nullable=True)  # NULL = all branches
    department_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("departments.id"), nullable=True)  # NULL = all departments
    role_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("roles.id"))

    is_default: Mapped[bool] = mapped_column(Boolean, default=False)

    user = relationship("User", back_populates="company_access")
    role = relationship("Role")
