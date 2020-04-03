select
    sys_optionset.id
from
	sys_optionset
where
	sys_optionset.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
