select
    example_activity.id
from
	example_activity
where
	example_activity.id {{.lt }}{{.gt}} ""
LIMIT {{.size}} OFFSET {{.offset}}
