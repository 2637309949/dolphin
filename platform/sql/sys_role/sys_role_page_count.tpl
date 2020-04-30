select
    count(*) records
from
	sys_role
where
	sys_role.id {{.ne}} ""
	and
	sys_role.del_flag {{.ne}} 1
