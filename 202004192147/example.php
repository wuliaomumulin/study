<?php 
//错误提示
ini_set('display_errors', 1);

read();

/**
 * 文档转数组的话属性不好处理
 * xml格式字符串
 */
function xmlConversionArray($str){
$data = json_decode(json_encode(simplexml_load_string($str)),TRUE);
return $data;
}

/**
 * 文档读取
 */
function read(){
	$doc = new DOMDocument();
	$doc->load("1.xml");
	$plays = $doc->getElementsByTagName("play");

	foreach ($plays as $play) {
		$titles = $play->getElementsByTagName('title');
		$title = $titles->item(0)->nodeValue;
		//var_dump($titles->item(0));

		$attributeName = $titles->item(0)->attributes->item(0)->nodeName;
		$attributeValue = $titles->item(0)->attributes->item(0)->nodeValue;

		$authors = $play->getElementsByTagName('author');
		$author = $authors->item(0)->nodeValue;

		$details = $play->getElementsByTagName('detail');


		echo "{$title} [{$attributeName} = {$attributeValue}]: {$author} <br>";
	}	
}

/*
 * 文档写入
 */
function write(){
	$local_array=array(
    	array("pid"=>"1", "name"=>"kitty","sex"=>"female"),
    	array("pid"=>"2", "name"=>"tom","sex"=>"male"),
	);

	$doc = new DOMDocument();
	//创建根节点
	$root = $doc->createElement('root');
	$root = $doc->appendChild($root);

	foreach($local_array as $a){
		$table_id = 'person';
		$occ = $doc->createElement($table_id);
		$occ = $root->appendChild($occ);
		$fieldname = 'pid';
		$child = $doc->createElement($fieldname);
		$child = $occ->appendChild($child);
		$fieldvalue = $a['pid'];
		$value = $doc->createTextNode($fieldvalue);
		$child->appendChild($value);
	
		$fieldname = 'name';
		$child = $doc->createElement($fieldname);
		$child = $occ->appendChild($child);
	
		//写入属性
		$child->setAttribute('attr1','attr1Value');
		$child->setAttribute('attr2','attr2Value');

		$fieldValue = $a['name'];
		$value = $doc->createTextNode($fieldValue);
		$child->appendChild($value);

		$fieldName = 'sex';
		$child = $doc->createElement($fieldName);
		$child = $occ->appendChild($child);
		$fieldValue = $a['sex'];
		$value = $doc->createTextNode($fieldValue);
		$child->appendChild($value);

		//我们在追加一组
		$fieldName ='age';
		$child = $doc->createElement($fieldName);
		$child = $occ->appendChild($child);
		$value = $doc->createTextNode(rand(10,100));
		$child->appendChild($value);
	}

	echo file_put_contents(basename(__FILE__,'.php').".xml",$doc->saveXML());
}


?>