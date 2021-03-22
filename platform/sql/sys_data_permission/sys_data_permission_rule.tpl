select
    sys_data_permission_detail.*
from
    `sys_data_permission`, `sys_data_permission_detail`,`sys_role_user`
where
    sys_data_permission.id = sys_data_permission_detail.data_permission_id
    and sys_role_user.role_id = sys_data_permission_detail.role_id
    and sys_data_permission_detail.is_delete = 0
    and sys_data_permission.code = "{{.code}}"
    and sys_role_user.user_id = "{{.user_id}}"