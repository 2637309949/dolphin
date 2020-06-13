select
    sys_user_template.id
from
	sys_user_template
where
	sys_user_template.id {{.ne}} ""
	and
	sys_user_template.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
LIMIT {{.size}} OFFSET {{.offset}}
