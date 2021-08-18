select
    count(*) records
from
	sys_role
where
	sys_role.is_delete {{.ne}} 1
