from database.imports import *

class ApprovalStep(Base):
    __tablename__ = "approval_steps"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    policy_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("approval_policies.id"), index=True)

    step_no: Mapped[int] = mapped_column(Integer)  # 1, 2, 3 ...
    step_name: Mapped[str] = mapped_column(String(100))  # e.g. Dept Head, CFO, MD
    approver_type: Mapped[str] = mapped_column(String(20), default="ROLE")  # SPECIFIC_USER, ROLE, DEPARTMENT_HEAD, DYNAMIC
    approver_role_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("roles.id"), nullable=True)
    approver_user_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"), nullable=True)
    approval_mode: Mapped[str] = mapped_column(String(20), default="ANY_ONE")  # ANY_ONE, ALL_MUST

    can_skip: Mapped[bool] = mapped_column(Boolean, default=False)
    allow_delegate: Mapped[bool] = mapped_column(Boolean, default=True)
    require_reason_on_approve: Mapped[bool] = mapped_column(Boolean, default=False)
    require_reason_on_reject: Mapped[bool] = mapped_column(Boolean, default=True)

    policy = relationship("ApprovalPolicy", back_populates="steps")
