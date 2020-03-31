select
    sys_attachment.id
from
	sys_attachment
where
	sys_attachment.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
