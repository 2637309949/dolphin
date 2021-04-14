select
    sys_domain.id
from
	sys_domain
where
	sys_domain.id {{.ne}} ""
	and
	sys_domain.is_delete {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
LIMIT {{.size}} OFFSET {{.offset}}
