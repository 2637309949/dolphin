select
    count(*) records
from
	sys_org
where
	sys_org.id {{.ne}} ""
	and
	sys_org.is_delete {{.ne}} 1
