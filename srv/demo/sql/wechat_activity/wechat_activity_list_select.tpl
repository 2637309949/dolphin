select
    example_wechat_activity.id
from
	example_wechat_activity
where
	example_wechat_activity.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
