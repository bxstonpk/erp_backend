from database.imports import *

class ApprovalPolicy(Base):
    __tablename__ = "approval_policies"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)

    doc_module: Mapped[str] = mapped_column(String(30), index=True)  # PO, SO, INV, PAY, GRN, JV, ADJ
    policy_name: Mapped[str] = mapped_column(String(255))
    description: Mapped[Optional[str]] = mapped_column(Text, nullable=True)
    conditions: Mapped[Optional[dict]] = mapped_column(JSONB, nullable=True)  # amount_gte, amount_lte, department_ids ...

    total_steps: Mapped[int] = mapped_column(Integer, default=1)
    on_reject: Mapped[str] = mapped_column(String(30), default="RETURN_TO_DRAFT")  # RETURN_TO_DRAFT, CANCEL, RETURN_TO_PREV_STEP
    sla_hours: Mapped[Optional[int]] = mapped_column(Integer, default=48)
    sla_action: Mapped[str] = mapped_column(String(20), default="NOTIFY_ONLY")  # NOTIFY_ONLY, AUTO_ESCALATE, AUTO_APPROVE
    priority: Mapped[int] = mapped_column(Integer, default=10)
    is_active: Mapped[bool] = mapped_column(Boolean, default=True)

    created_by: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)

    steps = relationship("ApprovalStep", back_populates="policy")
