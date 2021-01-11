mysql.md
```
MYSQL

	数据库操作
		创建数据库
			CREATE DATABASE db_name DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci
		查询数据库
		查询所有数据库
			SHOW DATABASES
			查询数据库建表时的sql
			SHOW CREATE DATABASE db_name;
		删除数据库
			DROP DATABASE db_name;
		修改数据库
			修改数据库的字符编码和排序方式
			ALTER DATABASE db_name DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
		选择数据库
			USE db_name;
		命令行设置之后操作的编码格式
			SET NAMES UTF8
		
	数据库表操作
		创建表
			SQL : CREATE TABLE tb_name (建表的字段、类型、长度、约束、默认、注释)
			约束
			非空
				NOT NULL
			非负
				UNSIGNED
			主键
				PRIMARY KEY
			自增
				AUTO_INCREMENT
			默认
				DEFAULT
			注释
				COMMENT
			常用类型
				极小整形
					TINYINT
						非负最大值255
						1个字节
				小整形
					SAMLLINT
						非负最大值65535
						2个字节
				整形
					INT
						非负最大值4 294 967 295
						4个字节
				单精度
					FLOAT
						4个字节
						定长字符串
					CHAR
						最大保存255个字节
						如果值没有到给定长度用空格补充
						变长字符串
					VARCHAR
						最大保存255个字节
						用多大占多大
				文本
					TEXT
						最大保存65535个字节
			表字段索引
				唯一索引
					添加
						创建索引
							CREATE UNIQUE INDEX index_name ON tb_name (account);
						表字段修改
							ALTER TABLE tb_name ADD UNIQUE index_name(field_name);
					删除
						DROP INDEX 索引名称 ON 表名
				普通索引
					添加
						表字段修改
							ALTER TABLE 表名 ADD INDEX 索引名称(字段名称);
						创建索引
							CREATE INDEX Index_name ON tb_name(`account`);
							CREATE INDEX 索引名称 ON 表名(字段名);
						删除
							DROP INDEX 索引名称 ON 表名
				主键
					添加
						ALTER TABLE tb_name ADD PRIMARY KEY (field_name);
						ALTER TABLE 表名 ADD PRIMARY KEY (字段名称)
					删除
						ALTER TABLE tb_name DROP PRIMARY KEY;
				联合索引
					添加
						ALTER TABLE tb_name ADD index_name (field_name1,field_name2);
					删除
						DROP INDEX 索引名称 ON 表名
				索引最左前缀原理：
					联合索引最左
						联合索引 (A,B,C),如果查询没有A,则不能走索引,查询条件有AC,只能走A的索引
					字段索引最左
						字符串 "abc",只能like "a%"走索引,like "%bc"则不能走索引
		修改表
			表字段的增删改查
				字段添加
					ALTER TABLE tb_name ADD address VARCHAR (100) NOT NULL DEFAULT '' COMMENT '地址';
					ALERT TABLE tb_name ADD 添加字段 字段类型  非空约束  默认  注释
				字段类型修改
					ALTER TABLE tb_name MODIFY address VARCHAR (50) NOT NULL DEFAULT '' COMMENT '地址';
					ALERT TABLE tb_name MODIFY 字段名称 新的字段类型  非负  非空  默认 注释
				字段名称类型修改
					ALTER TABLE tb_name CHANGE address addr VARCHAR (100) NOT NULL DEFAULT '' COMMENT '地址';
					ALTER TABLE tb_name CHANGE 旧的字段名  新的字段名  新的类型  约束  默认  注释
				字段类型查询
					DESC tb_name;
				字段删除
				ALTER TABLE tb_name DROP addr;
				ALTER TABLE 表名 DROP 删除的字段名
			表修改
				表名修改
					ALTER TABLE tb_name RENAME TO new_tb_name;
					ALTER TABLE 旧表名 RENAME TO 新表名
				引擎修改
					ALTER TABLE tb_name ENGINE = InnoDB;
					ALTER TABLE 表名 ENGINE = 新的引擎名称
		删除表
			DROP TABLE tb_name;
		查询表
			查询所有表
				SHOW TABLES;
			查询建表时的sql
				SHOW CREATE TABLE tb_name;
				
	MySQL函数
		1.数学函数
			ABS()绝对值
			SQRT(x) 求二次方根
			MOD(x,y)                 返回x/y的模（余数）
			CEIL 和 CEILING 向上取整
			FLOOR 向下取整
			RAND()返回０到１内的随机值,可以通过提供一个参数(种子)使RAND()随机数生成器生成一个指定的值。
			ROUND(x,y)返回参数x的四舍五入的有y位小数的值
			SIGN(x) 返回参数的符号，x 的值为负、零和正时返回结果依次为 -1、0 和 1。
			POW(x,y) 和 POWER(x,y) 函数对参数 x 进行 y 次方的求值。
			SIN(x) 函数计算正弦值
			ASIN(X) 求反正弦值，与函数 SIN 互为反函数
			COS(X) 求余弦值
			ACOS(X) 求反余弦值，与函数 COS 互为反函数
			TAN(X) 求正切值
			ATAN(X) 求反正切值，与函数 TAN 互为反函数
			COT(X) 求余切值
			PI()  π值
		2.聚合函数
			MAX(X) 查询指定列的最大值
			MIN(X) 查询指定列的最小值
			AVG(col)返回指定列的平均值
			COUNT(col)返回指定列中非NULL值的个数
			SUM(col)返回指定列的所有值之和
		3.字符串函数
			TRIM(str)去除字符串首部和尾部的所有空格
			RTRIM(str) 返回字符串str尾部的空格
			LTRIM(str) 从字符串str中切掉开头的空格
			LENGTH(s)返回字符串str中的字符数
			ASCII(char)返回字符的ASCII码值
			CONCAT(sl，s2，...) 合并字符串函数，返回结果为连接参数产生的字符串，参数可以使一个或多个
			INSERT(s1，x，len，s2) 返回字符串 s1，子字符串起始于 x 位置，并且用 len 个字符长的字符串代替 s2。
			UCASE(str)或UPPER(str) 返回将字符串str中所有字符转变为大写后的结果
			LOWER(str) 将字符串 str 中的字母字符全部转换成小写。
			LEFT(s，n) 返回字符串 s 最左边的 n 个字符。
			RIGHT(s，n) 返回字符串 s 最右边的 n 个字符。
			REPLACE(s，s1，s2) 使用字符串 s2 替换字符串 s 中所有的字符串 s1。
			SUBSTRING(s，n，len) 带有 len 参数的格式，从字符串 s 返回一个长度同 len 字符相同的子字符串，起始于位置 n。
			REVERSE(s) 字符串 s 反转，返回的字符串的顺序和 s 字符串的顺序相反。
		4.日期和时间函数
			CURDATE()或CURRENT_DATE() 返回当前的日期
			CURTIME()或CURRENT_TIME() 返回当前的时间
			NOW 和 SYSDATE 两个函数作用相同，返回当前系统的日期和时间值
			UNIX_TIMESTAMP()获取系统当前的时间戳
			FROM_UNIXTIME(时间戳) 格式化传入的时间戳，转成日期格式
			MONTH 获取指定日期中的月份
			MONTHNAME 获取指定日期中的月份英文名称
			DAYNAME 获取指定曰期对应的星期几的英文名称
			DAYOFWEEK 获取指定日期对应的一周的索引位置值
			WEEK 获取指定日期是一年中的第几周，返回值的范围是否为 0〜52 或 1〜53
			DAYOFYEAR 获取指定曰期是一年中的第几天，返回值范围是1~366
			DAYOFMONTH 获取指定日期是一个月中是第几天，返回值范围是1~31
			YEAR 获取年份，返回值范围是 1970〜2069
			TIME_TO_SEC 将时间参数转换为秒数
			SEC_TO_TIME 将秒数转换为时间，与TIME_TO_SEC 互为反函数
			DATE_ADD 和 ADDDATE 两个函数功能相同，都是向日期添加指定的时间间隔
			DATE_SUB 和 SUBDATE 两个函数功能相同，都是向日期减去指定的时间间隔
			ADDTIME 时间加法运算，在原始时间上添加指定的时间
			SUBTIME 时间减法运算，在原始时间上减去指定的时间
			DATEDIFF 获取两个日期之间间隔，返回参数 1 减去参数 2 的值
			DATE_FORMAT 格式化指定的日期，根据参数返回指定格式的值
			WEEKDAY 获取指定日期在一周内的对应的工作日索引
		5.加密函数
			MD5()    计算字符串str的MD5校验和
			PASSWORD(str)   返回字符串str的加密版本，这个加密过程是不可逆转的，和UNIX密码加密过程使用不同的算法。
			SHA()    计算字符串str的安全散列算法(SHA)校验和
			AES_DECRYPT(str,key) 返回用密钥key对字符串str利用高级加密标准算法解密后的结果
			DECODE(str,key) 使用key作为密钥解密加密字符串str
			ENCRYPT(str,salt) 使用UNIXcrypt()函数，用关键词salt(一个可以惟一确定口令的字符串，就像钥匙一样)加密字符串str
			ENCODE(str,key) 使用key作为密钥加密字符串str，调用ENCODE()的结果是一个二进制字符串，它以BLOB类型存储
		6.控制流程函数
			CASE()  比对某个字段值,返回相应的值
			CASE value WHEN [compare-value1] THEN result1 [WHEN [compare-value2] THEN result2 [ELSE result3] END
			IF(expr1,expr2,expr3)  如果expr1位true,则返回expr2,否则返回expr3
			IFNULL(expr1,expr2)  如果expr1为NULL,则返回expr2,否则返回expr1
		7.格式化函数
			FORMAT()
		8.类型转化函数
			CAST(expr AS type) 将任意类型的表达式expr转换成指定类型type的值。
		9.系统信息函数
			DATABASE() 或SCHEMA()   返回当前数据库名
			USER()、 SYSTEM_USER()、 SESSION_USER()、 CURRENT_USER() 和CURRENT_USER()  返回当前登陆用户名
			VERSION()   返回MySQL服务器的版本
			BENCHMARK(count,expr)  将表达式expr重复运行count次
	优化
		步骤：
			1.开启慢查询日志,设置阈值,比如超过5秒钟的就是慢SQL,并将它抓取出来;
			2.EXPLAIN+慢SQL分析;
			3.SHOW profile,查询SQL在MySQL服务器里面的执行细节和生命周期情况
			4.具体优化
		explain
			语法explain select * from xxl_job_log l where l.job_id in (select id from xxl_job_info)
			字段说明
				CONNECTION_ID()   返回当前客户的连接ID
				id：表示查询中执行select子句或操作表的顺序
					id相同,执行顺序由上至下;
					id不同,如果是子查询,id的序号会递增,id值越大优先级越高,越先被执行;
					id相同不同,都存在
				select_type：表示查询的类型，主要用于区别普通查询，联合查询，子查询等复杂查询
					https://www.cnblogs.com/danhuangpai/p/8475458.html
						SIMPLE:简单的select查询，查询中不包括子查询或者UNION
						PRIMARY：查询中若包括任何复杂的子部分，最外层查询则被标记为PRIMARY
						SUBQUERY:在select或where列表中包含了子查询
						DERIVED: 在FROM列表中,包含的子查询被标记为DERIVED(衍生),MySQL会递归执行这些子查询,
						把结果放在临时表里
						UNION: 若第二个SELECT出现在UNION之后,则被标记为UNION;若UNION包含在FROM子句的子查询中,
							外侧SELECT将被标记为 DERIVE
								UNION RESULT: 从UNION表获取结果的SELECT
				table：查询的表
					type：在表中找到所需行的方式  ALL、index、range、 ref、eq_ref、const、system、NULL（从左到右，性能从差到好）
						ALL：全部扫描，效率低
						index：Full Index Scan，index与ALL区别为index类型只遍历索引树
						explain select * from film; file中有name字段为索引
						
						range：只检索给定范围的行，使用一个索引来选择行
						explain select * from employee where rec_id < 3其中rec_id为主键
						
						ref：表示上述表的连接匹配条件，即哪些列或常量被用于查找索引列上的值
						查找条件列使用了索引而且不为主键和unique。其实，意思就是虽然使用了索引，但该索引列的值并不唯一，有重复。这样即使使用索引快速查找到了第一条数据，仍然不能停止，要进行目标值附近的小范围扫描。但它的好处是它并不需要扫全表，因为索引是有序的，即便有重复值，也是在一个非常小的范围内扫描。下面为了演示这种情形，给employee表中的name列添加一个普通的key（值允许重复）
						explain select * from film where name = "film1" ref name是film中的普通索引
						
						eq_ref：类似ref，区别就在使用的索引是唯一索引，对于每个索引键值，表中只有一条记录匹配，简单来说，就是多表连接中使用primary key或者 unique key作为关联条件
						eq_ref 与 ref相比牛的地方是，它知道这种类型的查找结果集只有一个？什么情况下结果集只有一个呢！那便是使用了主键或者唯一性索引进行查找的情况，比如根据学号查找某一学校的一名同学，在没有查找前我们就知道结果一定只有一个，所以当我们首次查找到这个学号，便立即停止了查询。这种连接类型每次都进行着精确查询，无需过多的扫描，因此查找效率更高，当然列的唯一性是需要根据实际情况决定的
						explain select * from film_actor left join film on film_actor.film_id= film.id eq_ref 根据主键或唯一索引查询film
						
						const、system：当MySQL对查询某部分进行优化，并转换为一个常量时，使用这些类型访问。如将主键置于where列表中，MySQL就能将该查询转换为一个常量，system是const类型的特例，当查询的表只有一行的情况下，使用system
						select 1 				from actor where id =1 const
						explain 				extended select * from (select * from film where id =1) tmp 临时表中只有一条记录，system
						
						NULL： MySQL在优化过程中分解语句，执行时甚至不用访问表或索引，例如从一个索引列里选取最小值可以通过单独索引查找完成explain select min(id) from film id是主键
					
					possible_keys：指出MySQL能使用哪个索引在表中找到记录，查询涉及到的字段上若存在索引，则该索引将被列出，但不一定被查询使用（该查询可以利用的索引，如果没有任何索引显示 null），可能出现 possible_keys 有列，而 key 显示 NULL 的情况，这种情况是因为表中数据不多，mys ql认为索引对此查询帮助不大，选择了全表查询
					
					key：显示MySQL实际决定使用的键（索引)
					
					key_len：表示索引中使用的字节数，可通过该列计算查询中使用的索引的长度（key_len显示的值为索引字段的最大可能长度，
					并非实际使用长度，即key_len是根据表定义计算而得，不是通过表内检索出的）不损失精确性的情况下，长度越短越好
					key_len计算规则如下：

					字符串
					char(n)：n字节长度
					varchar(n)：2字节存储字符串长度，如果是utf-8，则长度3n + 2
					数值类型
					tinyint：1字节
					smallint：2字节
					int：4字节
					bigint：8字节 
					时间类型 
					date：3字节timestamp：4字节
					datetime：8字节如果字段允许为 NULL，需要1字节记录是否为 NULL
					索引最大长度是768字节，当字符串过长时，mysql会做一个类似左前缀索引的处理，将前半部分的字符提取出来做索引。
					
					ref：列与索引的比较，表示上述表的连接匹配条件，即哪些列或常量被用于查找索引列上的值
					rows：估算出结果集行数，表示MySQL根据表统计信息及索引选用情况，估算的找到所需的记录所需要读取的行数
					filtered：一个半分比的值，rows * filtered/100 可以估算出将要和 explain 中前一个表进行连接的行数（前一个表指 explain 中的id值比当前表id值小的表）
					Extra
						Using where:不用读取表中所有信息，仅通过索引就可以获取所需数据，
						这发生在对表的全部的请求列都是同一个索引的部分的时候，表示mysql服务器将在存储引擎检索行后再进行过滤
						Using temporary：表示MySQL需要使用临时表来存储结果集，常见于排序和分组查询，常见 group by ; order by;distinct
						Using filesort：当Query中包含 order by 操作，而且无法利用索引完成的排序操作称为“文件排序”
						Using join buffer：改值强调了在获取连接条件时没有使用索引，并且需要连接缓冲区来存储中间结果。
						如果出现了这个值，那应该注意，根据查询的具体情况可能需要添加索引来改进能
						Impossible where：这个值强调了where语句会导致没有符合条件的行（通过收集统计信息不可能存在结果）
						Select tables optimized away：这个值意味着仅通过使用索引，优化器可能仅从聚合函数结果中返回一行
						No tables used：Query语句中使用from dual 或不含任何from子句
				缺点：
				• EXPLAIN不会告诉你关于触发器、存储过程的信息或用户自定义函数对查询的影响情况
				• EXPLAIN不考虑各种Cache
				• EXPLAIN不能显示MySQL在执行查询时所作的优化工作
				• 部分统计信息是估算的，并非精确值
				• EXPLAIN只能解释SELECT操作，其他操作要重写为SELECT后查看执行计划
			SHOW WARNINGS 在explain执行后执行，查看翻译后的sql
			最佳左前缀法则
				如果索引了多列，要遵守最左前缀法则。指的是查询从索引的最左前列开始并且不跳过索引中的列
			索引优化规则
				索引基数
					如果你的索引字段重复率太高,索引基数/表数据总数<=0.3,mysql优化器将不走索引
				索引回表
					如果select * from table where column=xx,
					column不为主键时,每次查询将会额外根据主键查询一次数据,这个操作称为:回表
					如果没有其他字段要求,可以select column from table where column
				order by column时,可能会因为回表的原因,mysql优化器不走索引,如果没有必要,尽量加上limit限制数量
				
				不在索引列上做任何操作（计算、函数、（自动or手动）类型转换），会导致索引失效而转向全表扫描
				EXPLAIN SELECT * FROM employees WHERE name = 'LiLei'; 使用索引
				EXPLAIN SELECT * FROM employees WHERE left(name,3) = 'LiLei'; 未使用索引

				使用索引中范围条件右边的列时,可能会因为数据量过多,导致不能使用索引,但不是一定
				EXPLAIN SELECT * FROM employees WHERE name= 'LiLei' AND age = 22 AND position ='manager'; 使用索引
				EXPLAIN SELECT * FROM employees WHERE name= 'LiLei' AND age > 22 AND position ='manager'; 未使用索引

				尽量使用覆盖索引（只访问索引的查询（索引列包含查询列）），减少select *语句
				EXPLAIN SELECT name,age FROM employees WHERE name= 'LiLei' AND age = 23 AND position ='manager'; 只查询索引不用查询具体的数据，效率更高
				EXPLAIN SELECT * FROM employees WHERE name= 'LiLei' AND age = 23 AND position ='manager';

				mysql在使用不等于（！=或者<>）的时候可能会因为数据量过多,导致不能使用索引,但不是一定

				is null,is not null 也无法使用索引  的时候可能会因为数据量过多,导致不能使用索引,但不是一定

				like以通配符开头（'%abc...'）mysql索引失效会变成全表扫描操作

				字符串不加单引号索引失效
				EXPLAIN SELECT * FROM employees WHERE name = '1000';
				EXPLAIN SELECT * FROM employees WHERE name = 1000;
				
				or 只有两边都有索引才走索引，如果都没有或者只有一个是不走索引的
				
				in操作能避免则避免，若实在避免不了，需要仔细评估in后边的集合元素数量，控制在1000个之内,如果太多了,会导致sql过长导致mysql拒绝处理

				union all 不去重复，union去重复，union使用了临时表，应尽量避免使用临时表

				order by如果根据多个值进行排序，那么排序方式必须保持一致，要么同时升续，要么同时降续，排序方式不一致不走索引
		
		优化方式
			优化数据库表结构的设计
				字段的数据类型
					不同的数据类型的存储和检索方式不同，对应的性能也不同，所以说要合理的选用字段的数据类型。比如人的年龄用无符号的unsigned tinyint即可，没必要用integer
				数据类型的长度
					数据库最终要写到磁盘上，所以字段的长度也会影响着磁盘的I/O操作，如果字段的长度很大，那么读取数据也需要更多的I/O, 所以合理的字段长度也能提升数据库的性能。比如用户的手机号11位长度，没必要用255个长度
				表的存储引擎
			分库分表
				分库
				分表
				垂直分表
				水平分表
			数据库参数配置优化
			主从复制，读写分离
			数据库编码: 采用utf8mb4而不使用utf8
			字段名
				MySQL 在 Windows 下不区分大小写，但在 Linux 下默认是区分大小写。因此，数据库名、 表名、字段名，都不允许出现任何大写字母，避免节外生枝。
				一般所有表都要有id, id必为主键，类型为bigint unsigned,单表时自增、步长为1; 有些特殊场景下(如在高并发的情况下该字段的自增可能对效率有比价大的影响)id是通过程序计算出来的一个唯一值而不是通过数据库自增长来实现的。
				一般情况下主键id和业务没关系的，例如订单号不是主键id，一般是订单表中的其他字段，一般订单号order_code为字符类型
				一般情况下每张表都有着四个字段create_id,create_time,update_id,update_time, 其中create_id表示创建者id，create_time表示创建时间，update_id表示更新者id，update_time表示更是时间，这四个字段的作用是为了能够追踪数据的来源和修改
				最好不要使用备用字段(个人观点), 禁用保留字，如 desc、range、match、delayed 等

				表达是与否概念的字段，必须使用 is_xxx 的方式命名，数据类型是 unsigned tinyint (1 表示是，0 表示否), 任何字段如果为非负数，必须是unsigned。表达逻辑删除的字段名 is_deleted，1 表示删除，0 表示未删除
				如果某个值能通过其他字段能计算出来就不需要用个字段来存储，减少存储的数据
				为了提高查询效率，可以适当的数据冗余，注意是适当
				
				强烈建议不使用外键, 数据的完整性靠程序来保证
				单条记录大小禁止超过8k， 一方面字段不要太多，有的都能上百，甚至几百个，另一方面字段的内容不易过大像文章内容等这种超长内容的需要单独存到另一张表
			
			字段类型
				字符类型
					不同存储引擎对char和varchar的使用原则不同，myisam:建议使用国定长度的数据列代替可变长度。
					innodb：建议使用varchar，大部分表都是使用innodb，所以varchar的使用频率更高
				数值类型
					金额类型的字段尽量使用long用分表示，尽量不要使用bigdecimal，严谨使用float和double因为计算时会丢失经度
					如果需要使用小数严谨使用float，double，使用定点数decimal，decimal实际上是以字符串的形式存储的，所以更加精确，java中与之对应的数据类型为BigDecimal
					如果值为非负数，一定要使用unsigned，无符号不仅能防止负数非法数据的保存，而且还能增大存储的范围
					不建议使用ENUM、SET类型，使用TINYINT来代替
				日期类型
					根据实际需要选择能够满足应用的最小存储日期类型。

					如果应用只需要记录年份，那么仅用一个字节的year类型。
					如果记录年月日用date类型, date占用4个字节，存储范围10000-01-01到9999-12-31
					如果记录时间时分秒使用它time类型
					如果记录年月日并且记录的年份比较久远选择datetime，而不要使用timestamp,因为timestamp表示的日期范围要比datetime短很多
					如果记录的日期需要让不同时区的用户使用，那么最好使用timestamp, 因为日期类型值只有它能够和实际时区相对应
					datetime默认存储年月日时分秒不存储毫秒fraction，如果需要存储毫秒需要定义它的宽度datetime(6)
					timestamp与datetime

					两者都可用来表示YYYY-MM-DD HH:MM:SS[.fraction]类型的日期。

					都可以使用自动更新CURRENT_TIMESTAMP

					对于TIMESTAMP，它把客户端插入的时间从当前时区转化为UTC（世界标准时间）进行存储。查询时，将其又转化为客户端当前时区进行返回。而对于DATETIME，不做任何改变，基本上是原样输入和输出。

					timestamp占用4个字节：timestamp所能存储的时间范围为：’1970-01-01 00:00:01.000000’ 到 ‘2038-01-19 03:14:07.999999’
					datetime占用8个字节 ：datetime所能存储的时间范围为：’1000-01-01 00:00:00.000000’ 到 ‘9999-12-31 23:59:59.999999’

					总结：TIMESTAMP和DATETIME除了存储范围和存储方式不一样，没有太大区别。如果需要使用到时区就必须使用timestamp，如果不使用时区就使用datetime因为datetime存储的时间范围更大

					注意：

					禁止使用字符串存储日期，一般来说日期类型比字符串类型占用的空间小，日期时间类型在进行查找过滤是可以利用日期进行对比，这比字符串对比高效多了，日期时间类型有丰富的处理函数，可以方便的对日期类型进行日期的计算
					也尽量不要使用int来存储时间戳
					
				是否为null
					MySQL字段属性应该尽量设置为NOT NULL，除非你有一个很特别的原因去使用 NULL 值，你应该总是让你的字段保持 NOT NULL 。

					在MySql中NULL其实是占用空间的，“可空列需要更多的存储空间”：需要一个额外字节作为判断是否为NULL的标志位“需要mysql内部进行特殊处理”， 而空值”“是不占用空间的。

					含有空值的列很难进行查询优化，而且对表索引时不会存储NULL值的，所以如果索引的字段可以为NULL，索引的效率会下降很多。因为它们使得索引、索引的统计信息以及比较运算更加复杂。你应该用0、一个特殊的值或者一个空串代替null。

					联表查询的时候，例如SELECT user.username, info.introduction FROM tbl_user user LEFT JOIN tbl_userinfo info ON user.id = info.user_id; 如果tbl_userinfo.introduction设置的可以为null, 假如这条sql查询出了对应的记录，但是username有值，introduction没有值，那么就不是很清楚这个introduction是没有关联到对应的记录，还是关联上了而这个值为null，null意思表示不明确，有歧义

					注意：NULL在数据库里是非常特殊的，任何数跟NULL进行运算都是NULL, 判断值是否等于NULL，不能简单用=，而要用IS NULL关键字。使用 ISNULL()来判断是否为 NULL 值，NULL 与任何值的直接比较都为 NULL。

					1) NULL<>NULL的返回结果是NULL，而不是false。
					2) NULL=NULL的返回结果是NULL，而不是true。
					3) NULL<>1的返回结果是NULL，而不是true。
			实例
				EXPLAIN SELECT * FROM tbl_user LIMIT 100000,2;
				EXPLAIN SELECT * FROM tbl_user u INNER JOIN (SELECT id FROM tbl_user ORDER BY id ASC LIMIT 10000,2) temp ON u.id = temp.id;
				id为主键，性能高于第一条全表扫描

				where中如果有多个过滤条件，在没有索引的情况下将过滤多的写在前面，过滤少的写在后面
			
			尽量少使用select *，需要什么字段就去取哪些字段

			不要使用count(列名)或 count(常量)来替代 count()，count()是SQL92定义的标准统计行数的语法，跟数据库无关，跟 NULL和非NULL无关。 说明:count(*)会统计值为NULL 的行，而count(列名)不会统计此列为NULL值的行

			禁止使用存储过程，存储过程难以调试和扩展，更没有移植性。避免使用存储过程、触发器

			除了 IO 瓶颈之外，SQL优化中需要考虑的就是 CPU 运算量的优化了。order by, group by,distinct … 都是消耗 CPU 的大户（这些操作基本上都是 CPU 处理内存中的数据比较运算）。当我们的 IO 优化做到一定阶段之后，降低 CPU 计算也就成为了我们 SQL 优化的重要目标
	
	锁
		分类
			乐观锁和悲观锁
				乐观锁
					根据版本号控制
				悲观锁
					锁定表或者行，让其他数据操作等待
						读锁(共享锁)
							针对同一份数据，多个读操作可以同时进行而不会互相影响
							不能进行写操作
						写锁(排他锁)
							当前写操作没有完成前，它会阻断其他写锁和读锁
			表锁和行锁
				表锁
					表锁偏向MyISAM存储引擎，开销小，加锁快，无思索，锁定粒度大，发生锁冲突的概率最高，并发度最低
					
					当前session对该表的增删改查都没有问题，其他session对该表的所有操作被阻塞 
					
						lock table表名称 read(write)
						unlock tables
							MyISAM在执行查询语句(SELECT)前,会自动给涉及的所有表加读锁,在执行增删改操作前,会自动给涉及的表加写锁。1、对MyISAM表的读操作(加读锁) ,不会阻寒其他进程对同一表的读请求,但会阻赛对同一表的写请求。只有当读锁释放后,才会执行其它进程的写操作。2、对MylSAM表的写操作(加写锁) ,会阻塞其他进程对同一表的读和写操作,只有当写锁释放后,才会执行其它进程的读写操作总结：简而言之，就是读锁会阻塞写，但是不会阻塞读。而写锁则会把读和写都阻塞

				行锁
					行锁偏向InnoDB存储引擎，开销大，加锁慢，会出现死锁，锁定粒度最小，发生锁冲突的概率最低，并发度也最高。InnoDB与MYISAM的最大不同有两点：一是支持事务（TRANSACTION）；二是采用了行级锁

		行锁分析
			show status like'innodb_row_lock%';
				Innodb_row_ lock_current_wait
					当前正在等待锁定的数量
				Innodb_row_ lock_time
					从系统启动到现在锁定总时间长度
				Innodb_row_ lock_time_avg
					每次等待所花平均时间
				Innodb_row_ lock_time_max
					从系统启动到现在等待最长的一次所花时间
				Innodb_row_ lock_waits
					系统启动后到现在总共等待的次数
		死锁
			Session _1执行：select *from account where i d= 1 for update;
			Session _2执行：select *from account where i d= 2 for update;
			Session _1执行：select *from account where i d= 2 for update;
			Session _2执行：select *from account where i d= 1 for update;
			查看近期死锁日志信息：show engine inno db statu s\G;

	表数据操作
		增
			添加单条
				INSERT INTO tb_name(`field1`,`field2`,....)VALUES('value1','value2',.....);
			添加多条
				INSERT INTO tb_name(`field1`,`field2`,....)VALUES('value1','value2',.....),('value1','value2',.....),('value1','value2',.....),....;
		删
			sql
				DELETE FROM tb_name WHERE ...
			注意
				删除时必须加where条件
		改
			sql
				UPDATE tb_name SET field1 = value1,field2 = value2,.....   WHERE  ....
			注意
				修改时必须加where条件
		查
			基础的查询
				SELECT * FROM tb_name
			where子句
				比较运算符
					大于、小于、等于、不等于、大于等于、小于等于
					SELECT * FROM tb_name WHERE user_id >10;
				逻辑运算符
					逻辑运算符是用来拼接其他条件的。用and或者or来连接两个条件，如果用or来连接的时候必须使用小括号
					SELECT * FROM tb_name WHERE user_id > 10 AND sex = '男'
			LIKE模糊查询
				通配符
					%（百分号）匹配零个或者多个任意字符
					_（下划线）匹配一个任意字符
			sql
				SELECT * FROM tb_name WHERE username LIKE '张%';查找username开头是张的数据
				SELECT * FROM tb_name WHERE username LIKE '%张%';查询username中含有张的数据
				SELECT * FROM tb_name WHERE username LIKE '%张';查询username字段的数据以张结尾的
				SELECT   *  FROM  tb_name WHERE username LIKE '张_';查询username以张开头后边有一个字符的数据
			IN字段指定多个值查询
				IN (value1,value2,value3,....)
				SELECT   *   FROM   tb_name  WHERE   user_id  IN (1,3,5,7,9,11);查询user_id是1，3，5，7，9，11的所有数据
			BETWEEN AND 区间查询
				field  BETWEEN  value1  AND   value2
				SELECT * FROM user WHERE user_id  BETWEEN  2 AND 9;查询user表中user_id大于等于2小于等于9的所有值
			GROUP BY 分组查询
				配合函数
					count(field)获取符合条件出现的非null值的次数
					SUM(field)获取所有符合条件的数据的总和
					AVG(field)或者平均值
				SELECT  sex,COUNT(*)  count  FROM class  GROUP BY sex;获取class表中男生和女生的数量
				group by会创建临时表
			ORDER BY 查询排序
				查询顺序
					ORDER BY field DESC;降序查询
					ORDER BY field ASC;升序查询
				SELECT  *  FROM  tb_name  ORDER BY   id  DESC; 查询tb_name表中所有数据，按id的降序来查找
				mysql有两种排序方式：

				通过有序索引顺序扫描直接返回有序数据，通过explain分析显示Using Index,不需要额外的排序，操作效率比较高。
				通过对返回数据进行排序，也就是Filesort排序，所有不是通过索引直接返回排序结果的都叫Filesort排序。Filesort是通过相应的排序算法将取得的数据在sort_buffer_size系统变量设置的内存排序中进行排序，如果内存装载不下，就会将磁盘上的数据进行分块，再对各个数据块进行排序，然后将各个块合并成有序的结果集

				order by 使用索引的严格要求：

					索引的顺序和order by子句的顺序完全一致
					索引中所有列的方向(升续、降续)和order by 子句完全一致
					当多表连接查询时order by中的字段必须在关联表中的第一张表中
					
			LIMIT 查询结果截取
				参数
					LIMIT 后边可以跟两个参数，如果只写一个表示从零开始查询指定长度，如果两个参数就是从第一个参数开始查询查询长度是第二个参数的值，俩个参数必须是整形。
				SELECT   *   FROM   tb_name   LIMIT   5;查询tb_name表中的所有数据，只要前边的5条数据
				SELECT    *   FROM   tb_name   LIMIT   5,5;查询tb_name中所有的数据，返回的结果是从第五条开始截取五条数据
				分页查询一般会全表扫描，优化的目的应尽可能减少扫描；

					第一种思路：在索引上完成排序分页的操作，最后根据主键关联回原表查询原来所需要的其他列。这种思路是使用覆盖索引尽快定位出需要的记录的id，覆盖索引效率高些

					第二中思路：limit m,n 转换为 n

				之前分页查询是传pageNo页码, pageSize分页数量，
				当前页的最后一行对应的id即last_row_id，以及pageSize，这样先根据条件过滤掉last_row_id之前的数据，然后再去n挑记录,此种方式只能用于排序字段不重复唯一的列，如果用于重复的列，那么分页数据将不准确
		关联查询
			外关联
				
				SELECT  *   FORM  tb_name1   LEFT  JOIN  tb_name2  ON  tb_name1.t2_id  =  tb_name2.t2_id;用表一的t2_id和表二的t2_id来关联，查询所有的值。
				FROM 之后的表是主表
				
			左关联				

		
			中关联
				SELECT   *    FORM  tb_name1  JOIN   tb_name2   ON   tb_name1.t2_id  =  tb_name2.t2_id;用表一的t2_id和表二的t2_id来关联，查询所有的值。
				中关联没有主表
			右关联
				SELECT   *    FORM  tb_name1   RIGHT  JOIN   tb_name2   ON   tb_name1.t2_id  =  tb_name2.t2_id;用表一的t2_id和表二的t2_id来关联，查询所有的值。
				在ON后边的表是主表
			内关联
				内管理的关联条件是用where来说关联的，多张表之间用AND来拼接where条件
				SELECT   *   FROM   tb_name1,tb_name2,....   WHERE   tb_name1.t2_id = tb_name2.t2_id   AND  ...
			外关联的说明
				主表关联副表，如果副表数据不够用NULL来补全，但是中关联的时候，如果不够了，左边的数据或者右边的数据不会显示。直接去掉
			横向连接：两个表字段一样，数据合并
				union会去重，union all不会
		事务
			关键词
				BEGIN开启事务
				ROLLBACK;事务回滚
				COMMIT;事务提交
			必备条件
				表的引擎为InnoDB
	数据库导入导出
		数据库导出
			1.打开cmd命令
			2.打开到mysql文件夹下的bin目录
			3.通过mysqldump来执行导出
			4.命令：mysqldump -u root -p 数据库（class15） >  要导出的文件名如：（test.sql）
			5.导出之后的文件会出现在bin目录下
		sql文件导入
			1.cmd打开到mysql的bin目录下
			2.通过   mysql -uroot -p 输入密码的形式进入到数据库中
			3.选择数据库   USE  db_name;
			4.执行导入命令：   source d:\datafilename.sql   后边路径是sql文件存放的物理路径
	引擎
		innoDB
			innodb主键使用自增bigint效率比uuid高
			1.方便比较大小
			2.不会破坏B+TREE结构
			聚集索引：索引和数据在同一张表
			非聚集索引：索引在一张表，数据在一张表
			innodb使用b+tree存索引和数据 不使用hash的原因：
			范围查找使用hash不合适，需要全表扫描，hash(主键)直接存储到位置，因此一般使用B+Tree
			支持事务
			支持行锁
		myisam
			myisam使用非聚集索引,主键和其他索引都是指向数据表
			不支持事务
			只支持表锁
			show engines查看所有引擎
	事务
		ACID属性
			持久性(Durable)
				事务完成之后,它对于数据的修改是永久性的,即使出现系统故障也能够保持
			原子性(Atomicity)
				事务是一个原子操作单元,其对数据的修改,要么全都执行,要么全都不执行
			隔离性(Isolation)
				数据库系统提供一定的隔离机制,保证事务在不受外部并发操作影响的“独立”环境执行。
				这意味着事务处理过程中的中间状态对外部是不可见的,反之亦然
			一致性(Consistent)
				在事务开始和完成时,数据都必须保持一致状态。这意味着所有相关的数据规则都必须应用于事务的修改,
				以保持数据的完整性;事务结束时,所有的内部数据结构(如B树索引或双向链表)也都必须是正确的
		并发事务处理带来的问题
			更新丢失（Lost Update）
				两个或多个事务选择同一行，然后基于最初选定的值更新该行时，由于每个事务都不知道其他事务的存在，
				就会发生丢失更新问题–最后的更新覆盖了由其他事务所做的更新
			脏读（Dirty Reads）
				事务A读取到了事务B已经修改但尚未提交的数据，还在这个数据基础上做了操作。
				此时，如果B事务回滚，A读取的数据无效，不符合一致性要求
			不可重读（Non-Repeatable Reads）
				一个事务在读取某些数据后的某个时间，再次读取以前读过的数据，
				却发现其读出的数据已经发生了改变、
				或某些记录已经被删除了！这种现象就叫做“不可重复读”。
				一句话：事务A读取到了事务B已经提交的修改数据，不符合隔离性
			幻读（Phantom Reads）
				个事务按相同的查询条件重新读取以前检索过的数据，却发现其他事务插入了满足其查询条件的新数据，
				这种现象就称为“幻读”
			脏读是事务B里面修改了数据,幻读是事务B里面新增了数据
		事务隔离级别
			分类
				读未提交(Read uncommitted)
					脏读/不可重复读/幻读都可能
				读已提交(Read uncommitted)
					脏读不可能，不可重复读/幻读都可能
				可重复读(Repeatable read)
					不可重复读/脏读不可能，幻读都可能
						可重复读的隔离级别下使用了MVCC机制，select操作不会更新版本号，是快照读（历史版本）
						insert、update和delete会更新版本号，是当前读（当前版本）
						要避免幻读可以用间隙锁在Session _1下面执行update acc ount se t name ='zhuge'
						where i d> 10and id<= 20;，则其他Session没法插入这个范围内的数据
				可串行化(Serializable)
					脏读/不可重复读/幻读都不可能
						mysql中事务隔离级别为serializable时会锁表，因此不会出现幻读的情况，
						这种隔离级别并发性极低，开发中很少会用到
			查看事务隔离级别
				show variables like 'transaction_isolation';
				select @@transaction_isolation;
			默认的事务隔离级别(Repeatable read)

	分表
		横向：表字段相同，数据量太大
			union会去重，union all不会
		纵向：一个表存基本信息，另外一个表存详情
			join

	分区
		优缺点
			优点
				和单个磁盘或者文件系统分区相比，可以存储更多数据
				优化查询。在where子句中包含分区条件时，可以只扫描必要的一个或者多个分区来提高查询效率；
				同时在涉及sum()和count()这类聚合函数的查询时，可以容易的在每个分区上并行处理，最终只需要汇总所有分区得到的结果
				对于已经过期或者不需要保存的数据，可以通过删除与这些数据有关的分区来快速删除数据
				跨多个磁盘来分散数据查询，以获得更大的查询吞吐量
		分区方式
			Range
				基于属于一个给定连续区间的列值，把多行分配给分区。
				这些区间要连续且不能相互重叠，使用VALUES LESS THAN操作符来进行定义
					CREATE TABLE employees (
					id INT NOT NULL,
					NAME VARCHAR ( 30 ),
					hired DATE NOT NULL DEFAULT '2018-12-01',
					job VARCHAR ( 30 ) NOT NULL,
					dept_id INT NOT NULL
					)
					PARTITION BY RANGE ( dept_id )(
					PARTITION p0 VALUES LESS THAN ( 6 ),
					PARTITION p1 VALUES LESS THAN ( 11 ),
					PARTITION p2 VALUES LESS THAN ( 16 ),
					PARTITION p3 VALUES LESS THAN ( 21 ) );
						1）、当需要删除一个分区上的"旧的"数据时,只删除分区即可。如果你使用上面最近的那个例 子给出 的分区方案，你只需简单地使用"ALTER TABLE employees DROP PARTITION p0；"来删除所有在1991年前就已经停止工作的雇员相对应的所有行。对于有大量行的表，这比 运行一个如“DELETE FROM employees WHERE YEAR (separated) <= 1990；”这样的一个DELETE查询要有效得多。
						2）、想要使用一个包含有日期或时间值，或包含有从一些其他级数开始增长的值的列。
						3）、 经常运行直接依赖于用于 分割 表的 列的查询。例如，当执行一个如“SELECT COUNT(*) FROM employees WHERE YEAR(separated) = 2000 GROUP BY dept_id；”这样的查询时，MySQL可以很 迅速地确定只有分区p2需要扫描 ，这是因为余下的分区不可能包含有符合该WHERE子句的任何记录。
			List
				设置若干个固定值进行分区，如果某个字段的值在这个设置的值列表中就会被分配到该分区。
				适用于字段的值区分度不高的，或者值是有限的，特别是像枚举这样特点的列
					CREATE TABLE employees (
							 id INT NOT NULL,
							 NAME VARCHAR ( 30 ),
							 hired DATE NOT NULL DEFAULT '2015-12-10',
							 job_code INT,
							 store_id INT )
					PARTITION BY LIST ( store_id )(
					PARTITION pQY VALUES IN ( 3, 5, 6, 17 ),
					PARTITION pJN VALUES IN ( 1, 10, 11, 19 ),
					PARTITION pCH VALUES IN ( 4, 12, 14, 18 ),
					PARTITION pJJ VALUES IN ( 2, 9, 13, 16 ),
					PARTITION pGX VALUES IN ( 7, 8, 15, 20 ));
					
			range columns
			
				create table rc3 (
				a int,
				b int
				)
				partition by range columns(a, b) (
				partition p01 values less than (0, 10),
				partition p02 values less than (10, 10),
				partition p03 values less than (10, 20),
				partition p04 values less than (10, 35)
				);
				insert into rc3(a, b) values(1, 10);

			hash分区			
				常规hash分区				
					常规hash分区使用的是取模算法，对应一个表达式expr是可以计算出它被保存到哪个分区中，N = MOD(expr, num)
				线性hash分区
					线性hash分区使用的是一个线性的2的幂运算法则
						常规hash分区在管理上带来了的代价太大，不适合需要灵活变动分区的需求。为了降低分区管理上的代价，mysql提供了线性hash分区，分区函数是一个线性的2的幂的运算法则。同样线性hash分区的记录被存在那个分区也是能被计算出来的。线性hash分区的优点是在分区维护(增加、删除、合并、拆分分区)时，mysql能够处理的更加迅速，缺点是：对比常规hash分区，线性hash各个分区之间数据的分布不太均衡
			key分区
				按照key进行分区非常类似于按照hash进行分区，只不过hash分区允许使用用户自定义的表达式，
				而key分区不允许使用用于自定义的表达式，需要使用mysql服务器提供的hash函数，
				同时hash分区只支持整数分区，而key分区支持使用出blob or text类型外的其他类型的列作为分区键
					
					partition by key(expr) partitions num;
					-- 不指定默认首选主键作为分区键，在没有主键的情况下会选择非空唯一键作为分区键
					partition by key() partitions num;
					-- linear key
					partition by linear key(expr)
						
			子分区
				是分区表中对每个分区的再次分割，又被称为复合分区，支持对range和list进行子分区，
					子分区即可以使用hash分区也可以使用key分区。
					复合分区适用于保存非常大量的数据记录
						create table ts (
						id int,
						purchased date
						)
						partition by range(year(purchased))
						subpartition by hash(to_days(purchased)) subpartitions 2
						(
						partition p0 values less than (1990),
						partition p0 values less than (2000),
						partition p0 values less than maxvalue
						);

			管理分区
				-- 删除list或者range分区(同时删除分区对应的数据)
					alter table <table> drop partition <分区名称>;

				-- 新增分区
				-- range添加新分区
				alter table <table> add partition(partition p4 values less than MAXVALUE);

				-- list添加新分区
				alter table <table> add partition(partition p4 values in (25,26,28));

				-- hash重新分区
				alter table <table> add partition partitions 4;

				-- key重新分区
				alter table <table> add partition partitions 4;

				-- 子分区添加新分区，虽然我没有指定子分区，但是系统会给子分区命名的
				alter table <table> add partition(partition p3 values less than MAXVALUE);

				-- range重新分区
				ALTER TABLE user REORGANIZE PARTITION p0,p1,p2,p3,p4 INTO (PARTITION p0 VALUES LESS THAN MAXVALUE);

				-- list重新分区
				ALTER TABLE <table> REORGANIZE PARTITION p0,p1,p2,p3,p4 INTO (PARTITION p0 VALUES in (1,2,3,4,5));

	数据库配置 my.cnf
		客户端配置
			client mysql所有客户端执行时配置
				socket
					执行mysql需要连接的socket
				host
					执行mysq 需要连接的ip,跟socket互斥
				port
					执行mysql的port
				password
					执行mysql 时默认的密码
			mysql
				no-auto-rehash
					关闭命令补全
				auto-rehash
					开启命令补全
				myisamchk
				mysqldump
					导出项
						all-databases
							导出全部数据库 
						all-tablespaces/no-tablespaces
							导出全部表空间/不导出全部表空间信息
						add-drop-database
							每个数据库创建之前添加drop数据库语句。
						add-drop-table
							每个数据表创建之前添加drop数据表语句。
						add-drop-trigger
							每个触发器创建之前添加drop语句
						add-locks
							在每个表导出之前增加LOCK TABLES并且之后UNLOCK TABLE。
						allow-keywords
							允许创建是关键词的列名字。这由表名前缀于每个列名做到。
						apply-slave-statements
							在'CHANGE MASTER'前添加'STOP SLAVE'，并且在导出的最后添加'START SLAVE'。
						character-sets-dir
							字符集文件的目录
						comments
							附加注释信息。默认为打开，
						compatible
							导出的数据将和其它数据库或旧版本的MySQL 相兼容。值可以为ansi、mysql323、mysql40、postgresql、oracle、mssql、db2、maxdb、no_key_options、no_tables_options、no_field_options等，
							要使用几个值，用逗号将它们隔开。它并不保证能完全兼容，而是尽量兼容。
						compact
							导出更少的输出信息(用于调试)。去掉注释和头尾等结构。
						complete-insert
							使用完整的insert语句(包含列名称)。这么做能提高插入效率，但是可能会受到max_allowed_packet参数的影响而导致插入失败。
						create-options
							在CREATE TABLE语句中包括所有MySQL特性选项。(默认为打开状态)
						databases
							导出几个数据库。参数后面所有名字参量都被看作数据库名。
						default-character-set
							设置默认字符集，默认值为utf8
						delete-master-logs
							master备份后删除日志. 这个参数将自动激活--master-data。
						disable-keys
							对于每个表，用/*!40000 ALTER TABLE tbl_name DISABLE KEYS */;和/*!40000 ALTER TABLE tbl_name ENABLE KEYS */;语句引用INSERT语句。这样可以更快地导入dump出来的文件，因为它是在插入所有行后创建索引的。该选项只适合MyISAM表，默认为打开状态。
						dump-slave
							该选项将导致主的binlog位置和文件名追加到导出数据的文件中。设置为1时，将会以CHANGE MASTER命令输出到数据文件；设置为2时，在命令前增加说明信息。该选项将会打开--lock-all-tables，除非--single-transaction被指定。该选项会自动关闭--lock-tables选项。默认值为0。
						events
							导出事件。
						extended-insert
							使用具有多个VALUES列的INSERT语法。这样使导出文件更小，并加速导入时的速度。默认为打开状态，使用--skip-extended-insert取消选项。
						fields-terminated-by
							导出文件中忽略给定字段。与--tab选项一起使用，不能用于--databases和--all-databases选项
						fields-enclosed-by
							输出文件中的各个字段用给定字符包裹。与--tab选项一起使用，不能用于--databases和--all-databases选项
						fields-optionally-enclosed-by
							输出文件中的各个字段用给定字符选择性包裹。与--tab选项一起使用，不能用于--databases和--all-databases选项
						fields-escaped-by
							输出文件中的各个字段忽略给定字符。与--tab选项一起使用，不能用于--databases和--all-databases选项
						flush-logs
							开始导出之前刷新日志。
							请注意：假如一次导出多个数据库(使用选项--databases或者--all-databases)，将会逐个数据库刷新日志。除使用--lock-all-tables或者--master-data外。在这种情况下，日志将会被刷新一次，相应的所以表同时被锁定。因此，如果打算同时导出和刷新日志应该使用--lock-all-tables 或者--master-data 和--flush-logs。
						flush-privileges
							在导出mysql数据库之后，发出一条FLUSH PRIVILEGES 语句。为了正确恢复，该选项应该用于导出mysql数据库和依赖mysql数据库数据的任何时候。
						force
							在导出过程中忽略出现的SQL错误。
						hex-blob
							使用十六进制格式导出二进制字符串字段。如果有二进制数据就必须使用该选项。影响到的字段类型有BINARY、VARBINARY、BLOB。
						host
							需要导出的主机信息
						ignore-error
							不导出指定表。指定忽略多个表时，需要重复多次，每次一个表。每个表必须同时指定数据库和表名。例如：--ignore-table=database.table1 --ignore-table=database.table2 ……
						include-master-host-port
							在--dump-slave产生的'CHANGE MASTER TO..'语句中增加'MASTER_HOST=<host>，MASTER_PORT=<port>'
						insert-ignore
							在插入行时使用INSERT IGNORE语句.
						lines-terminated-by
							输出文件的每行用给定字符串划分。与--tab选项一起使用，不能用于--databases和--all-databases选项。
						lock-all-tables
							提交请求锁定所有数据库中的所有表，以保证数据的一致性。这是一个全局读锁，并且自动关闭--single-transaction 和--lock-tables 选项。
						lock-tables
							开始导出前，锁定所有表。用READ LOCAL锁定表以允许MyISAM表并行插入。对于支持事务的表例如InnoDB和BDB，--single-transaction是一个更好的选择，因为它根本不需要锁定表。
						log-error
							附加警告和错误信息到给定文件
						master-data
							该选项将binlog的位置和文件名追加到输出文件中。如果为1，将会输出CHANGE MASTER 命令；如果为2，输出的CHANGE MASTER命令前添加注释信息。该选项将打开--lock-all-tables 选项，除非--single-transaction也被指定（在这种情况下，全局读锁在开始导出时获得很短的时间；其他内容参考下面的--single-transaction选项）。该选项自动关闭--lock-tables选项。
						no-autocommit
							使用autocommit/commit 语句包裹表。
						no-create-db
							只导出数据，而不添加CREATE DATABASE 语句。
						no-create-info
							只导出数据，而不添加CREATE TABLE 语句。
						no-data
							不导出任何数据，只导出数据库表结构。
						order-by-primary
							如果存在主键，或者第一个唯一键，对每个表的记录进行排序。在导出MyISAM表到InnoDB表时有效，但会使得导出工作花费很长时间。
						quick
							一行一行的读出数据并输出到标准输出中,不经过缓冲区
						quote-names
							使用（`）引起表和列名。默认为打开状态，使用--skip-quote-names取消该选项。
						replace
							使用REPLACE INTO 取代INSERT INTO.
						routines
							导出存储过程以及自定义函数。
						set-charset
							添加'SET NAMES default_character_set'到输出文件。默认为打开状态，使用--skip-set-charset关闭选项。
						single-transaction
							该选项在导出数据之前提交一个BEGIN SQL语句，BEGIN 不会阻塞任何应用程序且能保证导出时数据库的一致性状态。它只适用于多版本存储引擎，仅InnoDB。本选项和--lock-tables 选项是互斥的，因为LOCK TABLES 会使任何挂起的事务隐含提交。要想导出大表的话，应结合使用--quick 选项。
						dump-date
							将导出时间添加到输出文件中。默认为打开状态，使用--skip-dump-date关闭选项。
						tab
							为每个表在给定路径创建tab分割的文本文件。注意：仅仅用于mysqldump和mysqld服务器运行在相同机器上。
						triggers
							导出触发器。该选项默认启用，用--skip-triggers禁用它。
						tz-utc
							在导出顶部设置时区TIME_ZONE='+00:00' ，以保证在不同时区导出的TIMESTAMP 数据或者数据被移动其他时区时的正确性。
						where
							只转储给定的WHERE条件选择的记录。请注意如果条件包含命令解释符专用空格或字符，一定要将条件引用起来。
				连接项
					bind-address
						通过哪个ip来连接mysql服务器
					compress
						在客户端和服务器之间启用压缩传递所有信息
					max-allowed-packet
						服务器发送和接受的最大包长度。
					net-buffer-length
						TCP/IP和socket连接的缓存大小。
					port
						mysql服务器端口
					socket
						指定连接mysql的socket文件位置，默认路径/tmp/mysql.sock
					user
						指定连接的用户名。
				其他
					plugin-dir
						客户端插件的目录，用于兼容不同的插件版本。
					default-auth
						客户端插件默认使用权限。
					enable-cleartext-plugin
				ssl
					ssl
					ssl-verify-server-cert
					ssl-ca
					ssl-capath
					ssl-cert
					ssl-cipher
					ssl-key
					ssl-crl
					ssl-crlpath
					tls-version
					server-public-key-path
					get-server-public-key
		
		服务端配置		
			mysqld
				服务配置项
					bind-address
						mysql服务监听的端口
					server-id
						mysql本机的序号,集群下的唯一值
					port
						监听端口
					socket
						监听的socket
					pid-file
						mysql进程启动后的进程id存放位置
					basedir
						mysql的安装路径
					datadir
						mysql数据的存放路径
					tmpdir
						mysql临时文件的存放路径
					default-time-zone
						默认时区
				连接项
					skip-name-resolve 
						禁止 MySQL 对外部连接进行 DNS 解析，使用这一选项可以消除 MySQL 进行 DNS 解析的时间。但需要注意，如果开启该选项，则所有远程主机连接授权都要使用 IP 地址方式，否则 MySQL 将无法正常处理连接请求！
					skip-symbolic-links 
						不能使用连接文件，多个客户可能会访问同一个数据库，因此这防止外部客户锁定 MySQL 服务器。 该选项默认开启
					skip-external-locking 
						不使用系统锁定，要使用 myisamchk,必须关闭服务器 ,避免 MySQL的外部锁定，减少出错几率增强稳定性。
					skip-networking
						开启该选项可以彻底关闭 MySQL 的 TCP/IP 连接方式，如果 WEB 服务器是以远程连接的方式访问 MySQL 数据库服务器则不要开启该选项！否则将无法正常连接！ 如果所有的进程都是在同一台服务器连接到本地的 mysqld, 这样设置将是增强安全的方法
				系统资源项
					back_log
						接受队列，对于没建立 tcp 连接的请求队列放入缓存中，队列大小为 back_log，受限制与 OS 参数，试图设定 back_log 高于你的操作系统的限制将是无效的。默认值为 50。对于 Linux 系统推荐设置为小于512的整数。如果系统在一个短时间内有很多连接，则需要增大该参数的值
					max_connections
						指定MySQL允许的最大连接进程数。如果在访问数据库时经常出现"Too Many Connections"的错误提 示，则需要增大该参数值。
					max_connect_errors
						如果某个用户发起的连接 error 超过该数值，则该用户的下次连接将被阻塞，直到管理员执行 flush hosts ; 命令或者服务重启， 防止黑客 ， 非法的密码以及其他在链接时的错误会增加此值
					open_files_limit
						MySQL打开的文件描述符限制，默认最小1024;当open_files_limit没有被配置的时候，比较max_connections*5和ulimit-n的值，哪个大用哪个，当open_file_limit被配置的时候，比较open_files_limit和max_connections*5的值，哪个大用哪个。
					connect-timeout
						连接超时之前的最大秒数,在 Linux 平台上，该超时也用作等待服务器首次回应的时间
					wait-timeout
						等待关闭连接的时间
					interactive-timeout
						关闭连接之前，允许 interactive_timeout（取代了wait_timeout）秒的不活动时间。客户端的会话 wait_timeout 变量被设为会话interactive_timeout 变量的值。如果前端程序采用短连接，建议缩短这2个值, 如果前端程序采用长连接，可直接注释掉这两个选项，默认配置(8小时)  
					net_retry_count
						如果某个通信端口的读操作中断了，在放弃前重试多次
					net_buffer_length
						包消息缓冲区初始化为 net_buffer_length 字节，但需要时可以增长到 max_allowed_packet 字节
					max_allowed_packet
						服务所能处理的请求包的最大大小以及服务所能处理的最大的请求大小(当与大的BLOB 字段一起工作时相当必要)， 每个连接独立的大小.大小动态增加。 设置最大包,限制server接受的数据包大小，避免超长SQL的执行有问题 默认值为16M，当MySQL客户端或mysqld
						服务器收到大于 max_allowed_packet 字节的信息包时，将发出“信息包过大”错误，并关闭连接。对于某些客户端，如果通信信息包过大，在执行查询期间，可能会遇到“丢失与 MySQL 服务器的连接”错误。默认值 16M。
					table_cache
						所有线程所打开表的数量. 增加此值就增加了mysqld所需要的文件描述符的数量这样你需要确认在[mysqld_safe]中 “open-files-limit” 变量设置打开文件数量允许至少4096
					thread_stack
						线程使用的堆大小. 此容量的内存在每次连接时被预留.MySQL 本身常不会需要超过 64K 的内存如果你使用你自己的需要大量堆的 UDF 函数或者你的操作系统对于某些操作需要更多的堆,你也许需要将其设置的更高一点.默认设置足以满足大多数应用
					thread_cache_size
						我们在 cache 中保留多少线程用于重用.当一个客户端断开连接后,如果 cache 中的线程还少于 thread_cache_size,则客户端线程被放入 cache 中.这可以在你需要大量新连接的时候极大的减少线程创建的开销(一般来说如果你有好的线程模型的话, 这不会有明显的性能提升.)服务器线程缓存这个值表示可以重新利用保存在缓存中线程的数量,当断开连接时如果缓存中还有空间,那么客户端的线程将被放到缓存中,如果线程重新被请求，那么请求将从缓存中读取,如果缓存中是空的或者是新的请求，那么这个线程将被重新创建, 如果有很多新的线程，增加这个值可以改善系统性能.通过比较 Connections 和 Threads_created 状态的变量，可以看到这个变量的作用 根据物理内存设置规则如下： 1G —> 8 2G —> 16 3G —> 32 大于3G —> 64
					thread_concurrency
						此允许应用程序给予线程系统一个提示在同一时间给予渴望被运行的线程的数量.该参数取值为服务器逻辑CPU数量×2，在本例中，服务器有 2 颗物理CPU，而每颗物理CPU又支持H.T超线程，所以实际取值为 4 × 2 ＝ 8.设置 thread_concurrency的值的正确与否,
						对 mysql 的性能影响很大, 在多个 cpu(或多核)的情况下，错误设置了 thread_concurrency 的值, 会导致 mysql 不能充分利用多 cpu(或多核),出现同一时刻只能一个 cpu(或核)在工作的情况。 thread_concurrency 应设为 CPU 核数的 2 倍.比如有一个双核的 CPU,
						那么 thread_concurrency 的应该为 4; 2 个双核的 cpu,thread_concurrency 的值应为 8,属重点优化参数
					query_cache_limit
						不缓存查询大于该值的结果.只有小于此设定值的结果才会被缓冲, 此设置用来保护查询缓冲,防止一个极大的结果集将其他所有的查询结果都覆盖.
					query_cache_min_res_unit
						查询缓存分配的最小块大小.默认是 4KB，设置值大对大数据查询有好处，但如果你的查询都是小数据查询，就容易造成内存碎片和浪费
						查询缓存碎片率 = Qcache_free_blocks / Qcache_total_blocks * 100%
						如果查询缓存碎片率超过 20%，可以用 FLUSH QUERY CACHE 整理缓存碎片，或者试试减小query_cache_min_res_unit，如果你的查询都是小数据量的话。
						查询缓存利用率 = (query_cache_size – Qcache_free_memory) / query_cache_size *100%
						查询缓存利用率在 25%以下的话说明 query_cache_size 设置的过大，可适当减小;查询缓存利用率在 80%以上而且 Qcache_lowmem_prunes > 50 的话说明 query_cache_size 可能有点小，要不就是碎片太多。
						查询缓存命中率 = (Qcache_hits – Qcache_inserts) / Qcache_hits * 100%
					query_cache_size
						指定 MySQL 查询缓冲区的大小。可以通过在 MySQL 控制台执行以下命令观察：
							代码:
							> SHOW VARIABLES LIKE '%query_cache%';
							> SHOW STATUS LIKE 'Qcache%';如果 Qcache_lowmem_prunes 的值非常大，则表明经常出现缓冲不够的情况；
							如果 Qcache_hits 的值非常大，则表明查询缓冲使用非常频繁，如果该值较小反而会影响效率，那么可以考虑不用查询缓冲； Qcache_free_blocks，如果该值非常大，则表明缓冲区中碎片很多。
							memlock # 如果你的系统支持 memlock() 函数,你也许希望打开此选项用以让运行中的 mysql 在在内存高度
							紧张的时候,数据在内存中保持锁定并且防止可能被 swapping out,此选项对于性能有益
				binlog
				innodb
					skip-innodb 
						如果你的 MySQL 服务包含 InnoDB 支持但是并不打算使用的话,使用此选项会节省内存以及磁盘空间,并且加速某些部分
					innodb_file_per_table
						# InnoDB为独立表空间模式，每个数据库的每个表都会生成一个数据空间
						独立表空间优点：
						1．每个表都有自已独立的表空间。
						2．每个表的数据和索引都会存在自已的表空间中。
						3．可以实现单表在不同的数据库中移动。
						4．空间可以回收（除drop table操作处，表空不能自已回收）
						缺点：
							1.单表增加过大，如超过100G
						结论：
							共享表空间在Insert操作上少有优势。其它都没独立表空间表现好。当启用独立表空间时，请合理调整：innodb_open_files
					innodb_status_file
						#启用InnoDB的status file，便于管理员查看以及监控等
					innodb_open_files
						限制Innodb能打开的表的数据，如果库里的表特别多的情况，请增加这个。这个值默认是300
					innodb_additional_mem_pool_size
						设置InnoDB存储引擎用来存放数据字典信息以及一些内部数据结构的内存空间大小，所以当我们一个MySQL Instance中的数据库对象非常多的时候，是需要适当调整该参数的大小以确保所有数据都能存放在内存中提高访问效率的。 
					innodb_buffer_pool_size
						包括数据页、索引页、插入缓存、锁信息、自适应哈希所以、数据字典信息.InnoDB 使用一个缓冲池来保存索引和原始数据, 不像 MyISAM.这里你设置越大,你在存取表里面数据时所需要的磁盘 I/O 越少.在一个独立使用的数据库服务器上,你可以设置这个变量到服务器物理内存大小的 80%,不要设置过大,否则,由于物理内存的竞争可能导致操作系统的换页颠簸.注意在 32 位系统上你每个进程可能被限制在 2-3.5G 用户层面内存限制,所以不要设置的太高.
					innodb_write_io_threads/innodb_read_io_threads
						# innodb使用后台线程处理数据页上的读写 I/O(输入输出)请求,根据你的 CPU 核数来更改,默认是4 # 注:这两个参数不支持动态改变,需要把该参数加入到my.cnf里，修改完后重启MySQL服务,允许值的范围从 1-64
					innodb_data_home_dir
						设置此选项如果你希望 InnoDB 表空间文件被保存在其他分区.默认保存在 MySQL 的 datadir 中.
					innodb_data_file_path
						InnoDB将数据保存在一个或者多个数据文件中成为表空间.如果你只有单个逻辑驱动保存你的数据,一个单个的自增文件就足够好了.其他情况下.每个设备一个文件一般都是个好的选择.你也可以配置 InnoDB 来使用裸盘分区 – 请参考手册来获取更多相关内容
					innodb_file_io_threads
						用来同步 IO 操作的 IO 线程的数量. 此值在 Unix 下被硬编码为 4,但是在 Windows 磁盘 I/O 可能在一个大数值下表现的更好.
					innodb_thread_concurrency
						在 InnoDb 核心内的允许线程数量,InnoDB 试着在 InnoDB 内保持操作系统线程的数量少于或等于这个参数给出的限制,最优值依赖于应用程序,硬件以及操作系统的调度方式.过高的值可能导致线程的互斥颠簸.默认设置为 0,表示不限制并发数，这里推荐设置为0，更好去发挥CPU多核处理能力，提高并发量
					innodb_flush_log_at_trx_commit
						#如果设置为 1 ,InnoDB 会在每次提交后刷新(fsync)事务日志到磁盘上,这提供了完整的 ACID 行为.如果你愿意对事务安全折衷, 并且你正在运行一个小的食物, 你可以设置此值到 0 或者 2 来减少由事务日志引起的磁盘 I/O 0 代表日志只大约每秒写入日志文件并且日志文件刷新到磁盘. 2 代表日志写入日志文件在每次提交后,但是日志文件只有大约每秒才会刷新到磁盘上.
					innodb_log_buffer_size
						用来缓冲日志数据的缓冲区的大小.当此值快满时, InnoDB 将必须刷新数据到磁盘上.由于基本上每秒都会刷新一次,所以没有必要将此值设置的太大(甚至对于长事务而言)
					innodb_log_file_size
						事物日志大小.在日志组中每个日志文件的大小，你应该设置日志文件总合大小到你缓冲池大小的5%~100%，来避免在日志文件覆写上不必要的缓冲池刷新行为.不论如何, 请注意一个大的日志文件大小会增加恢复进程所需要的时间.
					innodb_log_files_in_group
						在日志组中的文件总数.通常来说 2~3 是比较好的.
					innodb_log_group_home_dir
						InnoDB 的日志文件所在位置. 默认是 MySQL 的 datadir.你可以将其指定到一个独立的硬盘上或者一个 RAID1 卷上来提高其性能innodb_max_dirty_pages_pct = 90 #innodb 主线程刷新缓存池中的数据，使脏数据比例小于 90%,这是一个软限制,不被保证绝对执行.
					innodb_lock_wait_timeout
						InnoDB 事务在被回滚之前可以等待一个锁定的超时秒数。InnoDB 在它自己的 锁定表中自动检测事务死锁并且回滚事务。 InnoDB 用 LOCK TABLES 语句注意到锁定设置。默认值是 50 秒
					innodb_flush_method
						InnoDB 用来刷新日志的方法.表空间总是使用双重写入刷新方法.默认值是 “fdatasync”, 另一个是 “O_DSYNC”.
					innodb_force_recovery
						如果你发现 InnoDB 表空间损坏, 设置此值为一个非零值可能帮助你导出你的表.从1 开始并且增加此值知道你能够成功的导出表.
					innodb_fast_shutdown
						加速 InnoDB 的关闭. 这会阻止 InnoDB 在关闭时做全清除以及插入缓冲合并.这可能极大增加关机时间, 但是取而代之的是 InnoDB 可能在下次启动时做这些操作.
				myisam
					key_buffer_size
						指定用于索引的缓冲区大小，增加它可得到更好的索引处理性能。如果是以InnoDB引擎为主的DB，专用于MyISAM引擎的 key_buffer_size 可以设置较小，8MB 已足够 如果是以MyISAM引擎为主，可设置较大，但不能超过4G. 在这里，强烈建议不使用MyISAM引擎，默认都是用InnoDB引擎.注意：该参数值设置的过大反而会是服务器整体效率降低！
					sort_buffer_size
						查询排序时所能使用的缓冲区大小。排序缓冲被用来处理类似 ORDER BY 以及 GROUP BY 队列所引起的排序.一个用来替代的基于磁盘的合并分类会被使用.查看 “Sort_merge_passes” 状态变量. 在排序发生时由每个线程分配 注意：该参数对应的分配内存是每连接独占！如果有 100 个连接，那么实际分配的总共排序缓冲区大小为 100 × 6 ＝600MB,所以,对于内存在 4GB 左右的服务器推荐设置为 6-8M。 
					join_buffer_size
						读查询操作所能使用的缓冲区大小。和 sort_buffer_size 一样，该参数对应的分配内存也是每连接独享！用来做 MyISAM 表全表扫描的缓冲大小.当全表扫描需要时,在对应线程中分配.
					read_buffer_size
						联合查询操作所能使用的缓冲区大小，和 sort_buffer_size 一样，该参数对应的分配内存也是每连接独享!此缓冲被使用来优化全联合(full JOINs 不带索引的联合).类似的联合在极大多数情况下有非常糟糕的性能表现, 但是将此值设大能够减轻性能影响.通过 “Select_full_join”状态变量查看全联合的数量， 当全联合发生时,在每个线程中分配。
					read_rnd_buffer_size
						MyISAM 以索引扫描(Random Scan)方式扫描数据的 buffer大小
					bulk_insert_buffer_size
						MyISAM 使用特殊的类似树的 cache 来使得突发插入(这些插入是,INSERT … SELECT, INSERT … VALUES (…), (…), …, 以及 LOAD DATAINFILE) 更快. 此变量限制每个进程中缓冲树的字节数.设置为 0 会关闭此优化.为了最优化不要将此值设置大于 “key_buffer_size”.当突发插入被检测到时此缓冲将被分配MyISAM 用在块插入优化中的树缓冲区的大小。注释：这是一个 per thread 的限制 （ bulk 大量）.此缓冲当 MySQL 需要在 REPAIR, OPTIMIZE, ALTER 以及 LOAD DATA INFILE到一个空表中引起重建索引时被分配.这在每个线程中被分配.所以在设置大值时需要小心.
					myisam_sort_buffer_size
						MyISAM 设置恢复表之时使用的缓冲区的尺寸,当在REPAIR TABLE 或用 CREATE INDEX 创建索引或 ALTER TABLE 过程中排序 MyISAM 索引分配的缓冲区
					myisam_max_sort_file_size
						mysql重建索引时允许使用的临时文件最大大小
					myisam_repair_threads
						如果该值大于 1，在 Repair by sorting 过程中并行创建MyISAM 表索引(每个索引在自己的线程内).如果一个表拥有超过一个索引, MyISAM 可以通过并行排序使用超过一个线程去修复他们.这对于拥有多个 CPU 以及大量内存情况的用户,是一个很好的选择.
					myisam_recover
						允许的 GROUP_CONCAT()函数结果的最大长度 transaction_isolation = REPEATABLE-READ # 设定默认的事务隔离级别.可用的级别如下:READ-UNCOMMITTED, READ-COMMITTED, REPEATABLE-READ,SERIALIZABLE 1.READ UNCOMMITTED-读未提交 2.READ COMMITTE-读已提交 3.REPEATABLE READ -可重复读 4.SERIALIZABLE -串行
			数据配置项
				default_table_type
					当创建新表时作为默认使用的表类型,如果在创建表示没有特别执行表类型,将会使用此值
				character-set-server
					server 级别字符集
				default-storage-engine
					默认存储引擎
				tmp_table_size
					临时表的最大大小，如果超过该值，则结果放到磁盘中,此限制是针对单个表的,而不是总和.
			日志项
				binlog
					binlog_cache_size
						在一个事务中 binlog 为了记录 SQL 状态所持有的 cache 大小,如果你经常使用大的,多声明的事务,你可以增加此值来获取更大的性能.所有从事务来的状态都将被缓冲在 binlog 缓冲中然后在提交后一次性写入到 binlog 中,如果事务比此值大, 会使用磁盘上的临时文件来替代.此缓冲在每个连接的事务第一次更新状态时被创建.session 级别
					log-bin
						打开二进制日志功能.在复制(replication)配置中,作为 MASTER 主服务器必须打开此项.如果你需要从你最后的备份中做基于时间点的恢复,你也同样需要二进制日志.这些路径相对于 datadir
					log-bin-index
						二进制的索引文件名
					expire_logs_days
						超过 30 天的 binlog 删除
					max_binlog_size
					如果二进制日志写入的内容超出给定值，日志就会发生滚动。你不能将该变量设置为大于1GB或小于4096字节。 默认值是1GB。如果你正使用大的事务，二进制日志还会超过max_binlog_size
				relaylog
					relay-log
						定义relay_log的位置和名称，如果值为空，则默认位置在数据文件的目录，文件名为host_name-relay-bin.nnnnnn（By default, relay log file names have the form host_name-relay-bin.nnnnnn in the data directory）；
					relay_log_index
						relay-log的索引文件名
					max_relay_log_size
						标记relaylog允许的最大值，如果该值为0，则默认值为max_binlog_size(1G)；如果不为0，则max_relay_log_size则为最大的relay_log文件大小；
					relay-log-purge
						是否自动清空不再需要中继日志时。默认值为1(启用)
			其他
				log-warnings
					将警告打印输出到错误 log 文件.如果你对于MySQL有任何问题，你应该打开警告 log 并且仔细审查错误日志,查出可能的原因.
				log-error
					错误日志路径
				log_output
					参数 log_output 指定了慢查询输出的格式，默认为 FILE，你可以将它设为 TABLE，然后就可以查询 mysql 架构下的 slow_log 表了
				log_slow_queries
					指定是否开启慢查询日志(该参数要被slow_query_log取代，做兼容性保留)
				slow_query_log
					指定是否开启慢查询日志. 慢查询是指消耗了比 “long_query_time” 定义的更多时间的查询.如果 log_long_format 被打开,那些没有使用索引的查询也会被记录.如果你经常增加新查询到已有的系统内的话. 一般来说这是一个好主意,
				long-query-time
					设定慢查询的阀值，超出次设定值的SQL即被记录到慢查询日志，缺省值为10s.所有的使用了比这个时间(以秒为单位)更多的查询会被认为是慢速查询.不要在这里使用”1″, 否则会导致所有的查询,甚至非常快的查询页被记录下来(由于MySQL 目前时间的精确度只能达到秒的级别).
				log_long_format
					在慢速日志中记录更多的信息.一般此项最好打开，打开此项会记录使得那些没有使用索引的查询也被作为到慢速查询附加到慢速日志里
				slow_query_log_file
					指定慢日志文件存放位置，可以为空，系统会给一个缺省的文件host_name-slow.log
				log-queries-not-using-indexes
					如果运行的SQL语句没有使用索引，则mysql数据库同样会将这条SQL语句记录到慢查询日志文件中。
				min_examined_row_limit
					记录那些由于查找了多余n次而引发的慢查询
				long-slow-admin-statements　　　　
					记录那些慢的optimize table，analyze table和alter table语句
				log-slow-slave-statements
					记录由Slave所产生的慢查询
				general_log
					将所有到达MySQL Server的SQL语句记录下来,默认关闭
				general_log_file
					general_log路径
		
		集群
			slave-load-tmpdir
				当 slave 执行 load data infile 时用
			skip-slave-start 
			slave-net-timeout
			net_read_timeout
			net_write_timeout
			log_slave_updates
				表示slave将复制事件写进自己的二进制日志
			replicate-wild-ignore-table
			slave_skip_errors
			杂项
				sysdate-is-now
			mysqld_safe
				open-files-limit
					#增加每个进程的可打开文件数量.确认你已经将全系统限制设定的足够高!打开大量表需要将此值设大
```				