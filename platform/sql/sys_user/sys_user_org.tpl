select
sys_org.name name,
sys_org.id id
from
    sys_org
where is_delete {{.ne}} 1
{{if ne .oids ""}}
    and sys_org.id in ({{.oids}})
{{else}}
    limit 0 offset 0
{{end}}