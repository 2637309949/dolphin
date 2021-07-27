select
    sys_data_permission.id
from
	sys_data_permission
where
	sys_data_permission.id {{.ne}} ""
	and
	sys_data_permission.is_delete {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
