from database.imports import *

class DocumentStamp(Base):
    __tablename__ = "document_stamps"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    request_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("approval_requests.id"))

    doc_module: Mapped[str] = mapped_column(String(30), index=True)
    doc_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), index=True)
    doc_no: Mapped[str] = mapped_column(String(50))

    stamp_uuid: Mapped[str] = mapped_column(String(36), unique=True)  # UUID printed as QR on the document
    content_hash: Mapped[str] = mapped_column(String(128))  # SHA-256 of document content at print time
    hash_algorithm: Mapped[str] = mapped_column(String(20), default="SHA-256")
    approvers_json: Mapped[dict] = mapped_column(JSONB)  # snapshot of approver names/positions/fingerprints
    combined_signature: Mapped[Optional[str]] = mapped_column(Text, nullable=True)

    printed_by: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"))
    printed_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    print_count: Mapped[int] = mapped_column(Integer, default=1)

    is_void: Mapped[bool] = mapped_column(Boolean, default=False)
    voided_by: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"), nullable=True)
    voided_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
    void_reason: Mapped[Optional[str]] = mapped_column(Text, nullable=True)
