select
    count(*) records
from
	dolphin_wechat_activity
where
	dolphin_wechat_activity.id {{.ne}} ""
