select
    sys_data_permission_detail.*
from
    `sys_data_permission`, `sys_data_permission_detail`,`sys_user_role`
where
    sys_data_permission.id = sys_data_permission_detail.data_permission_id
    and sys_user_role.role_id = sys_data_permission_detail.role_id
    and sys_data_permission_detail.del_flag = 0
    and sys_data_permission.code = "{{.code}}"
    and sys_user_role.user_id = "{{.user_id}}"