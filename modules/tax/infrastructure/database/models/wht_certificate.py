from database.imports import *

class WhtCertificate(Base):
    __tablename__ = "wht_certificates"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    branch_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("branches.id"))

    cert_no: Mapped[str] = mapped_column(String(30), index=True)  # WHT certificate number (50 ทวิ)
    cert_date: Mapped[datetime] = mapped_column(Date)
    partner_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("business_partners.id"))
    partner_name: Mapped[str] = mapped_column(String(255))
    partner_tax_id: Mapped[str] = mapped_column(String(20))
    partner_address: Mapped[Optional[str]] = mapped_column(Text, nullable=True)
    payer_name: Mapped[str] = mapped_column(String(255))
    payer_tax_id: Mapped[str] = mapped_column(String(20))
    payer_address: Mapped[Optional[str]] = mapped_column(Text, nullable=True)

    total_income: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    total_wht: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    condition_type: Mapped[str] = mapped_column(String(10), default="EVERY_TIME")  # ONE_TIME, EVERY_TIME, OTHERS
    status: Mapped[str] = mapped_column(String(10), default="DRAFT", index=True)  # DRAFT, ISSUED, CANCELLED

    issued_by: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"), nullable=True)
    issued_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
    e_wht_ref_no: Mapped[Optional[str]] = mapped_column(String(100), nullable=True)
    stamp_uuid: Mapped[Optional[str]] = mapped_column(String(36), nullable=True)
    content_hash: Mapped[Optional[str]] = mapped_column(String(128), nullable=True)
    notes: Mapped[Optional[str]] = mapped_column(Text, nullable=True)

    created_by: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"))
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
