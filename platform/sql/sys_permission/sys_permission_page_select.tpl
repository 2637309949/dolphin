select
    sys_permission.id
from
	sys_permission
where
	sys_permission.id {{.ne}} ""
	and
	sys_permission.is_delete {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
LIMIT {{.size}} OFFSET {{.offset}}
