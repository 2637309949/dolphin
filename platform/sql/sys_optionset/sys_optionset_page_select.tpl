select
    sys_optionset.id
from
	sys_optionset
where
	sys_optionset.id {{.ne}} ""
	and
	sys_optionset.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
