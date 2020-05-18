select
    sys_notification.id
from
	sys_notification
where
	sys_notification.id {{.ne}} ""
	and
	sys_notification.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}