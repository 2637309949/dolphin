select
    count(*) records
from
	example_applet_activity
where
	example_applet_activity.id {{.ne}} ""
