select
    count(*) records
from
	sys_domain
where
	sys_domain.is_delete {{.ne}} 1
