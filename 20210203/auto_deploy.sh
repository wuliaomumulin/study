#!/bin/bash
# liyb
#
# SRC=/etc/
flush(){
  if [ ! -f rsync.list ];then
    echo -e "\033[31mplease create rsync.list Files,the rsync.list contents as follow: \033[0m"
cat <<EOF
19.19.19.70 src_dir des_dir
192.168.1.86 src_fir des_dir
EOF
	  exit
  fi

  rm -rf rsync.list.swp;cat rsync.list |grep -v "#" > rsync.list.swp
  COUNT=`cat rsync.list.swp|wc -l`
  NUM=0

  while ((${NUM} < $COUNT))
  do
    NUM=`expr $NUM + 1`
    Line=`sed -n "${NUM}p" rsync.list.swp`
    SRC=`echo $Line |awk '{print $2}'`
    DES=`echo $Line |awk '{print $3}'`
    IP=`echo $Line |awk '{print $1}'`
    # 一致性强制同步
    rsync -aP --delete ${SRC}/ root@${IP}:${DES}/
  done
}
# 下面这个函数不好使
restart(){
  if [ ! -f restart.list ];then
    echo -e "\033[31mplease create restart.list Files,the restart.list contents as follow: \033[0m"
	  exit
  fi
  rm -rf restart.list.swp;cat restart.list |grep -v "#" > restart.list.swp
  COUNT=`cat restart.list.swp|wc -l`
  NUM=0

  while ((${NUM} < $COUNT))
  do
    NUM=`expr $NUM + 1`
    Line=`sed -n "${NUM}p" restart.list.swp`
    Command=`echo $Line |awk '{print $2}'`
    IP=`echo $Line |awk '{print $1}'`
    ssh -l root $IP "sh $Command;echo -e '-----------------------------\nThe $IP exec Command: sh $Command success !'"
  done
}

case $1 in
  flush )
  flush
  ;;
  restart )
  restart
  ;;
  * )
  echo -e "\033[31mUage:$0 command,example{flush | restart} \033[0m"
  ;;
esac
