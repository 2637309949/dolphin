select
    sys_data_permission.id
from
	sys_data_permission
where
	sys_data_permission.id {{.ne}} ""
	and
	sys_data_permission.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
