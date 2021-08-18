select
    count(*) records
from
	sys_notification
where
	sys_notification.is_delete {{.ne}} 1
