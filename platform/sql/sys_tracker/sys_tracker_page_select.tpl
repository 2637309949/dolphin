select
    sys_tracker.id,
	sys_tracker.user_id,
	sys_tracker.status_code,
	sys_tracker.latency,
	sys_tracker.client_ip,
	sys_tracker.method,
	sys_tracker.path,
	sys_tracker.create_time
from
	sys_tracker
where
	sys_tracker.id {{.ne}} ""
	and sys_tracker.is_delete {{.ne}} 1

{{if ne .sort ""}}
	order by {{.sort}}
{{else}}
	order by sys_tracker.update_time desc
{{end}}LIMIT {{.size}} OFFSET {{.offset}}
