import pycassa
from pycassa.pool import ConnectionPool
from pycassa.columnfamily import ColumnFamily

def fetch(key):
    pool = ConnectionPool('Keyspace1', ['127.0.0.1:9160'])
    col_fam = ColumnFamily(pool, 'ColumnFamily1')
    col_fam.insert('row_key', {'col_name': 'col_val'})
    return col_fam.get(str(key))

