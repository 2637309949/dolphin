select
    count(*) records
from
	sys_data_permission
where
	sys_data_permission.id {{.ne}} ""
	and
	sys_data_permission.is_delete {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}