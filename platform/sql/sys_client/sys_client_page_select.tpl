select
    sys_client.id
from
	sys_client
where
	sys_client.is_delete {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
