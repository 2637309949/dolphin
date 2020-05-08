select
    count(*) records
from
	base_article
where
	base_article.id {{.ne}} ""
	and
	base_article.del_flag {{.ne}} 1
