select
    sys_permission.id
from
	sys_permission
where
	sys_permission.id {{.ne}} ""
	and
	sys_permission.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
LIMIT {{.size}} OFFSET {{.offset}}
