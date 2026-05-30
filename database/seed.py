"""Seed the database with the reference data shipped in ``accounting_db.sql``.

Run from the repo root (requires a configured ``.env`` / database):

    python -m database.seed

The source rows use MariaDB integer ids; this project uses UUID primary keys, so
every integer id is mapped to a deterministic UUID via :func:`sid` (uuid5 of a
fixed namespace). That keeps foreign-key references (parent accounts, account
types, role permissions, user access ...) consistent across runs and makes the
seeder idempotent: re-running merges the same UUIDs instead of duplicating rows.
"""
from __future__ import annotations

import uuid
from datetime import datetime, date, timezone

from database.session import SessionLocal
import database.models as models  # noqa: F401  (registers all mappers)
from database import seed_data as sd

from modules.shared.infrastructure.database.models import Currency, UnitOfMeasure
from modules.organization.infrastructure.database.models import Company
from modules.auth.infrastructure.database.models import (
    Permission, Role, User, UserCompanyAccess, role_permissions,
)
from modules.accounting.infrastructure.database.models import AccountType, Account
from modules.tax.infrastructure.database.models import VatRate, WhtType

# Fixed namespace so generated UUIDs are stable across runs/machines.
SEED_NS = uuid.UUID("a1f0c0de-5eed-4000-8000-0123456789ab")


def sid(table: str, key) -> uuid.UUID:
    """Deterministic UUID for a source (table, integer/string id)."""
    return uuid.uuid5(SEED_NS, f"{table}:{key}")


def _dt(value):
    if value is None:
        return None
    return datetime.strptime(value, "%Y-%m-%d %H:%M:%S").replace(tzinfo=timezone.utc)


def _date(value):
    if value is None:
        return None
    return date.fromisoformat(value)


def _bool(value) -> bool:
    return bool(value)


def seed(session) -> None:
    # ---- currencies (PK = code, no integer id) -------------------------------
    for r in sd.CURRENCIES:
        session.merge(Currency(id=sid("currencies", r["code"]), code=r["code"], name=r["name"], symbol=r["symbol"]))

    # ---- companies -----------------------------------------------------------
    for r in sd.COMPANIES:
        session.merge(Company(
            id=sid("companies", r["id"]), code=r["code"], name_th=r["name_th"], name_en=r["name_en"],
            tax_id=r["tax_id"], registered_no=r["registered_no"], address=r["address"], phone=r["phone"],
            email=r["email"], logo_url=r["logo_url"], fiscal_year_start=r["fiscal_year_start"],
            base_currency=r["base_currency"], is_active=_bool(r["is_active"]),
            created_at=_dt(r["created_at"]), updated_at=_dt(r["updated_at"]),
        ))

    # ---- units of measure ----------------------------------------------------
    for r in sd.UNITS_OF_MEASURE:
        session.merge(UnitOfMeasure(
            id=sid("units_of_measure", r["id"]), company_id=sid("companies", r["company_id"]),
            code=r["code"], name_th=r["name_th"],
        ))

    # ---- account types -------------------------------------------------------
    for r in sd.ACCOUNT_TYPES:
        session.merge(AccountType(
            id=sid("account_types", r["id"]), code=r["code"], name_th=r["name_th"], name_en=r["name_en"],
            normal_balance=r["normal_balance"], fs_section=r["fs_section"],
        ))

    # ---- chart of accounts (parent_id self-ref) ------------------------------
    for r in sd.ACCOUNTS:
        session.merge(Account(
            id=sid("accounts", r["id"]), company_id=sid("companies", r["company_id"]),
            account_type_id=sid("account_types", r["account_type_id"]),
            parent_id=sid("accounts", r["parent_id"]) if r["parent_id"] is not None else None,
            code=r["code"], name_th=r["name_th"], name_en=r["name_en"], level=r["level"],
            is_header=_bool(r["is_header"]), is_active=_bool(r["is_active"]), created_at=_dt(r["created_at"]),
        ))

    # ---- permissions ---------------------------------------------------------
    for r in sd.PERMISSIONS:
        session.merge(Permission(
            id=sid("permissions", r["id"]), module=r["module"], action=r["action"], description=r["description"],
        ))

    # ---- roles ---------------------------------------------------------------
    for r in sd.ROLES:
        session.merge(Role(
            id=sid("roles", r["id"]), company_id=sid("companies", r["company_id"]), code=r["code"],
            name=r["name"], description=r["description"], is_system=_bool(r["is_system"]),
            created_at=_dt(r["created_at"]),
        ))

    # ---- users ---------------------------------------------------------------
    for r in sd.USERS:
        session.merge(User(
            id=sid("users", r["id"]), username=r["username"], email=r["email"], password_hash=r["password_hash"],
            full_name=r["full_name"], employee_code=r["employee_code"], is_active=_bool(r["is_active"]),
            last_login_at=_dt(r["last_login_at"]), created_at=_dt(r["created_at"]), updated_at=_dt(r["updated_at"]),
        ))

    # ---- user company access -------------------------------------------------
    for r in sd.USER_COMPANY_ACCESS:
        session.merge(UserCompanyAccess(
            id=sid("user_company_access", r["id"]), user_id=sid("users", r["user_id"]),
            company_id=sid("companies", r["company_id"]),
            branch_id=sid("branches", r["branch_id"]) if r["branch_id"] is not None else None,
            department_id=sid("departments", r["department_id"]) if r["department_id"] is not None else None,
            role_id=sid("roles", r["role_id"]), is_default=_bool(r["is_default"]),
        ))

    # ---- vat rates -----------------------------------------------------------
    for r in sd.VAT_RATES:
        session.merge(VatRate(
            id=sid("vat_rates", r["id"]), code=r["code"], name_th=r["name_th"], name_en=r["name_en"],
            rate=r["rate"], vat_category=r["vat_category"], effective_from=_date(r["effective_from"]),
            effective_until=_date(r["effective_until"]),
            output_vat_account_id=sid("accounts", r["output_vat_account_id"]) if r["output_vat_account_id"] is not None else None,
            input_vat_account_id=sid("accounts", r["input_vat_account_id"]) if r["input_vat_account_id"] is not None else None,
            is_active=_bool(r["is_active"]),
        ))

    # ---- wht types -----------------------------------------------------------
    for r in sd.WHT_TYPES:
        session.merge(WhtType(
            id=sid("wht_types", r["id"]), code=r["code"], section=r["section"], description_th=r["description_th"],
            description_en=r["description_en"], default_rate=r["default_rate"], reduced_rate=r["reduced_rate"],
            form_type=r["form_type"], is_active=_bool(r["is_active"]),
        ))

    session.flush()

    # ---- role <-> permission links (association table) -----------------------
    existing = {tuple(row) for row in session.execute(role_permissions.select()).all()}
    for r in sd.ROLE_PERMISSIONS:
        key = (sid("roles", r["role_id"]), sid("permissions", r["permission_id"]))
        if key not in existing:
            session.execute(role_permissions.insert().values(role_id=key[0], permission_id=key[1]))


def main() -> None:
    session = SessionLocal()
    try:
        seed(session)
        session.commit()
        print("Seed complete.")
    except Exception:
        session.rollback()
        raise
    finally:
        session.close()


if __name__ == "__main__":
    main()
