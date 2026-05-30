from database.imports import *

class Currency(Base):
    __tablename__ = "currencies"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)

    code: Mapped[str] = mapped_column(String(3), unique=True, index=True)  # ISO 4217 e.g. THB, USD
    name: Mapped[str] = mapped_column(String(100), index=True)
    symbol: Mapped[Optional[str]] = mapped_column(String(10), nullable=True)
