select
    sys_tag_group.id
from
	sys_tag_group
where
	sys_tag_group.id {{.ne}} ""
	and
	sys_tag_group.is_delete {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
LIMIT {{.size}} OFFSET {{.offset}}
