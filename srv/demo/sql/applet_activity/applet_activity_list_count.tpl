select
    count(*) records
from
	dolphin_applet_activity
where
	dolphin_applet_activity.id {{.ne}} ""
