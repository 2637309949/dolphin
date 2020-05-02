select
    sys_client.id
from
	sys_client
where
	sys_client.id {{.ne}} ""
	and
	sys_client.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
