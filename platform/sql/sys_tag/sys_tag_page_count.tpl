select
    count(*) records
from
	sys_tag
where
	sys_tag.id {{.ne}} ""
	and
	sys_tag.is_delete {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}