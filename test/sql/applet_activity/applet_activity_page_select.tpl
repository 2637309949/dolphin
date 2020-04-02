select
    applet_activity.id
from
	applet_activity
where
	applet_activity.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
