select
    count(*) records
from
	sys_permission
where
	sys_permission.is_delete {{.ne}} 1
