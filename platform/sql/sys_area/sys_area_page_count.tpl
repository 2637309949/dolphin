select
    count(*) records
from
	sys_area
where
	sys_area.id {{.ne}} ""
	and
	sys_area.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}