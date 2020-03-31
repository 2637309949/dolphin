select
    count(*) records
from
	sys_attachment
where
	sys_attachment.id {{.ne}} ""
