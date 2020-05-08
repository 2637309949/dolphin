select
    wechat_activity.id
from
	wechat_activity
where
	wechat_activity.id {{.ne}} ""
	and
	wechat_activity.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
