#!/bin/bash

# sed

# 预修改 sed 's/192.10/192.168/g' test.txt
# 修改 sed -i 's/192.10/192.168/g' test.txt

# 在行首添加一个空格、字符
sed 's/^/& /g' test.txt
sed 's/^/&id/g' test.txt

# 在行末添加一个字符
sed 's/$/&id/g' test.txt

# 特定字符之后添加一行
sed '/wugangke/a ######' test.txt

# 特定字符之前添加一行
sed '/wugangke/i ######' test.txt

# 检索某一行的数据并打印到终端
sed -n '/wugangke/p' test.txt

# 打印第一行
sed -n '1p' test.txt

# 打印第一行到第三行
sed -n '1,3p' test.txt

# 现将空格替换成换行，之后排除掉以空为开头并且以空为结尾的行,然后再从大到小排序，最后取出第一个和最后一个值
more number.txt | sed 's/ /\n/g'|grep -v "^$" | sort -nr | sed -n '1p;$p'



## password admin123Aa
echo "bbb"
echo "aaa"
