from database.imports import *

class BusinessPartner(Base):
    __tablename__ = "business_partners"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)

    code: Mapped[str] = mapped_column(String(30), index=True)
    name_th: Mapped[str] = mapped_column(String(255), index=True)
    name_en: Mapped[Optional[str]] = mapped_column(String(255), nullable=True)

    partner_type: Mapped[str] = mapped_column(String(20), index=True)  # CUSTOMER, VENDOR (comma-separated set)
    tax_id: Mapped[Optional[str]] = mapped_column(String(20), nullable=True, index=True)
    branch_no: Mapped[Optional[str]] = mapped_column(String(10), nullable=True)  # for tax invoice

    credit_days: Mapped[Optional[int]] = mapped_column(Integer, default=30, nullable=True)
    credit_limit: Mapped[float] = mapped_column(Numeric(18, 2), default=0)

    billing_address: Mapped[Optional[str]] = mapped_column(Text, nullable=True)
    shipping_address: Mapped[Optional[str]] = mapped_column(Text, nullable=True)
    phone: Mapped[Optional[str]] = mapped_column(String(100), nullable=True)
    email: Mapped[Optional[str]] = mapped_column(String(255), nullable=True)

    ar_account_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("accounts.id"), nullable=True)  # default AR account
    ap_account_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("accounts.id"), nullable=True)  # default AP account

    is_active: Mapped[bool] = mapped_column(Boolean, default=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
