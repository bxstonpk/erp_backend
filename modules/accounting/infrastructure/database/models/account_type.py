from database.imports import *

class AccountType(Base):
    __tablename__ = "account_types"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)

    code: Mapped[str] = mapped_column(String(20), unique=True, index=True)  # ASSET, LIABILITY, EQUITY ...
    name_th: Mapped[str] = mapped_column(String(100))
    name_en: Mapped[Optional[str]] = mapped_column(String(100), nullable=True)
    normal_balance: Mapped[str] = mapped_column(String(2))  # DR, CR
    fs_section: Mapped[str] = mapped_column(String(2))  # BS, PL, CF

    accounts = relationship("Account", back_populates="account_type")
