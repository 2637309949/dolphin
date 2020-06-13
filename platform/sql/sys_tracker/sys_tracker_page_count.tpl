select
    count(*) records
from
	sys_tracker
where
	sys_tracker.id {{.ne}} ""
	and
	sys_tracker.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}