from database.imports import *

class ApprovalNotification(Base):
    __tablename__ = "approval_notifications"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    request_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("approval_requests.id"), index=True)
    recipient_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"))

    channel: Mapped[str] = mapped_column(String(10), default="IN_APP")  # EMAIL, LINE, IN_APP, SMS
    noti_type: Mapped[str] = mapped_column(String(30))  # APPROVAL_REQUESTED, APPROVED, REJECTED, SLA_WARNING, DELEGATED, RECALLED
    sent_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
    read_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
    status: Mapped[str] = mapped_column(String(10), default="PENDING")  # PENDING, SENT, FAILED, READ
    error_msg: Mapped[Optional[str]] = mapped_column(Text, nullable=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
