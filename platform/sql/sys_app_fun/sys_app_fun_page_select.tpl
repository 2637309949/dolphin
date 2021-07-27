select
    sys_app_fun.id
from
	sys_app_fun
where
	sys_app_fun.id {{.ne}} ""
	and
	sys_app_fun.is_delete {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
