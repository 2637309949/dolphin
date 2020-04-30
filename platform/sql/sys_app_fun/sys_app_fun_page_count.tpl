select
    count(*) records
from
	sys_app_fun
where
	sys_app_fun.id {{.ne}} ""
	and
	sys_app_fun.del_flag {{.ne}} 1
