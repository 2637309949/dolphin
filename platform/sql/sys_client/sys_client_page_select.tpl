select
    sys_client.id
from
	sys_client
where
	sys_client.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
