select
    count(*) records
from
	applet_activity
where
	applet_activity.id {{.ne}} ""
