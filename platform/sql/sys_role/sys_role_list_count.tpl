select
    count(*) records
from
	sys_role
where
	sys_role.id {{.ne}} ""
