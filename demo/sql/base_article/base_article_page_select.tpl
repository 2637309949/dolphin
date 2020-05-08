select
    base_article.id
from
	base_article
where
	base_article.id {{.ne}} ""
	and
	base_article.del_flag {{.ne}} 1
LIMIT {{.size}} OFFSET {{.offset}}
