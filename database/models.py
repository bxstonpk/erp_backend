"""Central model registry.

Importing this module registers every ORM model on ``Base.metadata`` and makes
all cross-module relationships resolvable. Import it once at startup (and from
Alembic's env.py) before calling ``configure_mappers`` or ``create_all``.

The schema mirrors ``accounting_db.sql`` (the system of record) using UUID
primary keys per project convention.
"""

from database.base import Base

from modules.shared.infrastructure.database.models import *  # noqa: F401,F403
from modules.organization.infrastructure.database.models import *  # noqa: F401,F403
from modules.auth.infrastructure.database.models import *  # noqa: F401,F403
from modules.approval.infrastructure.database.models import *  # noqa: F401,F403
from modules.accounting.infrastructure.database.models import *  # noqa: F401,F403
from modules.sales.infrastructure.database.models import *  # noqa: F401,F403
from modules.purchasing.infrastructure.database.models import *  # noqa: F401,F403
from modules.inventory.infrastructure.database.models import *  # noqa: F401,F403
from modules.tax.infrastructure.database.models import *  # noqa: F401,F403

__all__ = ["Base"]
