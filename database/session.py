from sqlalchemy.orm import sessionmaker
from .sqlalchemy import engine

SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
