import pycassa
from pycassa.pool import ConnectionPool
from pycassa.columnfamily import ColumnFamily
from settings import KEY_SPACE
from settings import DB_URI
from settings import COLUMN_FAMILY

def fetch(key):
    pool = ConnectionPool(KEY_SPACE, [DB_URI])
    col_fam = ColumnFamily(pool, COLUMN_FAMILY)
    col_fam.insert('row_key', {'col_name': 'col_val'})
    return col_fam.get(str(key))