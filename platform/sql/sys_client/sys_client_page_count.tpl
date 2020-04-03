select
    count(*) records
from
	sys_client
where
	sys_client.id {{.ne}} ""
