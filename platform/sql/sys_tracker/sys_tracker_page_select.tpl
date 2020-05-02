select
    sys_tracker.id,
	sys_tracker.user_id,
	sys_user.name user_name,
	sys_tracker.status_code,
	sys_tracker.latency,
	sys_tracker.client_ip,
	sys_tracker.method,
	sys_tracker.path,
	sys_tracker.create_time
from
	sys_tracker
left join sys_user on sys_user.id = sys_tracker.user_id
where
	sys_tracker.id {{.ne}} ""
	and sys_tracker.domain = "{{.domain}}"
	and sys_tracker.app_name = "{{.app_name}}"
	and sys_tracker.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
