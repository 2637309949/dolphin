select
    count(*) records
from
	article
where
	article.id {{.ne}} ""
	and
	article.del_flag {{.ne}} 1
