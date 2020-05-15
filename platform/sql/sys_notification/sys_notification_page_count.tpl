select
    count(*) records
from
	sys_notification
where
	sys_notification.id {{.ne}} ""
	and
	sys_notification.del_flag {{.ne}} 1
