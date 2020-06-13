select
    count(*) records
from
	sys_role
where
	sys_role.id {{.ne}} ""
	and
	sys_role.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}