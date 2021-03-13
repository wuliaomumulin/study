package main
import(
	"fmt"
	"runtime"
	"sync"
	"errors"

	"time"//本功能新版本需要的依赖
	"context"//新版本所需要的依赖
	"github.com/go-redis/redis/v8"//注意导入的是新版本
	//"github.com/go-redis/redis"//注意导入的是旧版本
)
/*



简单操作redis

1、支持字符串、hash、lists、sets、带范围查询的排序集合(sorted sets)、位图(bitmaps)、hyperloglogs、带半径查询和流的地理空间索引等数据结构
2、通过复制、持久化、客户端分片等特性，我们可以很方便的将Redis扩展成一个能够包含数百GB数据，每秒处理上百万次请求的系统
3、应用场景
	a、缓存系统，减轻主数据压力。
	b、计数场景，比如微博、抖音中的关注数和粉丝数
	c、热门排行榜,需要排序的场景特别适合使用ZSET
	d、利用list可以实现队列的功能
下载依赖:
	go get -u github.com/go-redis/redis
更改window-cmd字符集为utf8:
	chcp 65001
查看gopath
	go env GOPATH
报错:no required module provides package github.com/go-redis/redis/v8: working directory is not part of a module

解决记录:
1、go mod init lin.com/m
2、go get -u github.com/go-redis/redis/v8


查看帮助或者函数原型
https://pkg.go.dev/github.com/go-redis/redis

*/
func main(){
	//普通连接,v8之前的版本
	//fmt.Println(test1())
	//v8版本
	test2()

	//fmt.Println("哨兵模式:",initFailoverClient())
	//fmt.Println("集群模式:",initClusterClient())
	//test3()

	test4()
	test5()
	test6()
	test7()
	test8()
	test9()
	test10()
}
//返回调用者函数名称
func getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
/*
	
*/


/*
普通连接,v8之前的版本
否则报错:
not enough arguments in call to rdb.cmdable.Ping
have ()
want (context.Context)

*/
var rdb *redis.Client//初始化一个全局redis指针
//v8之前的版本
/*func test1()(err error){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	rdb=redis.NewClient(&redis.Options{
		Addr:"19.19.19.70:6379",
		Password:"",
		DB:0,
	})
	_,err=rdb.Ping().Result()
	//返回nil即为成功
	return err
}*/
//普通连接,v8版本
func initClient()(err error){
	rdb=redis.NewClient(&redis.Options{
		Addr:"19.19.19.70:6379",
		Password:"",
		DB:0,
		PoolSize:100,//连接池大小
	})
	//Background()返回一个空的context
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	_,err=rdb.Ping(ctx).Result()
	//返回nil即为成功
	return err
}
//v8 example
func test2(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	
	ctx:=context.Background()
	if err:=initClient();err!=nil{return }

	err:=rdb.Set(ctx,"rand-key","rand-value",0).Err()
	if err!=nil {panic(err)}	

	//已经存在的值
	val,err:=rdb.Get(ctx,"rand-key").Result()
	if err!=nil {panic(err)}
	fmt.Println("rand-key is",val)

	//不存在的值
	val2,err:=rdb.Get(ctx,"rand-key2").Result()
	if err==redis.Nil{
		fmt.Println("rand-key2 doesn't exist")
	}else if(err!=nil){
		panic(err)
	}else{
		fmt.Println("rand-key2 is",val2)
	}
}


/**
	连接redis哨兵模式
*/
/*func initFailoverClient()(err error){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	rdb:=redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:"master",
		SentinelAddrs:[]string{"19.19.19.70:6379","192.168.1.86:6379"},
	})
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	_,err=rdb.Ping(ctx).Result()
	return err
}
//连接redis集群
func initClusterClient()(err error){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	rdb:=redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:[]string{":6379",":7000",":7001",":7002",":7003"},
	})
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	_,err=rdb.Ping(ctx).Result()
	return err
}*/
//zset有序集合示例 ,旧版本的
/*func test3(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	zsetkey:="language_rank"
	languages:=[]redis.Z{
		redis.Z{Score:77,Member:"C++"},
		redis.Z{Score:60,Member:"PHP"},
		redis.Z{Score:74,Member:"Goland"},
		redis.Z{Score:63,Member:"Python"},
	}

	//Zadd
	num,err:=rdb.ZAdd(zsetkey,languages...).Result()
	if err!=nil{
		fmt.Println("zadd failed,err:",err)
		return
	}
	fmt.Println("zadd success,affected rows is ",num)

	//把php的分数追加5
	newScore,err:=rdb.ZIncrBy(zsetkey,5,"PHP").Result()
	if err!=nil{
		fmt.Println("ZIncrBy failed,err:",err)
		return
	}
	fmt.Println("PHP Score is ",newScore)

	//取分数最高的三个
	fmt.Println("取分数最高的三个")
	ret,err:=rdb.ZRevRangeWithScores(zsetkey,0,2).Result()
	if err!=nil{
		fmt.Println("ZRevRangeWithScores failed,err:",err)
		return
	}
	for _,z:=range ret{
		fmt.Println(z.Member,"=",z.Score)
	}

	//取60~70分的
	fmt.Println("取60~70分的")
	op:=redis.ZRangeBy{
		Min:"60",
		Max:"70",
	}
	ret,err=rdb.ZRangeByScoreWithScores(zsetkey,op).Result()
	if err!=nil{
		fmt.Println("ZRevRangeWithScores failed,err:",err)
		return
	}
	for _,z:=range ret{
		fmt.Println(z.Member,"=",z.Score)
	}

}*/

