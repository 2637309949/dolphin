select
    count(*) records
from
	sys_user
where
	sys_user.id {{.ne}} ""
