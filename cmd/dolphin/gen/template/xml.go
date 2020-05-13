package template

// TmplXML defined template
var TmplXML = `<?xml version="1.0" encoding="utf-8" ?>
<application name="%v" desc="dolphin template" packagename="%v">
    <controllers>
        <controller name="base_article" desc="文章接口定义">
            <api name="add" function="add" table="base_article" desc="添加文章" method="post">
                <param name="user" type="$base_article" desc="文章信息" />
                <return>
                    <success type="$success"/>
                    <failure type="$fail"/>
                </return>
            </api>
            <api name="del" function="delete" table="base_article" desc="删除文章" method="delete">
                <param name="base_article" type="$base_article" desc="文章" />
                <return>
                    <success type="$success"/>
                    <failure type="$fail"/>
                </return>
            </api>
            <api name="update" function="update" table="base_article" desc="更新文章" method="put">
                <param name="user" type="$base_article" desc="文章信息" />
                <return>
                    <success type="$success"/>
                    <failure type="$fail"/>
                </return>
            </api>
            <api name="page" function="page" table="base_article" desc="文章分页查询" method="get">
                <param name="page" type="int" desc="页码" value="1"/>
                <param name="size" type="int" desc="单页数" value="15" />
                <param name="app_name" type="string" desc="所属应用"/>
                <return>
                    <success type="$success"/>
                    <failure type="$fail"/>
                </return>
            </api>
            <api name="get" desc="获取文章信息" function="one" table="base_article" method="get">
                <param name="id" type="string" desc="文章id" />
                <return>
                    <success type="$success"/>
                    <failure type="$fail"/>
                </return>
            </api>
        </controller>
    </controllers>
    <tables>
        <table name="base_article" desc="文章" packages="github.com/2637309949/dolphin/packages/null">
            <column name="id" desc="主键" type="null.String" xorm="varchar(36) notnull unique pk" />

            <column name="create_by" desc="创建人" type="null.String" xorm="varchar(36)" />
            <column name="create_time" desc="创建时间" type="null.Time" xorm="datetime" />
            <column name="update_by" desc="最后更新人" type="null.String" xorm="varchar(36)" />
            <column name="update_time" desc="最后更新时间" type="null.Time" xorm="datetime" />
            <column name="del_flag" desc="删除标记" type="null.Int" xorm="notnull" />
            <column name="remark" desc="备注" type="null.String" xorm="varchar(200)" />
        </table>
    </tables>
</application>
`
