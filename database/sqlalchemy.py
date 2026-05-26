from sqlalchemy import create_engine
from core.config import settings

sql_uri = f"postgresql://{settings.DB_USER}:{settings.DB_PASSWORD}@{settings.DB_HOST}:{settings.DB_PORT}/{settings.DB_NAME}"

engine = create_engine(sql_uri, echo=True)