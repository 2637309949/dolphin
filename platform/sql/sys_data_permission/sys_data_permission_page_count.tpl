select
    count(*) records
from
	sys_data_permission
where
	sys_data_permission.is_delete {{.ne}} 1
