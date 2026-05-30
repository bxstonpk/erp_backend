from database.imports import *

class ApprovalRequest(Base):
    __tablename__ = "approval_requests"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    policy_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("approval_policies.id"))

    doc_module: Mapped[str] = mapped_column(String(30), index=True)  # PO, SO, INV, PAY, GRN, JV, ADJ
    doc_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), index=True)  # ID of the document
    doc_no: Mapped[str] = mapped_column(String(50))
    doc_amount: Mapped[Optional[float]] = mapped_column(Numeric(18, 2), nullable=True)
    doc_date: Mapped[Optional[datetime]] = mapped_column(Date, nullable=True)

    status: Mapped[str] = mapped_column(String(20), default="PENDING", index=True)  # PENDING, APPROVED, REJECTED, CANCELLED, RECALLED
    current_step: Mapped[int] = mapped_column(Integer, default=1)
    total_steps: Mapped[int] = mapped_column(Integer)

    requested_by: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"))
    requested_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    approved_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
    rejected_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
    sla_deadline: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
    doc_snapshot: Mapped[Optional[dict]] = mapped_column(JSONB, nullable=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)

    actions = relationship("ApprovalAction", back_populates="request")
