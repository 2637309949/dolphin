select
    sys_user.id,
	sys_user.name,
	sys_user.nickname,
	sys_user.email,
	sys_user.mobile,
	sys_user.status,
	sys_user.org_id
from
	sys_user
where
	sys_user.id {{.ne}} ""
	and sys_user.del_flag {{.ne}} 1
{{if ne .org_id ""}}
	and sys_user.org_id = "{{.org_id}}"
{{end}}
{{if ne .cn_org_id ""}}
	and sys_user.org_id in ({{.cn_org_id}})
{{end}}
LIMIT {{.size}} OFFSET {{.offset}}
