select
    count(*) records
from
	sys_org
where
	sys_org.id {{.ne}} ""
