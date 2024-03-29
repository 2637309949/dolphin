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
    and is_delete = 0

{{if ne .name ""}}
    and sys_org.name = "{{.name}}"
{{end}}

order by `order`
