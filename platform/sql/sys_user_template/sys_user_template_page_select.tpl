select
    sys_user_template.id,
	sys_user_template.name,
	sys_user_template.type,
	sys_user_template.default
from
	sys_user_template
where
	sys_user_template.id {{.ne}} ""
	and
	sys_user_template.is_delete {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
LIMIT {{.size}} OFFSET {{.offset}}
