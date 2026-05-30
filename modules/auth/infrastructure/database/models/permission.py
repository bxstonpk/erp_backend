from database.imports import *

class Permission(Base):
    __tablename__ = "permissions"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)

    module: Mapped[str] = mapped_column(String(50), index=True)  # SALES, PURCHASE, INVENTORY, GL ...
    action: Mapped[str] = mapped_column(String(50), index=True)  # VIEW, CREATE, EDIT, DELETE, APPROVE, POST
    description: Mapped[Optional[str]] = mapped_column(String(255), nullable=True)

    roles = relationship("Role", secondary="role_permissions", back_populates="permissions")
