from database.imports import *

class JournalLine(Base):
    __tablename__ = "journal_lines"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    journal_entry_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("journal_entries.id"))
    account_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("accounts.id"))

    description: Mapped[Optional[str]] = mapped_column(String)
    debit: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    credit: Mapped[float] = mapped_column(Numeric(18, 2), default=0)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    journal_entry = relationship("JournalEntry", back_populates="lines")
    account = relationship("Account", back_populates="journal_lines")
