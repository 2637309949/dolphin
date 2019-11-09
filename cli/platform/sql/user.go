package sql

func init() {
	SQLTPL["user_page_select.tpl"] = `select user.* from user`
	SQLTPL["user_page_count.tpl"] = `select count(*) records from user `
}
