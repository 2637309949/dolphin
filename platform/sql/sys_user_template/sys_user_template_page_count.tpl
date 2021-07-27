select
    count(*) records
from
	sys_user_template
where
	sys_user_template.id {{.ne}} ""
	and
	sys_user_template.is_delete {{.ne}} 1