/**
vals,err:=rdb.keys(ctx,"prefix*").Result()//根据前缀获取key
ret,err:=rdb.Do(ctx,"set","key","val").Result()//执行自定义命令

按通配符删除key
*/

func test4(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	ctx:=context.TODO()
	iter:=rdb.Scan(ctx,0,"*_rank",0).Iterator()
	for iter.Next(ctx) {
		err:=rdb.Del(ctx,iter.Val()).Err()
		if err!=nil{panic(err)}
	}
	if err:=iter.Err();err!=nil {panic(err)}
}
/*
	pipeline主要是一种网络优化，它本质上意味着客户端缓冲一些命令并一次性将他们发送给服务器，这些命令不能保证在事务中执行，这样做的好处是节省了每个命令的网络往返时间(RTT)
*/
func test5(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	ctx:=context.TODO()
	pipe:=rdb.Pipeline()
	incr:=pipe.Incr(ctx,"pipeline_counter")
	pipe.Expire(ctx,"pipeline_counter",time.Hour)
	_,err:=pipe.Exec(ctx)
	fmt.Println(incr.Val(),err)
	//相当于两个命令依次发给redis server端执行，与不使用相比减少一次RTT
	
}
//也可以使用pipelined
func test6(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var incr *redis.IntCmd
	ctx:=context.TODO()
	_,err:=rdb.Pipelined(ctx,func(pipe redis.Pipeliner) error {
		incr=pipe.Incr(ctx,"pipeline_counter")
		pipe.Expire(ctx,"pipeline_counter",time.Minute)
		return nil
	})
	fmt.Println(incr.Val(),err)
	//在某些场景下，当我们有多条命令要执行时，就可以考虑使用pipeline来优化
}
/*
事务
Multi/exec能够确保在multi/exec两个语句之间的命令之间没有其他客户端正在执行命令。
在这种场景下我们需要使用TxPipeline.TxPipeline总体上类似于上面的Pipeline,但是它内部会使用MULTI/EXEC包裹排队的命令，例如
*/
func test7(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	pipe:=rdb.TxPipeline()
	ctx:=context.TODO()
	incr:=pipe.Incr(ctx,"pipeline_counter")
	pipe.Expire(ctx,"pipeline_counter",time.Minute)
	_,err:=pipe.Exec(ctx)
	fmt.Println(incr.Val(),err)
	/*
		上面代码相当于:

		MULTI
		incr pipeline_counter
		expire pipeline_counter 60
		EXEC
	*/
}
//还有一个与上下文类似的TxPipelined方法，使用方法如下:
func test8(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var incr *redis.IntCmd
	ctx:=context.TODO()
	_,err:=rdb.TxPipelined(ctx,func(pipe redis.Pipeliner) error {
		incr=pipe.Incr(ctx,"pipeline_counter")
		pipe.Expire(ctx,"pipeline_counter",time.Minute)
		return nil
	})
	fmt.Println(incr.Val(),err)
}
/*
Watch
在某些场景下，我们除了要使用MULTI/EXEC命令外，还需要配合使用watch命令，在用户使用watch命令监视某个键之后，直到该用户执行exec命令之短时间里，如果有其他用户抢先对被监视的键进行了替换、更新、删除等操作，那么当用户尝试执行EXEC的时候，事务失败并将返回一个错误，用户可以根据这个错误选择重试事务或者放弃事务.
*/
func test9(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	//监视pipeline_counter的值，并在值不变的情况下将其值+1
	key:="pipeline_counter"
	ctx:=context.TODO()

	err:=rdb.Watch(ctx,func(tx *redis.Tx) error{
		n,err:=tx.Get(ctx,key).Int()
		if err!=nil && err!=redis.Nil{
			return err
		}
		_,err=tx.Pipelined(ctx,func(pipe redis.Pipeliner) error{
			pipe.Set(ctx,key,n+1,0)
			return nil
		})
		return err
	//Watch方法额可以接收一个函数和一个或多个key作为参数
	},key)

	fmt.Println(err)
}
/*
	v8 example
*/
func test10(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var (
		MaxRetries=1000
		routineCount=10
	)
	ctx,cannel:=context.WithTimeout(context.TODO(),5*time.Second)
	defer cannel()

	//Increment使用Get和Set命令以事务方式递增key的值
	//将匿名函数赋值给increment，意思就是func increment(key string) error{}
	increment:=func(key string) error{
		//事务函数
		txf:=func(tx *redis.Tx) error{
			//获得key的值或者当前值
			n,err:=tx.Get(ctx,key).Int()
			if err!=nil && err!=redis.Nil{
				return err
			}
			//实际的操作代码。乐观锁定中的本地操作
			n++

			//操作仅在Watch的key没发生变化的情况下提交
			_,err=tx.TxPipelined(ctx,func(pipe redis.Pipeliner) error{
				pipe.Set(ctx,key,n,time.Minute)
				return nil
			})
			return err
		}

		//最多重试MaxRetries次
		for i:=0;i<MaxRetries;i++ {
			err:=rdb.Watch(ctx,txf,key)
			if err==nil{
				//成功
				return nil
			}
			if err==redis.TxFailedErr{
				//乐观锁丢失，重试
				continue
			}
			//返回其他错误信息
			return err
		}
		return errors.New("increment reached maximum number of retries")
	}

	//模拟 routineCount个并发同时去修改counter3的值
	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i:=0;i<routineCount;i++{
		go func(){
			defer wg.Done()
			if err:=increment("pipeline_counter");err!=nil{
				fmt.Println("increment error:",err)
			}
		}()
	}
	wg.Wait()

	n,err:=rdb.Get(context.TODO(),"pipeline_counter").Int()
	fmt.Println("ended with",n,err)
}