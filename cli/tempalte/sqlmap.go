package tempalte

// TmplSQLMap defined template
var TmplSQLMap = `
<sqlMap>
    <sql id="insert_{{.Table.TableName .Name .Table.Name}}">
        insert into {{.Table.TableName .Name .Table.Name}}
        (
        {{- range .Table.Columns}}
        {{- end}}
        )
		values
		(?id,?activity_title,?activity_cover,?activity_intro,?activity_detail,?apply_begin_time,?apply_end_time,?activity_begin_time,?activity_end_time,?activity_join_condition,?activity_selection_rules,?activity_state,?activity_create_time,?single_dubbing,?create_by,?create_time,?last_update_by,?last_update_time,?del_flag)
    </sql>
    <sql id="update_{{.Table.TableName .Name .Table.Name}}">
        update {{.Table.TableName .Name .Table.Name}} set
		where id = ?id
    </sql>
    <sql id="deleteone_{{.Table.TableName .Name .Table.Name}}">
        delete from {{.Table.TableName .Name .Table.Name}}
		where id =?id
    </sql>
    <sql id="selectone_{{.Table.TableName .Name .Table.Name}}">
        select 
        from {{.Table.TableName .Name .Table.Name}}
		where id = ?id
    </sql>
    <sql id="selectall_{{.Table.TableName .Name .Table.Name}}">
        select 
        from {{.Table.TableName .Name .Table.Name}}
    </sql>
</sqlMap>
`
