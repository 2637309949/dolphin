select
    count(*) records
from
	sys_app_fun
where
	sys_app_fun.id {{.ne}} ""
	and
	sys_app_fun.is_delete {{.ne}} 1
