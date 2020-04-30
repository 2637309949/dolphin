select
    count(*) records
from
	sys_attachment
where
	sys_attachment.id {{.ne}} ""
	and
	sys_attachment.del_flag {{.ne}} 1
