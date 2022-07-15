## mysql和mongo的查询比较

### fragments 1
```sql
select first_name as `名`,last_name as `姓` from users where gender = 1 skip 100 limit 20;
```

```shell
db.users.aggregate([
	{$match:{gender:1}},
	{$skip:100},
	{$limit:20},
	{$project:{
		'名':'$first_name',
		'姓':'$last_name'
	}}
]);
```

### fragments 2
```sql
select department,count(NULL) AS emp_qty from users where gender=2 group by department having count(1)<10;
```

```shell
db.users.aggregate([
	{$match:{gender:2}},
	{$group:{
		"_id":$department,
		"emp_qty":{$sum:1}
	}},
	{$match:{
		'emp_qty':{$lt:10}}}
]);
```
