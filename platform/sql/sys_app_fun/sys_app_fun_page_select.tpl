select
    sys_app_fun.id
from
	sys_app_fun
where
	sys_app_fun.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
