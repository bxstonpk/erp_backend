from database.imports import *

class JournalEntry(Base):
    __tablename__ = "journal_entries"

    id: Mapped[UUID] = mapped_column(UUID, primary_key=True, index=True)
    company_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("companies.id"))
    fiscal_period_id: Mapped[UUID] = mapped_column(UUID, ForeignKey("fiscal_periods.id"))

    entry_number: Mapped[str] = mapped_column(String, unique=True, index=True)
    entry_date: Mapped[datetime] = mapped_column(DateTime(timezone=True), index=True)
    description: Mapped[Optional[str]] = mapped_column(String)
    reference: Mapped[Optional[str]] = mapped_column(String, index=True)

    status: Mapped[str] = mapped_column(String, index=True) # draft, posted, reversed

    posted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)
    deleted_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), nullable=True)

    lines = relationship("JournalLine", back_populates="journal_entry")
