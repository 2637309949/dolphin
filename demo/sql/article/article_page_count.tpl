-- Code generated by dol build. Only Generate by tools if not existed.
select
    count(*) records
from
	article
where
	article.id {{.ne}} ""
	and
	article.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
