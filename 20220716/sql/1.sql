CREATE TABLE `calendar` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `calendar_date` datetime DEFAULT NULL COMMENT '日期',
  `calendar_year` year(4) DEFAULT extract(year from `calendar_date`) COMMENT '年',
  `calendar_month` tinyint(1) DEFAULT extract(month from `calendar_date`) COMMENT '月',
  `calenday_day` tinyint(1) DEFAULT extract(day from `calendar_date`) COMMENT '日',
  `is_work_day` enum('Y','N') NOT NULL DEFAULT (case when date_format(`calendar_date`,'%w') between 1 and 5 then 'Y' else 'N' end) COMMENT '是否工作日',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 数据透视表
select coalesce(product,'--') as '产品',coalescs(channel,'-') as '渠道',
sum(case extract(month from saledate) when 1 then amount else 0 END) AS '一月',
sum(case extract(month from saledate) when 2 then amount else 0 END) AS '二月',
sum(case extract(month from saledate) when 3 then amount else 0 END) AS '三月',
sum(case extract(month from saledate) when 4 then amount else 0 END) AS '四月',
sum(case extract(month from saledate) when 5 then amount else 0 END) AS '五月',
sum(case extract(month from saledate) when 6 then amount else 0 END) AS '六月',
sum(amount) as '合计'
from sales_data sd group by product,channel WITH ROLLUP;
-- 行列转换

-- 行列合并
WITH t (a, b, c, d) AS (
	SELECT
		a.*, b.b
	FROM
		ta a
	JOIN ta b ON FIND_IN_SET(b.a, a.c)
	ORDER BY
		a.a
) SELECT
	a,
	b,
	GROUP_CONCAT(d)
FROM
	t
GROUP BY
	a,
	b