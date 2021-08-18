select
    sys_notification.id
from
	sys_notification
where
	sys_notification.is_delete {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
