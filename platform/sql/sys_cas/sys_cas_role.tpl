select
    name,
    code
from sys_role
left join sys_role_user
on sys_role.id = sys_role_user.role_id and sys_role_user.user_id = {{.user_id}}