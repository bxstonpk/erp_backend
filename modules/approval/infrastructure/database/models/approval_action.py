from database.imports import *

class ApprovalAction(Base):
    __tablename__ = "approval_actions"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    request_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("approval_requests.id"), index=True)

    step_no: Mapped[int] = mapped_column(Integer)
    actor_user_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"))  # who acted
    delegated_from: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"), nullable=True)
    action: Mapped[str] = mapped_column(String(20))  # SUBMIT, APPROVE, REJECT, DELEGATE, RECALL, COMMENT, ESCALATE, AUTO_APPROVE
    reason: Mapped[Optional[str]] = mapped_column(Text, nullable=True)

    digital_signature: Mapped[Optional[str]] = mapped_column(Text, nullable=True)  # Base64 signature
    signature_key_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("user_signatures.id"), nullable=True)
    signed_payload: Mapped[Optional[str]] = mapped_column(Text, nullable=True)  # canonical JSON that was signed
    ip_address: Mapped[Optional[str]] = mapped_column(String(45), nullable=True)
    user_agent: Mapped[Optional[str]] = mapped_column(String(500), nullable=True)

    acted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)

    request = relationship("ApprovalRequest", back_populates="actions")
