select
    dolphin_applet_activity.id
from
	dolphin_applet_activity
where
	dolphin_applet_activity.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
