from database.imports import *

class VatRate(Base):
    __tablename__ = "vat_rates"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)

    code: Mapped[str] = mapped_column(String(20), unique=True, index=True)  # VAT7, VAT0, EXEMPT, OUT_OF_SCOPE
    name_th: Mapped[str] = mapped_column(String(100))
    name_en: Mapped[Optional[str]] = mapped_column(String(100), nullable=True)
    rate: Mapped[float] = mapped_column(Numeric(5, 2))  # percentage e.g. 7.00
    vat_category: Mapped[str] = mapped_column(String(15))  # STANDARD, ZERO_RATED, EXEMPT, OUT_OF_SCOPE
    effective_from: Mapped[datetime] = mapped_column(Date)
    effective_until: Mapped[Optional[datetime]] = mapped_column(Date, nullable=True)  # NULL = still in use

    output_vat_account_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("accounts.id"), nullable=True)
    input_vat_account_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("accounts.id"), nullable=True)
    is_active: Mapped[bool] = mapped_column(Boolean, default=True)
