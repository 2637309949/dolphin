### application
```xml
<?xml version="1.0" encoding="utf-8" ?>
<application name="scene" desc="dolphin template" packagename="scene"/>
```

### bean
```xml
<bean name="article_info" desc="文章信息" 
    packages="github.com/2637309949/dolphin/packages/null" extends="$article">
    <prop name="url" desc="地址" type="null.String" />
</bean>
```

### controller
```xml
<controller name="article" desc="文章">
    <api name="add" func="add" table="article" method="post" desc="添加文章">
        <param name="article" type="$article" desc="文章信息"/>
        <return>
            <success type="$success"/>
            <failure type="$fail"/>
        </return>
    </api>
    <api name="batch_add" func="add" table="article" desc="添加文章" method="post">
        <param name="article" type="[]$article" desc="文章信息" />
        <return>
            <success type="$success"/>
            <failure type="$fail"/>
        </return>
    </api>
    <api name="del" func="delete" table="article" method="delete" desc="删除文章">
        <param name="article" type="$article" desc="文章"/>
        <return>
            <success type="$success"/>
            <failure type="$fail"/>
        </return>
    </api>
    <api name="batch_del" func="delete" table="article" desc="删除文章" method="put">
        <param name="article" type="[]$article" desc="文章信息" />
        <return>
            <success type="$success"/>
            <failure type="$fail"/>
        </return>
    </api>
    <api name="update" func="update" table="article" desc="更新文章" method="put">
        <param name="article" type="$article" desc="文章信息" />
        <return>
            <success type="$success"/>
            <failure type="$fail"/>
        </return>
    </api>
    <api name="batch_update" func="update" table="article" desc="更新文章" method="put">
        <param name="article" type="[]$article" desc="文章信息" />
        <return>
            <success type="$success"/>
            <failure type="$fail"/>
        </return>
    </api>
    <api name="page" func="page" table="article" method="get" desc="文章分页查询">
        <param name="page" type="int" value="1" desc="页码"/>
        <param name="size" type="int" value="15" desc="单页数"/>
        <return>
            <success type="$success"/>
            <failure type="$fail"/>
        </return>
    </api>
    <api name="get" func="one" table="article" method="get" desc="获取文章信息">
        <param name="id" type="string" desc="文章id" />
        <return>
            <success type="$success"/>
            <failure type="$fail"/>
        </return>
    </api>
    <api name="payment" method="post" desc="文章付费">
        <param name="article" type="$article_info" desc="文章"/>
        <return>
            <success type="$success"/>
            <failure type="$fail"/>
        </return>
    </api>
</controller>
```

### rpc
```xml
<service name="message_srv" desc="消息服务">
    <rpc name="send_mail" desc="发送邮件">
        <request type="$message_mail" desc="邮件"/>
        <reply type="$message_reply" desc="反馈"/>
    </rpc>
</service>
```

### table
```xml
<table name="article" desc="文章" 
    packages="github.com/2637309949/dolphin/packages/null,github.com/shopspring/decimal">
    <column name="id" desc="主键" type="null.Int" xorm="bigint(20) notnull autoincr unique pk" />
    <column name="unanswered_count" desc="未答复数目" type="null.Int" />
    <column name="best_answerers_count" desc="最佳答复数目" type="null.Int" />
    <column name="is_super_topic_vote" desc="是否超级话题投票" type="null.Int" />
    <column name="excerpt" desc="摘录" type="null.String" xorm="varchar(512)" />
    <column name="is_vote" desc="是否投票" type="null.Int" />
    <column name="is_black" type="null.Int" />
    <column name="questions_count" desc="提问数目" type="null.Int" />
	<column name="category" desc="分类" type="null.String" xorm="varchar(36)" />
	<column name="name" desc="标题" type="null.String" xorm="varchar(108)" />
    <column name="introduction" desc="简介" type="null.String" xorm="varchar(512)" />
    <column name="url" desc="地址" type="null.String" xorm="varchar(512)" />
    <column name="followers_count" desc="粉丝数" type="null.Int" />
	<column name="type" desc="类别" type="null.String" xorm="varchar(36)" />
    <column name="reward" desc="打赏" type="decimal.Decimal" xorm="decimal(6,2)" />

	<column name="creater" desc="创建人" type="null.Int" xorm="bigint(20) notnull" />
	<column name="create_time" desc="创建时间" type="null.Time" xorm="datetime" />
	<column name="updater" desc="最后更新人" type="null.Int" xorm="bigint(20) notnull" />
	<column name="update_time" desc="最后更新时间" type="null.Time" xorm="datetime" />
	<column name="is_delete" desc="删除标记" type="null.Int" xorm="notnull" />
	<column name="remark" desc="备注" type="null.String" xorm="varchar(200)" />
</table>
```