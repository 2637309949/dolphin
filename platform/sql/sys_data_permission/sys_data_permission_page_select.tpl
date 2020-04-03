select
    sys_data_permission.id
from
	sys_data_permission
where
	sys_data_permission.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
