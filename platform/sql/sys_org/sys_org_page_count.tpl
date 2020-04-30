select
    count(*) records
from
	sys_org
where
	sys_org.id {{.ne}} ""
	and
	sys_org.del_flag {{.ne}} 1
