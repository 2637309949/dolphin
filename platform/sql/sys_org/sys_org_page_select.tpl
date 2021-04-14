select
    sys_org.id,
    sys_org.parent,
    sys_org.code,
    sys_org.name,
    sys_org.type,
    sys_org.order
from
	sys_org
where
	sys_org.id {{.ne}} ""
	and
	sys_org.is_delete {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
LIMIT {{.size}} OFFSET {{.offset}}
