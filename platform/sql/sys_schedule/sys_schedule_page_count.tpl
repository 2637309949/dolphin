-- Code generated by dol build. Only Generate by tools if not existed.
select
    count(*) records
from
	sys_schedule
where
	sys_schedule.is_delete {{.ne}} 1

