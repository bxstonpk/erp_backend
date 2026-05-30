from database.imports import *

class Product(Base):
    __tablename__ = "products"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    category_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("product_categories.id"), nullable=True)

    code: Mapped[str] = mapped_column(String(50), index=True)
    name_th: Mapped[str] = mapped_column(String(255), index=True)
    name_en: Mapped[Optional[str]] = mapped_column(String(255), nullable=True)
    product_type: Mapped[str] = mapped_column(String(20), default="PRODUCT")  # PRODUCT, SERVICE, RAW_MATERIAL, FINISHED_GOOD
    base_uom_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("units_of_measure.id"))

    standard_cost: Mapped[Optional[float]] = mapped_column(Numeric(18, 4), default=0)
    list_price: Mapped[Optional[float]] = mapped_column(Numeric(18, 4), default=0)
    tax_type: Mapped[str] = mapped_column(String(10), default="VAT7")  # NO_TAX, VAT7, EXEMPT
    is_inventory: Mapped[bool] = mapped_column(Boolean, default=True)

    inventory_account_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("accounts.id"), nullable=True)
    cogs_account_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("accounts.id"), nullable=True)
    sales_account_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("accounts.id"), nullable=True)
    purchase_account_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("accounts.id"), nullable=True)

    is_active: Mapped[bool] = mapped_column(Boolean, default=True)
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)

    category = relationship("ProductCategory", back_populates="products")
