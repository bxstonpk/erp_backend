from .base import Base

from sqlalchemy import (
    String,
    Integer,
    Float,
    Numeric,
    Boolean,
    ForeignKey,
    DateTime,
    Table,
    Column,
)

from typing import Optional

from sqlalchemy.dialects.postgresql import (
    UUID,
    JSONB,
)

from sqlalchemy.orm import (
    Mapped,
    mapped_column,
    relationship,
)

from datetime import datetime, timezone


def utcnow() -> datetime:
    return datetime.now(timezone.utc)