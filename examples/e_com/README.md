# 常见电商项目的数据库表设计

## 简介：

#### 目的：

* 电商常用功能模块的数据库设计
* 常见问题的数据库解决方案

#### 环境：

* MySQL5.7
* 图形客户端，SQLyog
* Linux

#### 模块：

* 用户：注册、登陆
* 商品：浏览、管理
* 订单：生成、管理
* 仓配：库存、管理

## 电商实例数据库结构设计

### 电商项目用户模块

* 用户表涉及的实体
    
<figure>
   ![](https://ask.qcloudimg.com/http-save/yehe-1219068/8izjf0zj0j.jpeg?imageView2/2/w/1620)
</figure>

* 改进1：第三范式：将依赖传递的列分离出来。比如：登录名<-用户级别<-级别积分上限，级别积分下限 

<figure>
    ![](https://ask.qcloudimg.com/http-save/yehe-1219068/gcdaumy42s.jpeg?imageView2/2/w/1620)
</figure>

* 改进2：尽量做到冷热数据的分离，减小表的宽度

<figure>
    ![](https://ask.qcloudimg.com/http-save/yehe-1219068/m6liepskc1.jpeg?imageView2/2/w/1620)
</figure>

* 用户登录表(customer_login)

<pre class="prism-token token  language-javascript">
CREATE TABLE customer_login(
  customer_id INT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '用户ID',
  login_name VARCHAR(20) NOT NULL COMMENT '用户登录名',
  password CHAR(32) NOT NULL COMMENT 'md5加密的密码',
  user_stats TINYINT NOT NULL DEFAULT 1 COMMENT '用户状态',
  modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY pk_customerid(customer_id)
) ENGINE = innodb COMMENT '用户登录表'
</pre>

* 用户信息表(customer_inf)

<pre class="prism-token token  language-javascript">
CREATE TABLE customer_inf(
  customer_inf_id INT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '自增主键ID',
  customer_id INT UNSIGNED NOT NULL COMMENT 'customer_login表的自增ID',
  customer_name VARCHAR(20) NOT NULL COMMENT '用户真实姓名',
  identity_card_type TINYINT NOT NULL DEFAULT 1 COMMENT '证件类型：1 身份证，2 军官证，3 护照',
  identity_card_no VARCHAR(20) COMMENT '证件号码',
  mobile_phone INT UNSIGNED COMMENT '手机号',
  customer_email VARCHAR(50) COMMENT '邮箱',
  gender CHAR(1) COMMENT '性别',
  user_point INT NOT NULL DEFAULT 0 COMMENT '用户积分',
  register_time TIMESTAMP NOT NULL COMMENT '注册时间',
  birthday DATETIME COMMENT '会员生日',
  customer_level TINYINT NOT NULL DEFAULT 1 COMMENT '会员级别：1 普通会员，2 青铜，3白银，4黄金，5钻石',
  user_money DECIMAL(8,2) NOT NULL DEFAULT 0.00 COMMENT '用户余额',
  modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY pk_customerinfid(customer_inf_id)
) ENGINE = innodb COMMENT '用户信息表';
</pre>

* 用户级别表(customerlevelinf)

<pre class="prism-token token  language-javascript">
CREATE TABLE customer_level_inf(
  customer_level TINYINT NOT NULL AUTO_INCREMENT COMMENT '会员级别ID',
  level_name VARCHAR(10) NOT NULL COMMENT '会员级别名称',
  min_point INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '该级别最低积分',
  max_point INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '该级别最高积分',
  modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY pk_levelid(customer_level)
) ENGINE = innodb COMMENT '用户级别信息表';
</pre>

* 用户地址表(customer_addr)

<pre class="prism-token token  language-javascript">
CREATE TABLE customer_addr(
  customer_addr_id INT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '自增主键ID',
  customer_id INT UNSIGNED NOT NULL COMMENT 'customer_login表的自增ID',
  zip SMALLINT NOT NULL COMMENT '邮编',
  province SMALLINT NOT NULL COMMENT '地区表中省份的ID',
  city SMALLINT NOT NULL COMMENT '地区表中城市的ID',
  district SMALLINT NOT NULL COMMENT '地区表中的区ID',
  address VARCHAR(200) NOT NULL COMMENT '具体的地址门牌号',
  is_default TINYINT NOT NULL COMMENT '是否默认',
  modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY pk_customeraddid(customer_addr_id)
) ENGINE = innodb COMMENT '用户地址表';
</pre>

* 用户积分日志表(customerpointlog)

<pre class="prism-token token  language-javascript">
CREATE TABLE customer_point_log(
  point_id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '积分日志ID',
  customer_id INT UNSIGNED NOT NULL COMMENT '用户ID',
  source TINYINT UNSIGNED NOT NULL COMMENT '积分来源：0订单，1登陆，2活动',
  refer_number INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '积分来源相关编号',
  change_point SMALLINT NOT NULL DEFAULT 0 COMMENT '变更积分数',
  create_time TIMESTAMP NOT NULL COMMENT '积分日志生成时间',
  PRIMARY KEY pk_pointid(point_id)
) ENGINE = innodb COMMENT '用户积分日志表';
</pre>

* 用户余额变动表(customerbalancelog)

<pre class="prism-token token  language-javascript">
CREATE TABLE customer_balance_log(
  balance_id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '余额日志ID',
  customer_id INT UNSIGNED NOT NULL COMMENT '用户ID',
  source TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '记录来源：1订单，2退货单',
  source_sn INT UNSIGNED NOT NULL COMMENT '相关单据ID',
  create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录生成时间',
  amount DECIMAL(8,2) NOT NULL DEFAULT 0.00 COMMENT '变动金额',
  PRIMARY KEY pk_balanceid(balance_id)
) ENGINE = innodb COMMENT '用户余额变动表';
</pre>

* 用户登陆日志表(customerloginlog)

<pre class="prism-token token  language-javascript">
CREATE TABLE customer_login_log(
  login_id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '登陆日志ID',
  customer_id INT UNSIGNED NOT NULL COMMENT '登陆用户ID',
  login_time TIMESTAMP NOT NULL COMMENT '用户登陆时间',
  login_ip INT UNSIGNED NOT NULL COMMENT '登陆IP',
  login_type TINYINT NOT NULL COMMENT '登陆类型：0未成功，1成功',
  PRIMARY KEY pk_loginid(login_id)
) ENGINE = innodb COMMENT '用户登陆日志表';
</pre>

### Hash分区表

分区表特点：逻辑上为一个表，在物理上存储在多个文件中

<pre class="prism-token token  language-javascript">
CREATE TABLE customer_login_log(
  login_id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '登陆日志ID',
  customer_id INT UNSIGNED NOT NULL COMMENT '登陆用户ID',
  login_time TIMESTAMP NOT NULL COMMENT '用户登陆时间',
  login_ip INT UNSIGNED NOT NULL COMMENT '登陆IP',
  login_type TINYINT NOT NULL COMMENT '登陆类型：0未成功，1成功',
  PRIMARY KEY pk_loginid(login_id)
) ENGINE = innodb COMMENT '用户登陆日志表'
PARTITION BY HASH(customer_id) PARTITIONS 4;
</pre>

区别就在于加了 `PARTITION`这个命令。 文件结构上的区别

* 普通表结构:
    * `customer_login_log.frm`
    * `customer_login_log.ibd`
* 分区表结构:

    * `customer_login_log.frm`
    * `customer_login_log#P#p0.ibd`
    * `customer_login_log#P#p1.ibd`
    * `customer_login_log#P#p2.ibd`
    * `customer_login_log#P#p3.ibd`

按HASH分区的特点

* 根据MOD（分区建，分区数）的值把数据行存储到表的不同分区
* 数据可以平均的分布在各个分区中
* HASH分区的键值必须是一个INT类型的值，或是通过函数可以转为INT类型比如 `UNIX_TIMESTAMP(login_time)`

### Range分区表

特点：

* 根据分区键值的范围把数据行存储到表的不同分区中
* 多个分区的范围要连续，但是不能重复
* 默认情况下使用VALUES LESS THAN属性，即每个分区不包括指定的那个值

<pre class="prism-token token  language-javascript">
CREATE TABLE customer_login_log(
  login_id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '登陆日志ID',
  customer_id INT UNSIGNED NOT NULL COMMENT '登陆用户ID',
  login_time TIMESTAMP NOT NULL COMMENT '用户登陆时间',
  login_ip INT UNSIGNED NOT NULL COMMENT '登陆IP',
  login_type TINYINT NOT NULL COMMENT '登陆类型：0未成功，1成功',
  PRIMARY KEY pk_loginid(login_id)
) ENGINE = innodb COMMENT '用户登陆日志表'
PARTITION BY RANGE (customer_id) (
    PARTITION p0 VALUES LESS THAN (10000),
    PARTITION p1 VALUES LESS THAN (10000),
    PARTITION p2 VALUES LESS THAN (10000),
    PARTITION p3 VALUES LESS THAN MAXVALUE
);
</pre>

Range分区的适用范围

* 分区键为日期或是时间类型
* 所有SELECT查询中都包括分区键

### LIST分区

特点：

* 按分区键取值的列表进行分区
* 同范围分区一样，各分区的列表值不能重复
* 每一行数据必须能找到对应的分区列表，否则数据插入失败

<pre class="prism-token token  language-javascript">
CREATE TABLE customer_login_log(
  login_id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '登陆日志ID',
  customer_id INT UNSIGNED NOT NULL COMMENT '登陆用户ID',
  login_time TIMESTAMP NOT NULL COMMENT '用户登陆时间',
  login_ip INT UNSIGNED NOT NULL COMMENT '登陆IP',
  login_type TINYINT NOT NULL COMMENT '登陆类型：0未成功，1成功',
  PRIMARY KEY pk_loginid(login_id)
) ENGINE = innodb COMMENT '用户登陆日志表'
PARTITION BY LIST (login_type) (
    PARTITION p0 VALUES (1,3,5,7,9),
    PARTITION p1 VALUES (2,4,6,8)
);
</pre>

### 如何选择正确的分区类型

#### 如何为customerloginlog表分区

业务场景：

* 用户每次登录都会记录
* 日志保存一年，一年后可删除

解决：

* 使用RANGE范围分区
* 以login_type作为分区键

如何查看分区是否正确：

* 使用SELECT查询 `information_schema.PARTITIONS`
* 这里不使用MAXVALUE，防止后续的日期全部归到一个分区中，而是使用定时计划修改增加分区 `ALTER TABLE customer_login_log ADD PARTITION(PARTITION p4 VALUES LESS THAN(2018))`
* 删除以前一年的分区 `ALTER TABLE customer_login_log DROP PARTITION p0;`
* 过期数据归档
    * 建立用户登陆日志归档 `CREATE TABLE arch_customer_login_log(login_id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT'登陆日志ID',customer_id INT UNSIGNED NOT NULL COMMENT'登陆用户ID',login_time TIMESTAMP NOT NULL COMMENT'用户登陆时间',login_ip INT UNSIGNED NOT NULL COMMENT'登陆IP',login_type TINYINT NOT NULL COMMENT'登陆类型：0未成功，1成功',PRIMARY KEY pk_loginid(login_id))ENGINE=innodb COMMENT'用户登陆日志归档表'`
    * 归档操作： `ALTER TABLE customer_login_log EXCHANGE PARTITION p1 WITH TABLE arch_customer_login_log`
    * 迁移后删除： `ALTER TABLE customer_login_log DROP PARTITION p2`
    * 根据需要可以把归档的表引擎改为 `ARCHIVE`
    * 分区数据归档迁移条件
    * 操作步骤
    1.  [mysql](https://cloud.tencent.com/product/cdb?from=10680) >= 5.7
    2.  结构相同
    3.  归档到的数据表一定是非分区表
    4.  非临时表；不能有外键约束
    5.  归档引擎要是：archive

#### 使用分区表的注意事项

* 结合业务场景选择分区键，避免跨分区查询
* 对分区表进行查询最好在WHERE从句中包含分区键
* 具有主键或唯一索引的表，主键或唯一索引必须是分区键的一部分

### 商品实体

<figure>
    ![](https://ask.qcloudimg.com/http-save/yehe-1219068/59dvxul1tv.jpeg?imageView2/2/w/1620)
</figure>

* 品牌信息表（brand_info）

<pre class="prism-token token  language-javascript">
CREATE TABLE brand_info(
  brand_id SMALLINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '品牌ID',
  brand_name VARCHAR(50) NOT NULL COMMENT '品牌名称',
  telephone VARCHAR(50) NOT NULL COMMENT '联系电话',
  brand_web VARCHAR(100) COMMENT '品牌网络',
  brand_logo VARCHAR(100) COMMENT '品牌logo URL',
  brand_desc VARCHAR(150) COMMENT '品牌描述',
  brand_status TINYINT NOT NULL DEFAULT 0 COMMENT '品牌状态,0禁用,1启用',
  brand_order TINYINT NOT NULL DEFAULT 0 COMMENT '排序',
  modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY pk_brandid (brand_id)
)ENGINE=innodb COMMENT '品牌信息表';
</pre>

* 分类信息表(product_category)

<pre class="prism-token token  language-javascript">
CREATE TABLE product_category(
  category_id SMALLINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '分类ID',
  category_name VARCHAR(10) NOT NULL COMMENT '分类名称',
  category_code VARCHAR(10) NOT NULL COMMENT '分类编码',
  parent_id SMALLINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '父分类ID',
  category_level TINYINT NOT NULL DEFAULT 1 COMMENT '分类层级',
  category_status TINYINT NOT NULL DEFAULT 1 COMMENT '分类状态',
  modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT  '最后修改时间',
  PRIMARY KEY pk_categoryid(category_id)
)ENGINE=innodb COMMENT '商品分类表'
</pre>

* 供应商信息表（supplier_info）

<pre class="prism-token token  language-javascript">
CREATE TABLE supplier_info(
  supplier_id INT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '供应商ID',
  supplier_code CHAR(8) NOT NULL COMMENT '供应商编码',
  supplier_name CHAR(50) NOT NULL COMMENT '供应商名称',
  supplier_type TINYINT NOT NULL COMMENT '供应商类型：1.自营，2.平台',
  link_man VARCHAR(10) NOT NULL COMMENT '供应商联系人',
  phone_number VARCHAR(50) NOT NULL COMMENT '联系电话',
  bank_name VARCHAR(50) NOT NULL COMMENT '供应商开户银行名称',
  bank_account VARCHAR(50) NOT NULL COMMENT '银行账号',
  address VARCHAR(200) NOT NULL COMMENT '供应商地址',
  supplier_status TINYINT NOT NULL DEFAULT 0 COMMENT '状态：0禁止，1启用',
  modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT  '最后修改时间',
  PRIMARY KEY pk_supplierid(supplier_id)
) ENGINE = innodb COMMENT '供应商信息表';
</pre>

* 商品信息表(product_info)
    * 宽度较宽，字段差不多一起使用
    * 可以被缓存

<pre class="prism-token token  language-javascript">
CREATE TABLE product_info(
  product_id INT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '商品ID',
  product_core CHAR(16) NOT NULL COMMENT '商品编码',
  product_name VARCHAR(20) NOT NULL COMMENT '商品名称',
  bar_code VARCHAR(50) NOT NULL COMMENT '国条码',
  brand_id INT UNSIGNED NOT NULL COMMENT '品牌表的ID',
  one_category_id SMALLINT UNSIGNED NOT NULL COMMENT '一级分类ID',
  two_category_id SMALLINT UNSIGNED NOT NULL COMMENT '二级分类ID',
  three_category_id SMALLINT UNSIGNED NOT NULL COMMENT '三级分类ID',
  supplier_id INT UNSIGNED NOT NULL COMMENT '商品的供应商ID',
  price DECIMAL(8,2) NOT NULL COMMENT '商品销售价格',
  average_cost DECIMAL(18,2) NOT NULL COMMENT '商品加权平均成本',
  publish_status TINYINT NOT NULL DEFAULT 0 COMMENT '上下架状态：0下架1上架',
  audit_status TINYINT NOT NULL DEFAULT 0 COMMENT '审核状态：0未审核，1已审核',
  weight FLOAT COMMENT '商品重量',
  length FLOAT COMMENT '商品长度',
  height FLOAT COMMENT '商品高度',
  width FLOAT COMMENT '商品宽度',
  color_type ENUM('红','黄','蓝','黑'),
  production_date DATETIME NOT NULL COMMENT '生产日期',
  shelf_life INT NOT NULL COMMENT '商品有效期',
  descript TEXT NOT NULL COMMENT '商品描述',
  indate TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '商品录入时间',
  modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY pk_productid(product_id)
) ENGINE = innodb COMMENT '商品信息表';
</pre>

* 商品图片表(productpicinfo)

<pre class="prism-token token  language-javascript">
CREATE TABLE product_pic_info(
  product_pic_id INT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '商品图片ID',
  product_id INT UNSIGNED NOT NULL COMMENT '商品ID',
  pic_desc VARCHAR(50) COMMENT '图片描述',
  pic_url VARCHAR(200) NOT NULL COMMENT '图片URL',
  is_master TINYINT NOT NULL DEFAULT 0 COMMENT '是否主图：0.非主图1.主图',
  pic_order TINYINT NOT NULL DEFAULT 0 COMMENT '图片排序',
  pic_status TINYINT NOT NULL DEFAULT 1 COMMENT '图片是否有效：0无效 1有效',
  modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT  '最后修改时间',
  PRIMARY KEY pk_picid(product_pic_id)
)ENGINE=innodb COMMENT '商品图片信息表';
</pre>

* 商品评论表（product_comment）

<pre class="prism-token token  language-javascript">
CREATE TABLE product_comment(
  comment_id INT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '评论ID',
  product_id INT UNSIGNED NOT NULL COMMENT '商品ID',
  order_id BIGINT UNSIGNED NOT NULL COMMENT '订单ID',
  customer_id INT UNSIGNED NOT NULL COMMENT '用户ID',
  title VARCHAR(50) NOT NULL COMMENT '评论标题',
  content VARCHAR(300) NOT NULL COMMENT '评论内容',
  audit_status TINYINT NOT NULL COMMENT '审核状态：0未审核，1已审核',
  audit_time TIMESTAMP NOT NULL COMMENT '评论时间',
  modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY pk_commentid(comment_id)
) ENGINE = innodb COMMENT '商品评论表';
</pre>

### 订单模块

<figure>
    ![](https://ask.qcloudimg.com/http-save/yehe-1219068/blh2j3vjtp.jpeg?imageView2/2/w/1620)
</figure>

* 订单主表（order_master）

<pre class="prism-token token  language-javascript">
CREATE TABLE order_master(
  order_id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  order_sn BIGINT UNSIGNED NOT NULL COMMENT '订单编号 yyyymmddnnnnnnnn',
  customer_id INT UNSIGNED NOT NULL COMMENT '下单人ID',
  shipping_user VARCHAR(10) NOT NULL COMMENT '收货人姓名',
  province SMALLINT NOT NULL COMMENT '省',
  city SMALLINT NOT NULL COMMENT '市',
  district SMALLINT NOT NULL COMMENT '区',
  address VARCHAR(100) NOT NULL COMMENT '地址',
  payment_method TINYINT NOT NULL COMMENT '支付方式：1现金，2余额，3网银，4支付宝，5微信',
  order_money DECIMAL(8,2) NOT NULL COMMENT '订单金额',
  district_money DECIMAL(8,2) NOT NULL DEFAULT 0.00 COMMENT '优惠金额',
  shipping_money DECIMAL(8,2) NOT NULL DEFAULT 0.00 COMMENT '运费金额',
  payment_money DECIMAL(8,2) NOT NULL DEFAULT 0.00 COMMENT '支付金额',
  shipping_comp_name VARCHAR(10) COMMENT '快递公司名称',
  shipping_sn VARCHAR(50) COMMENT '快递单号',
  create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '下单时间',
  shipping_time DATETIME COMMENT '发货时间',
  pay_time DATETIME COMMENT '支付时间',
  receive_time DATETIME COMMENT '收货时间',
  order_status TINYINT NOT NULL DEFAULT 0 COMMENT '订单状态',
  order_point INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单积分',
  invoice_time VARCHAR(100) COMMENT '发票抬头',
  modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY pk_orderid(order_id)
)ENGINE = innodb COMMENT '订单主表';
</pre>

* 订单详情表（order_detail）

<pre class="prism-token token  language-javascript">
CREATE TABLE order_detail(
  order_detail_id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单详情表ID',
  order_id INT UNSIGNED NOT NULL COMMENT '订单表ID',
  product_id INT UNSIGNED NOT NULL COMMENT '订单商品ID',
  product_name VARCHAR(50) NOT NULL COMMENT '商品名称',
  product_cnt INT NOT NULL DEFAULT 1 COMMENT '购买商品数量',
  product_price DECIMAL(8,2) NOT NULL COMMENT '购买商品单价',
  average_cost DECIMAL(8,2) NOT NULL COMMENT '平均成本价格',
  weight FLOAT COMMENT '商品重量',
  fee_money DECIMAL(8,2) NOT NULL DEFAULT 0.00 COMMENT '优惠分摊金额',
  w_id INT UNSIGNED NOT NULL COMMENT '仓库ID',
    modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY pk_orderdetailid(order_detail_id)
)ENGINE = innodb COMMENT '订单详情表'
</pre>

* 购物车表（order_cart）

<pre class="prism-token token  language-javascript">
CREATE TABLE order_cart(
  cart_id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '购物车ID',
  customer_id INT UNSIGNED NOT NULL COMMENT '用户ID',
  product_id INT UNSIGNED NOT NULL COMMENT '商品ID',
  product_amount INT NOT NULL COMMENT '加入购物车商品数量',
  price DECIMAL(8,2) NOT NULL COMMENT '商品价格',
  add_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入购物车时间',
      modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY pk_cartid(cart_id)
) ENGINE = innodb COMMENT '购物车表';
</pre>

* 仓库信息表（warehouse_info）

<pre class="prism-token token  language-javascript">
CREATE TABLE warehouse_info(
  w_id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '仓库ID',
  warehouse_sn CHAR(5) NOT NULL COMMENT '仓库编码',
  warehoust_name VARCHAR(10) NOT NULL COMMENT '仓库名称',
  warehouse_phone VARCHAR(20) NOT NULL COMMENT '仓库电话',
  contact VARCHAR(10) NOT NULL COMMENT '仓库联系人',
  province SMALLINT NOT NULL COMMENT '省',
  city SMALLINT NOT NULL COMMENT '市',
  distrct SMALLINT NOT NULL COMMENT '区',
  address VARCHAR(100) NOT NULL COMMENT '仓库地址',
  warehouse_status TINYINT NOT NULL DEFAULT 1 COMMENT '仓库状态：0禁用，1启用',
        modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY pk_wid(w_id)
)ENGINE = innodb COMMENT '仓库信息表';
</pre>

* 商品库存表（warehouse_product）

<pre class="prism-token token  language-javascript">
CREATE TABLE warehouse_product(
  wp_id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品库存ID',
  product_id INT UNSIGNED NOT NULL COMMENT '商品ID',
  w_id SMALLINT UNSIGNED NOT NULL COMMENT '仓库ID',
  current_cnt INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '当前商品数量',
  lock_cnt INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '当前占用数据',
  in_transit_cnt INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '在途数据',
  average_cost DECIMAL(8,2) NOT NULL DEFAULT 0.00 COMMENT '移动加权成本',
  modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY pk_wpid(wp_id)
)ENGINE = innodb COMMENT '商品库存表'
</pre>

* 物流公司信息表（shipping_info）

<pre class="prism-token token  language-javascript">
CREATE TABLE shipping_info(
  ship_id TINYINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  ship_name VARCHAR(20) NOT NULL COMMENT '物流公司名称',
  ship_contact VARCHAR(20) NOT NULL COMMENT '物流公司联系人',
  telephone VARCHAR(20) NOT NULL COMMENT '物流公司联系电话',
  price DECIMAL(8,2) NOT NULL DEFAULT 0.00 COMMENT '配送价格',
    modified_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY pk_shipid(ship_id)
)ENGINE = innodb COMMENT '物流公司信息表';
</pre>

### DB规划

* 为以后数据库迁移提供方便
* 避免跨库操作，把经常一起关联查询的表放到一个DB中
* 为方便识别表所在的DB，在表名前增加库名前缀
* 用户数据库(mccustomerdb)
    * customerinf
    * customer_login
    * customerlevelinf
    * customerloginlog
    * customerpointlog
    * customerbalancelog
* 商品数据库(mcproductdb)
    * productinfo
    * productpicinfo
    * productcategory
    * productsupplierinfo
    * productcomment
    * productbrandinfo
* 订单数据库(mcorderdb)
    * ordermaster
    * orderdetail
    * ordercustomeraddr
    * ordercart
    * shippinginfo
    * warehouseinfo
    * warehouse_product

