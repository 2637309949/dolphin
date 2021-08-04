select
    sys_menu.id,
    sys_menu.parent,
    sys_menu.code,
    sys_menu.hidden,
    sys_menu.icon,
    sys_menu.name,
    sys_menu.`order`,
    sys_menu.url,
    sys_menu.type,
    sys_menu.component
from
	sys_menu
where
	sys_menu.id {{.ne}} ""
{{if ne .name ""}}
    and sys_menu.name like "%{{.name}}%"
{{end}}
{{if ne .code ""}}
    and sys_menu.code like "%{{.code}}%"
{{end}}
	and
	sys_menu.is_delete {{.ne}} 1
order by {{.sort}}
LIMIT {{.size}} OFFSET {{.offset}}
