select
    sys_role.id,
	sys_role.name,
	sys_role.code,
	sys_role.status,
	sys_role.app_index,
	sys_role.admin_index
from
	sys_role
where
	sys_role.id {{.ne}} ""
	and
	sys_role.is_delete {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
LIMIT {{.size}} OFFSET {{.offset}}
