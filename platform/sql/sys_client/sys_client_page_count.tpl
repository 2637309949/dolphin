select
    count(*) records
from
	sys_client
where
	sys_client.is_delete {{.ne}} 1
