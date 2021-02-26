<?php 
namespace App\Model;

class User extends AbstructModel{
	protected $tableName = 'user';
	protected $connectionName = 'write';//write|read
	protected $autoTimeStamp = true;
	protected $createTime = 'create_at';
	protected $updateTime = 'update_at';

}