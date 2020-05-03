select
sys_role_user.user_id,
IFNULL(GROUP_CONCAT(sys_role.name), '') role_name,
IFNULL(GROUP_CONCAT(sys_role.id), '') user_role
from
    sys_role_user
left join
    sys_role
on sys_role.id = sys_role_user.role_id
where sys_role_user.user_id in ({{.uids}})
GROUP BY sys_role_user.user_id