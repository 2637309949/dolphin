select
    count(*) records
from
	sys_domain
where
	sys_domain.id {{.ne}} ""
	and
	sys_domain.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}