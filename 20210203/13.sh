#!/bin/bash

# Automatic Backup Linux System Files
# 自动备份文件系统
# Author liyb
# date -s 20210131
# date -s "20210201 23:06:45"

# Define Variable
SOURCE_DIR=(
$*
)

TARGET_DIR=/data/backup/
YEAR=`date +%Y`
MONTH=`date +%m`
DAY=`date +%d`
WEEK=`date +%u`
A_NAME=`date +%H%M`
FILES=system_backup.tgz
CODE=$?

if
  [ -z "$*" ];then
    echo -e "\033[32mUsage:\nPlease   Enter	Your Backup Files	or Directories\n	\n\nUsage: { $0 /boot /etc}\033[0m"
    exit
fi

#Determine Whether the Target Directory Exists
if
  [ ! -d $TARGET_DIR/$YEAR/$MONTH/$DAY ];then
    mkdir -p $TARGET_DIR/$YEAR/$MONTH/$DAY
    echo -e "\033[32mThe $TARGET_DIR Created Successfully !\033[0m"
fi
#EXEC Full_Backup Function Command
# 全量备份
Full_Backup()
{
  if
  [ "$WEEK" -eq "7" ];then

    rm -rf $TARGET_DIR/snapshot
    cd $TARGET_DIR/$YEAR/$MONTH/$DAY ;tar -g $TARGET_DIR/snapshot -czvf $FILES ${SOURCE_DIR[@]}
    [ "$CODE" == "0" ]&&echo -e "	\n\033[32mThese Full_Backup System Files Backup Successfully !\033[0m"
  fi
}

#增量备份
Add_Backup()
{
  if
  [ $WEEK -ne "7" ];then
    cd $TARGET_DIR/$YEAR/$MONTH/$DAY ;tar -g $TARGET_DIR/snapshot -czvf $A_NAME$FILES ${SOURCE_DIR[@]}
    [ "$CODE" == "0" ]&&echo -e "	\n\033[32mThese Add_Backup System Files $TARGET_DIR/$YEAR/$MONTH/$DAY/$A_NAME$FILES Backup Successfully !\033[0m"
  fi
}

sleep 3
Full_Backup;Add_Backup