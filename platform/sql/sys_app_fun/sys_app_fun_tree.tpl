select
    sys_app_fun.id,
    sys_app_fun.parent,
    sys_app_fun.code,
    sys_app_fun.hidden,
    sys_app_fun.icon,
    sys_app_fun.name,
    sys_app_fun.order,
    sys_app_fun.url
from
	sys_app_fun
where
	sys_app_fun.id {{.ne}} ""
    and del_flag = 0
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
    {{if ne .name ""}}
    and sys_app_fun.name = "{{.name}}"
    {{end}}

order by `order`
