from database.imports import *

class Account(Base):
    __tablename__ = "accounts"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    company_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("companies.id"), index=True)
    account_type_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("account_types.id"), index=True)
    parent_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("accounts.id"), nullable=True)  # self-ref hierarchy

    code: Mapped[str] = mapped_column(String(20), index=True)
    name_th: Mapped[str] = mapped_column(String(255), index=True)
    name_en: Mapped[Optional[str]] = mapped_column(String(255), nullable=True)
    level: Mapped[int] = mapped_column(Integer, default=1)
    is_header: Mapped[bool] = mapped_column(Boolean, default=False)  # header accounts cannot post
    is_active: Mapped[bool] = mapped_column(Boolean, default=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)

    account_type = relationship("AccountType", back_populates="accounts")
    parent = relationship("Account", remote_side=[id], back_populates="children")
    children = relationship("Account", back_populates="parent")
    journal_lines = relationship("JournalLine", back_populates="account")
