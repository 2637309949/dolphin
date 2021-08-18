select
    sys_attachment.id
from
	sys_attachment
where
	sys_attachment.is_delete {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
