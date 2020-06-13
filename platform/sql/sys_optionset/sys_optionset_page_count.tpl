select
    count(*) records
from
	sys_optionset
where
	sys_optionset.id {{.ne}} ""
	and
	sys_optionset.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}