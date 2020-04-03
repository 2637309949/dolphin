select
    count(*) records
from
	sys_domain
where
	sys_domain.id {{.ne}} ""
