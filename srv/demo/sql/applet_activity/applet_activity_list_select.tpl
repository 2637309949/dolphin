select
    example_applet_activity.id
from
	example_applet_activity
where
	example_applet_activity.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
