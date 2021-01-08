本文主要向大家介绍了MySQL数据库之MySQL的binary类型操作 ，通过具体的内容向大家展现，希望对大家学习MySQL数据库有所帮助。
 
示例数据表： 
CREATE TABLE test_bin (
   bin_id BINARY(16) NOT NULL
) Engine=InnoDB; 
 
插入数据(内容是一个32位的UUID字符串值)： 
INSERT INTO test_bin(bin_id) VALUES(UNHEX(‘FA34E10293CB42848573A4E39937F479‘));
INSERT INTO test_bin(bin_id) VALUES(UNHEX(?));
或
INSERT INTO test_bin(bin_id) VALUES(x‘FA34E10293CB42848573A4E39937F479‘); 
 
查询数据： 
SELECT HEX(bin_id) AS bin_id FROM test_bin;
 
SELECT HEX(bin_id) AS bin_id FROM test_bin WHERE bin_id = UNHEX(‘FA34E10293CB42848573A4E39937F479‘);
SELECT HEX(bin_id) AS bin_id FROM test_bin WHERE bin_id = UNHEX(?);
 
SELECT HEX(bin_id) AS bin_id FROM test_bin WHERE bin_id = x‘FA34E10293CB42848573A4E39937F479‘;
 
查询结果：
bin_id
--------------------------
FA34E10293CB42848573A4E39937F479 
 
 
备注：使用MySQL内置的 UUID() 创建一个函数返回 BINARY（16）类型的UUID值 
CREATE FUNCTION uu_id() RETURNS binary(16) RETURN UNHEX(REPLACE(UUID(),‘-‘,‘‘));
或
CREATE FUNCTION uu_id() RETURNS binary(16) RETURN UNHEX(REVERSE(REPLACE(UUID(),‘-‘,‘‘)));
 
使用：
INSERT INTO test_bin(bin_id) VALUES(uu_id());