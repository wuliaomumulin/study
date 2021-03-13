package main
import(
	"fmt"
	"runtime"
	"database/sql"
	_"github.com/go-sql-driver/mysql"//引用该包，只期望他执行init()函数，所以无法通过包名调用其他函数
)
/*



简答操作mysql
	下载依赖:
		go get -u github.com/go-sql-driver/mysql
	原型:
		func Open(dirverName,dataSourceName string)(*DB,error)
	更改window-cmd字符集为utf8:
		chcp 65001
*/
func main(){
	test1()
	fmt.Println(test2())
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
	test9()
	test10()
	defer db.Close() //关闭句柄要卸载err判断下面
}
//返回调用者函数名称
func getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
/*
	
*/
var db *sql.DB//初始化一个全局db指针


func test1(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	//DSN:Data Source Name
	dsn:="root:123456@tcp(19.19.19.70:3306)/test"
	db,err:=sql.Open("mysql",dsn)
	if err!=nil {
		panic(err)
	}
	defer db.Close() //关闭句柄要卸载err判断下面

}
/**
	可以安全的被多个goroutine同时使用
*/
func test2()(err error){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	//DSN:Data Source Name
	dsn:="root:123456@tcp(19.19.19.70:3306)/test?charset=utf8mb4&parseTime=True"
	//这个阶段不会验证账号密码错误或成功
	db,err=sql.Open("mysql",dsn)
	if err!=nil {
		return err
	}
	//尝试与数据库建立连接，验证dsn是否正确
	err=db.Ping()
	if err!=nil {
		return err
	}

	return nil
}
/*
func (db *DB) SetMaxOpenConns(n int)	
SetMaxOpenConns设置与数据库连接的最大数目，如果n大于0小于最大闲置连接数目，会将最大闲置连接数减少到匹配最大开启连接数的限制，如果n<=0,不会限制最大开启连接数，默认为O(无限制)。
func (db *DB) SetMaxIdleConns(n int)
SetMaxIdleConns设置连接池中的最大闲置连接数，如果n大于最大开启连接数，则新的最大闲置连接数会减少到匹配最大开启连接数的限制。如果n<=0,不会保留闲置连接

curd操作

*/
//查询数据
type user struct{
	id int
	username string
	nickname string
}
//查询单条数据
func queryRow(){
	//单条数据
	sqlstr:="select id,username,nickname from yaf_user where id=?"//注意条件
	var u user;

	//非常重要:确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err:=db.QueryRow(sqlstr,1).Scan(&u.id,&u.username,&u.nickname)
	if err!=nil {
		fmt.Println("scan failed,error:",err)
		return
	}
	fmt.Println(u);
}
//多行查询
func queryMultiRow(){
	sqlstr:="select id,username,nickname from yaf_user where id>?"//注意条件
	rows,err:=db.Query(sqlstr,0)
	if err!=nil{
		fmt.Println("query failed,error:",err)
		return
	}
	//非常重要:关闭rows释放持有的数据库链接
	defer rows.Close()

	//循环读取结果集中的数据
	for rows.Next(){
		var u user
		err=rows.Scan(&u.id,&u.username,&u.nickname)
		if err!=nil{
			fmt.Println("scan failed,error:",err)
			return
		}
		fmt.Println(u)
	}
}
//查询
func test3(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	//查询单条数据
	queryRow()
	fmt.Println("--查询多行数据--")
	queryMultiRow()
}

