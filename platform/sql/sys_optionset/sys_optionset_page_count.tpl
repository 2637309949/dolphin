select
    count(*) records
from
	sys_optionset
where
	sys_optionset.id {{.ne}} ""
	and
	sys_optionset.is_delete {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}