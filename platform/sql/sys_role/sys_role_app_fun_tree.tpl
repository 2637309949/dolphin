select
    distinct
    sys_app_fun.id,
    sys_app_fun.parent,
    sys_app_fun.code,
    sys_app_fun.hidden,
    sys_app_fun.icon,
    sys_app_fun.name,
    sys_app_fun.order,
    sys_app_fun.url
from
	sys_app_fun{{if ne .is_admin true}}, sys_role_app_fun{{end}}
where
	sys_app_fun.id {{.ne}} ""
{{if ne .is_admin true}}
    and sys_app_fun.id = sys_role_app_fun.app_fun_id
    and sys_role_app_fun.role_id = "{{.role_id}}"
{{end}}
    and sys_app_fun.is_delete = 0

{{if ne .name ""}}
    and sys_app_fun.name = "{{.name}}"
{{end}}

order by `order`
