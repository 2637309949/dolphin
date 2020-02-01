## Dolphin
![dolphin flash](./assets/dolphin.webp)

### Bean
```xml
<bean
    name="activity_info"
    desc="活动信息"
    packages="github.com/2637309949/dolphin/cli/packages/null"
    extends="$applet_activity">
    <prop name="code" desc="编码" type="null.String" />
    <prop name="name" desc="名称" type="null.String" />
</bean>
```

### Controler
```xml
<!-- rule of controller name: ?<end>_?<module>_<ctr>_?<action> -->
<controller
    name="applet_activity"
    desc="活动">
    <api name="batch_add" function="add" table="applet_activity" desc="添加活动" method="post">
        <param name="activity" type="[]$applet_activity" desc="活动对象" />
        <return>
            <success type="$response" />
            <failure type="$response" />
        </return>
    </api>
    <api name="add" function="add" table="applet_activity" desc="添加活动" method="post">
        <param name="activity" type="$applet_activity" desc="活动对象" />
        <return>
            <success type="$response" />
            <failure type="$response" />
        </return>
    </api>
    <api name="batch_del" function="delete" table="applet_activity" desc="删除活动" method="delete">
        <param name="activity" type="[]$applet_activity" desc="活动对象" />
        <return>
            <success type="$response" />
            <failure type="$response" />
        </return>
    </api>
    <api name="del" function="delete" table="applet_activity" desc="删除活动" method="delete">
        <param name="activity" type="$applet_activity" desc="活动对象" />
        <return>
            <success type="$response" />
            <failure type="$response" />
        </return>
    </api>
    <api name="batch_update" function="update" table="applet_activity" desc="更新活动" method="put">
        <param name="activity" type="[]$applet_activity" desc="活动对象" />
        <return>
            <success type="$response" />
            <failure type="$response" />
        </return>
    </api>
    <api name="update" function="update" table="applet_activity" desc="更新活动" method="put">
        <param name="activity" type="$applet_activity" desc="活动对象" />
        <return>
            <success type="$response" />
            <failure type="$response" />
        </return>
    </api>
    <api name="list" function="list" table="applet_activity" desc="活动分页查询" method="get">
        <param name="page" type="int" desc="页码" value="1" />
        <param name="size" type="int" desc="单页数" value="20" />
        <param name="title" type="string" desc="标题筛选" value="nn" />
        <param name="hidden" type="int" desc="是否隐藏筛选" />
        <return>
            <success type="$response" />
            <failure type="$response" />
        </return>
    </api>
    <api name="one" function="one" table="applet_activity" desc="获取活动" method="get">
        <param name="id" type="string" desc="活动id" />
        <return>
            <success type="$response" />
            <failure type="$response" />
        </return>
    </api>
    <api name="increase" table="applet_activity" desc="增加次数" method="post" version="1.0" auth="false">
        <param name="applet_activity" type="$applet_activity" desc="记录id" />
        <return>
            <success type="$response" />
            <failure type="$response" />
        </return>
    </api>
    <api name="increase_v2" table="applet_activity" desc="增加次数" method="post" version="2.0" auth="false">
        <param name="id" type="string" desc="记录id" />
        <return>
            <success type="$response" />
            <failure type="$response" />
        </return>
    </api>
</controller>
```

### Table
```xml
<table
    name="applet_activity"
    desc="活动信息"
    packages="github.com/2637309949/dolphin/cli/packages/null">
	<column name="id" desc="主键" type="null.String" xorm="varchar(36) notnull unique pk" />
	<column name="title" desc="标题" type="null.String" xorm="varchar(36)" />
	<column name="content" desc="内容" type="null.String" xorm="varchar(1500)" />
	<column name="campus" desc="校区" type="null.String" xorm="varchar(2000)" />
	<column name="image" desc="图片" type="null.String" xorm="varchar(500)" />
	<column name="type" desc="类型" type="null.Int" />
	<column name="forward_count" desc="转发次数" type="null.Int" />
	<column name="like_count" desc="点赞次数" type="null.Int" />
	<column name="share_count" desc="分享次数" type="null.Int" />
	<column name="read_count" desc="阅读次数" type="null.Int" />
	<column name="priority" desc="优先级" type="null.Int" />
	<column name="is_entry" desc="是否报名 0(否) 1(是)" type="null.Int" />
	<column name="hidden" desc="是否隐藏 0(否) 1(是)" type="null.Int" />
	<column name="create_by" desc="创建人" type="null.String" xorm="varchar(36)" />
	<column name="create_time" desc="创建时间" type="null.Time" xorm="datetime" />
	<column name="update_by" desc="最后更新人" type="null.String" xorm="varchar(36)" />
	<column name="update_time" desc="最后更新时间" type="null.Time" xorm="datetime" />
	<column name="delete_by" desc="删除人" type="null.String" xorm="varchar(36)" />
	<column name="delete_time" desc="删除时间" type="null.Time" xorm="datetime" />
</table>
```