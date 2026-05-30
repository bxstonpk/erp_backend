from database.imports import *

class JournalLine(Base):
    __tablename__ = "journal_lines"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    journal_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("journals.id"), index=True)
    line_no: Mapped[int] = mapped_column(Integer)
    account_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("accounts.id"))
    department_id: Mapped[Optional[UUID]] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("departments.id"), nullable=True)

    debit: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    credit: Mapped[float] = mapped_column(Numeric(18, 2), default=0)
    description: Mapped[Optional[str]] = mapped_column(String(500), nullable=True)

    journal = relationship("Journal", back_populates="lines")
    account = relationship("Account", back_populates="journal_lines")
