select
    sys_optionset.id,
	sys_optionset.name,
	sys_optionset.code,
	sys_optionset.value,
	sys_optionset.remark
from
	sys_optionset
where
	sys_optionset.id {{.ne}} ""
	and
	sys_optionset.is_delete {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
LIMIT {{.size}} OFFSET {{.offset}}
