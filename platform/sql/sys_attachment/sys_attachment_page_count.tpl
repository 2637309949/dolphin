select
    count(*) records
from
	sys_attachment
where
	sys_attachment.id {{.ne}} ""
	and
	sys_attachment.is_delete {{.ne}} 1

