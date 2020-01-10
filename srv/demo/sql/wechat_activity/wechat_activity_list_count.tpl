select
    count(*) records
from
	example_wechat_activity
where
	example_wechat_activity.id {{.ne}} ""
