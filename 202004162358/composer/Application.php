<?php
use Api\User;
use Service\User as User2;
class Application{
  public static function main(){
    self::registe();
    new User();
    new User2();
  }
  public static function registe(){
    spl_autoload_register("Application::loadClass");
  }
  public static function loadClass($class){
    $class=str_replace('\\', '/', lcfirst($class));
    $class="./".$class.".php";
    require_once $class;    
  }
}
Application::main();
//大小写的对应关系模糊不清