package sql

func init() {
	SQLTPL["platform_user_page_select.tpl"] = `
	select
    user.*
	from
    user
	`
}
