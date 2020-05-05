select
sys_org.name name,
sys_org.id id
from
    sys_org
{{if ne .oids ""}}
where sys_org.id in ({{.oids}})
{{else}}
limit 0 offset 0
{{end}}