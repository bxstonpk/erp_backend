from database.imports import *

class WhtType(Base):
    __tablename__ = "wht_types"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)

    code: Mapped[str] = mapped_column(String(20), unique=True, index=True)
    section: Mapped[str] = mapped_column(String(20))  # e.g. 40(1), 40(2), 40(8)
    description_th: Mapped[str] = mapped_column(String(255))
    description_en: Mapped[Optional[str]] = mapped_column(String(255), nullable=True)
    default_rate: Mapped[float] = mapped_column(Numeric(5, 2))  # normal rate %
    reduced_rate: Mapped[Optional[float]] = mapped_column(Numeric(5, 2), nullable=True)
    form_type: Mapped[Optional[str]] = mapped_column(String(5), nullable=True)  # ND1, ND3, ND53
    is_active: Mapped[bool] = mapped_column(Boolean, default=True)

    transactions = relationship("WhtTransaction", back_populates="wht_type")
