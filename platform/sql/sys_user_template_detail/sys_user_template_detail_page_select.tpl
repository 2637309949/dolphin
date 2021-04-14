-- Code generated by dol build. Only Generate by tools if not existed.
select
    sys_user_template_detail.id,
	sys_user_template_detail.name,
	sys_user_template_detail.value,
	sys_user_template_detail.type,
	sys_user_template_detail.content
from
	sys_user_template_detail
where
	sys_user_template_detail.id {{.ne}} ""
	and
	sys_user_template_detail.is_delete {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
{{if ne .temp_id ""}}
	and sys_user_template_detail.temp_id="{{.temp_id}}"
{{end}}
	LIMIT {{.size}} OFFSET {{.offset}}
