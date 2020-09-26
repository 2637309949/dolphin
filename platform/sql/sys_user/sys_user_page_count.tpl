select
    count(*) records
from
	sys_user
where
	sys_user.id {{.ne}} ""
	and
	sys_user.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
{{if ne .org_id ""}}
	and sys_user.org_id = "{{.org_id}}"
{{end}}
{{if ne .mobile ""}}
	and sys_user.mobile like "%{{.mobile}}%"
{{end}}
{{if ne .cn_org_id ""}}
	and sys_user.org_id in ({{.cn_org_id}})
{{end}}
{{if ne .name ""}}
	and sys_user.name like "%{{.name}}%"
{{end}}