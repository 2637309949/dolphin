select
    sys_org.id,
    sys_org.parent,
    sys_org.code,
    sys_org.name,
    sys_org.type,
    sys_org.order
from
	sys_org
where
	sys_org.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
