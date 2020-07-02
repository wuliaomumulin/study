package main
import(
	"fmt"
	"strings"
)

/*
* string
* 字符串需要用双引号包裹，单引号的叫字符;
* 字符是单独的一个字符，可能是一个汉字或一个字母或一个符号
* UTF-8 一个字符暂居多个字节 1个字符=1个字节，一个utf8编码的汉字"沙"==一般3个字节
* UNCODE  一个utf8编码的汉字"沙"==一般2个字节
* byte and rune类型,Go语言中为了处理非ASCIII码类型的字符，定义了新的类型rune
*/
func main(){
	s:="hello,"
	a:='A'//asiii值，场景：用户输入计算
	fmt.Printf("%s%d\n",s,a)
	url := "'E:\\www\\go'"
	fmt.Println("\"E:\\www\\go\"")
	fmt.Println(url)
	/* 定义多行字符串` */
	fmt.Println(`"F:\www\go我"`)
	b :=`
世情薄，
人情恶，
	雨送黄昏花易落。
	`
	fmt.Println(b)
	//字符串相关操作
	fmt.Println(len(b))//字符串不对
	fmt.Println(s+b)//字符串拼接
	fmt.Println(len(fmt.Sprintf("%s%v,你好",s,a)))//字符串拼接+
	fmt.Println(fmt.Sprintf("%s%v,你好",s,a))//字符串拼接Sprintf,返回拼接之后的字符串
	fmt.Println(strings.Split(url,"\\"))//回一个列表，类似数组,类似PHP的explode
	fmt.Println(strings.Contains(url,"go啊"))//检测是否包含
	fmt.Println(strings.HasPrefix(url,"go"))//前缀判断
	fmt.Println(strings.HasSuffix(url,"go"))//后缀判断
	fmt.Println(strings.Index(url,"w"))//判断字符串第一次出现在主字符串的位置，返回下标
	fmt.Println(strings.LastIndex(url,"w"))//判断字符串最后一次出现在主字符串的位置，返回下标
	fmt.Println(strings.Join(strings.Split(url,"\\"),"+"))//拼接字符串，类似PHP的implode
	fmt.Println(`------------字符串修改----------------`)
	s1 := "白萝卜"
	s2 := []rune(s1)//把字符串强制转换成一个rune切片
	s2[0] = '红'//切片修改
	s1 = string(s2)
	fmt.Printf("s2=%T ,%T,%v\n",s2,s1,s1)//int32的别名就是rune
	fmt.Printf("%T",byte('H'))//uint8
	fmt.Println(`------------类型转换----------------`)
	fmt.Printf("%T",float64(10));//类型转换，整形和浮点型转换
}