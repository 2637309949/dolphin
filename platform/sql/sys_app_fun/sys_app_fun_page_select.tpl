select
    sys_app_fun.id
from
	sys_app_fun
where
	sys_app_fun.id {{.ne}} ""
	and
	sys_app_fun.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
LIMIT {{.size}} OFFSET {{.offset}}
