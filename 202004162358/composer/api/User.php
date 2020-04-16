<?php
namespace Api;
  
use Service\Login;

class User{
  public function __construct(){
    echo "User类<br/>";
    new Login();
    new Account();
  }
}