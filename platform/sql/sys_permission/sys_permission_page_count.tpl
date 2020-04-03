select
    count(*) records
from
	sys_permission
where
	sys_permission.id {{.ne}} ""
