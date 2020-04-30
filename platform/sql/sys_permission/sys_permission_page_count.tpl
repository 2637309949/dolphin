select
    count(*) records
from
	sys_permission
where
	sys_permission.id {{.ne}} ""
	and
	sys_permission.del_flag {{.ne}} 1
