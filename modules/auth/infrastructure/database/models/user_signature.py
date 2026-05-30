from database.imports import *

class UserSignature(Base):
    __tablename__ = "user_signatures"

    id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), primary_key=True, index=True, default=uuid4)
    user_id: Mapped[UUID] = mapped_column(PG_UUID(as_uuid=True), ForeignKey("users.id"), index=True)

    signature_image_url: Mapped[Optional[str]] = mapped_column(String(1000), nullable=True)
    signature_image_b64: Mapped[Optional[str]] = mapped_column(Text, nullable=True)
    display_name_th: Mapped[Optional[str]] = mapped_column(String(255), nullable=True)
    display_name_en: Mapped[Optional[str]] = mapped_column(String(255), nullable=True)
    position_title_th: Mapped[Optional[str]] = mapped_column(String(255), nullable=True)
    position_title_en: Mapped[Optional[str]] = mapped_column(String(255), nullable=True)

    public_key: Mapped[Optional[str]] = mapped_column(Text, nullable=True)  # RSA/ECDSA public key (PEM)
    key_fingerprint: Mapped[Optional[str]] = mapped_column(String(128), nullable=True)  # SHA-256 fingerprint
    key_algorithm: Mapped[str] = mapped_column(String(20), default="RSA_2048")  # RSA_2048, RSA_4096, ECDSA_P256, ECDSA_P384
    key_issued_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
    key_expires_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)
    key_revoked_at: Mapped[Optional[datetime]] = mapped_column(DateTime(timezone=True), nullable=True)

    is_active: Mapped[bool] = mapped_column(Boolean, default=True)

    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)

    user = relationship("User", back_populates="signatures")
