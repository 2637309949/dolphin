select
sys_role_user.user_id,
IFNULL(GROUP_CONCAT(sys_role.name), '') role_name,
IFNULL(GROUP_CONCAT(sys_role.id), '') user_role
from
    sys_role_user
left join
    sys_role
on sys_role.id = sys_role_user.role_id
where sys_role_user.is_delete=0
{{if ne .uids ""}}
    and sys_role_user.user_id in ({{.uids}})
{{end}}
GROUP BY sys_role_user.user_id
{{if eq .uids ""}}
    limit 0 offset 0
{{end}}
