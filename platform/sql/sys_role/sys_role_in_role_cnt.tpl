select
    count(distinct(role_id))
from sys_role_user
where
role_id in (select id from sys_role where code in ({{.roles}})) and user_id = "{{.uid}}" and del_flag != 1