select
    count(*) records
from
	sys_tracker
where
	sys_tracker.is_delete {{.ne}} 1

