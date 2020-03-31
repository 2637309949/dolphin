select
    sys_user_template.id
from
	sys_user_template
where
	sys_user_template.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
