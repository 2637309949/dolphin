select
    count(*) records
from
	sys_attachment
where
	sys_attachment.id {{.ne}} ""
	and
	sys_attachment.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
