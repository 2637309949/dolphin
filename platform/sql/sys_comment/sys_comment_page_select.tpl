-- Code generated by dol build. Only Generate by tools if not existed.
select
    sys_comment.id
from
	sys_comment
where
	sys_comment.id {{.ne}} ""
	and
	sys_comment.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
	LIMIT {{.size}} OFFSET {{.offset}}
