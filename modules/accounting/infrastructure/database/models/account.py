from database.imports import *

class Account(Base):
    __tablename__ = "accounts"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))

    parent_account_id: Mapped[Optional[UUID]] = mapped_column(UUID, ForeignKey("accounts.id"), nullable=True)

    code: Mapped[str] = mapped_column(String, index=True)
    name: Mapped[str] = mapped_column(String, index=True)
    account_type: Mapped[str] = mapped_column(String, index=True) # asset, liability, equity, income, expense

    is_active: Mapped[bool] = mapped_column(Boolean, default=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    parent_account = relationship("Account", remote_side=[id], back_populates="child_accounts")
    child_accounts = relationship("Account", back_populates="parent_account")
    journal_lines = relationship("JournalLine", back_populates="account")
