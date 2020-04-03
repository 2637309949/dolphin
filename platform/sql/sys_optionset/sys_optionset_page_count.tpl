select
    count(*) records
from
	sys_optionset
where
	sys_optionset.id {{.ne}} ""
