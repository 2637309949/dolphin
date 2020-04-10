select
    wechat_activity.id
from
	wechat_activity
where
	wechat_activity.id {{.ne}} ""
	and
	wechat_activity.del_flag {{.ne}} 0
LIMIT {{.size}} OFFSET {{.offset}}
