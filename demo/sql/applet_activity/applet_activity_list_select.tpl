select
    applet_activity.id
from
	applet_activity
where
	applet_activity.id {{.ne}} ""
	and
	applet_activity.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
