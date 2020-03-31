select
    count(*) records
from
	wechat_activity
where
	wechat_activity.id {{.ne}} ""
