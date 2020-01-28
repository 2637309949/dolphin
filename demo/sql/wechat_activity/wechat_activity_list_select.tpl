select
    dolphin_wechat_activity.id,
	dolphin_wechat_activity.title
from
	dolphin_wechat_activity
where
	dolphin_wechat_activity.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