//插入
func test4(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	
	sqlstr:="insert into yaf_user(username,nickname) value(?,?)"
	ret,err:=db.Exec(sqlstr,"wuwang","王五")
	if err!=nil{
		fmt.Println("insert failed,error:",err)
		return
	}
	theID,err:=ret.LastInsertId()//新插入数据的id
	if err!=nil{
		fmt.Println("get LastInsertId failed,error:",err)
		return
	}
	fmt.Println("insert seccess,the id is ",theID)
}
/*
更新
*/
func test5(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	sqlstr:="update yaf_user set mobile=?,email=? where username=?"
	ret,err:=db.Exec(sqlstr,"15210193108","39","wuwang")
	if err!=nil{
		fmt.Println("update failed,error:",err)
		return
	}
	n,err:=ret.RowsAffected()//影响行数
	if err!=nil{
		fmt.Println("get RowsAffecteds failed,error:",err)
		return
	}
	fmt.Println("update seccess,affected rows is ",n)

}
/*
删除
*/
func test6(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	sqlstr:="delete from yaf_user where username=?"
	ret,err:=db.Exec(sqlstr,"wuwang")
	if err!=nil{
		fmt.Println("delete failed,error:",err)
		return
	}
	n,err:=ret.RowsAffected()//影响行数
	if err!=nil{
		fmt.Println("get RowsAffecteds failed,error:",err)
		return
	}
	fmt.Println("delete seccess,affected rows is ",n)

}
//查询预处理
func test7(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	sqlstr:="select id,username,nickname from yaf_user where id>?"//注意条件
	stmt,err:=db.Prepare(sqlstr)
	if err!=nil{
		fmt.Println("Prepare failed,error:",err)
		return
	}
	defer stmt.Close()

	rows,err:=stmt.Query(0)
	if err!=nil{
		fmt.Println("query failed,error:",err)
		return
	}
	//非常重要:关闭rows释放持有的数据库链接
	defer rows.Close()

	//循环读取结果集中的数据
	for rows.Next(){
		var u user
		err=rows.Scan(&u.id,&u.username,&u.nickname)
		if err!=nil{
			fmt.Println("scan failed,error:",err)
			return
		}
		fmt.Println(u)
	}

}
/**
插入预处理
*/
func test8(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	sqlstr:="insert into yaf_user(username,nickname) value(?,?)"
	stmt,err:=db.Prepare(sqlstr)
	if err!=nil{
		fmt.Println("Prepare failed,error:",err)
		return
	}
	defer stmt.Close()

	_,err=stmt.Exec("wuwang","王五")
	if err!=nil{
		fmt.Println("insert failed,error:",err)
		return
	}
	_,err=stmt.Exec("wuwang","张三")
	if err!=nil{
		fmt.Println("insert failed,error:",err)
		return
	}
	fmt.Println("insert seccess.affected rows is 2")
}
/**
	sql注入问题
*/
func sqlInject(name string){
	sqlstr:=fmt.Sprintf("select id,username,nickname from yaf_user where username='%s'",name)
	fmt.Println("SQL:",sqlstr)
	var u user
	err:=db.QueryRow(sqlstr).Scan(&u.id,&u.username,&u.nickname)
	if err!=nil{
		fmt.Println("QueryRow failed,error:",err)
		return
	}
	fmt.Println("user:",u)
}
/*
此时输入以下字符串都可以引发SQL注入问题:
	sqlInject("xxx' or 1=1")
	sqlInject("xxx' union select * from yaf_user #")
	sqlInject("xxx' and (select count(*) from yaf_user) <10 #")
*/
func test9(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	sqlInject("wuwang")
}
/*
事务操作
*/
func test10(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	tx,err:=db.Begin()
	if err!=nil{
		if tx!=nil{
			tx.Rollback()
		}
		fmt.Println("begin trans failed,error:",err)
		return
	}
	
	sqlstr:="update yaf_user set mobile=?,email=? where username=?"
	ret,err:=db.Exec(sqlstr,"15210193108","39","wuwang")
	if err!=nil{
		tx.Rollback()//回滚
		fmt.Println("first exec failed,error:",err)
		return
	}
	n,err:=ret.RowsAffected()//影响行数

	sqlstr="delete from yaf_user where username=?"
	ret,err=db.Exec(sqlstr,"wuwang")
	if err!=nil{
		tx.Rollback()//回滚
		fmt.Println("second exec failed,error:",err)
		return
	}
	n2,err:=ret.RowsAffected()//影响行数

	fmt.Println(n,n2)

	if(n>0 && n2>0){
		fmt.Println("事务提交了...")
		tx.Commit()
	}else{
		tx.Rollback()
		fmt.Println("事务回滚了...")
	}

	fmt.Println("exec trans success")

}