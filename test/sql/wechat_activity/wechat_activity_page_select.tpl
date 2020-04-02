select
    wechat_activity.id
from
	wechat_activity
where
	wechat_activity.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
