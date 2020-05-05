select
sys_org.name name,
sys_org.id id
from
    sys_org
where del_flag=0
{{if ne .oids ""}}
    and sys_org.id in ({{.oids}})
{{else}}
    limit 0 offset 0
{{end}}