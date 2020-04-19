<?php

ini_set('display_errors', 1);

$doc= new DOMDocument();
$doc->load('2.xml');

$years = $doc->getElementsByTagName("year");

//对每一个年份进行处理
foreach($years as $year){
	//var_dump($year);
	//获得具体的年份值
	echo $year->getElementsByTagName('yearName')->item(0)->nodeValue.'年<br>';
	//获得该年份下的所有假日
	$holidays = $year->getElementsByTagName('holiday');
	foreach ($holidays as $holiday) {
		$holidayName =  $holiday->getElementsByTagName('holidayName')->item(0)->nodeValue;
		//echo iconv('utf-8','gb2312',$holidayName).': <br>';
		echo $holidayName.': <br>';
	
		//获得假日的具体开放日期
		$from = $holiday->getElementsByTagName('daysOff')->item(0)->getElementsByTagName('from')->item(0)->nodeValue;
		$to =  $holiday->getElementsByTagName('daysOff')->item(0)->getElementsByTagName('to')->item(0)->nodeValue;
		echo "假期为:{$from}~{$to}<br>";

		//获得针对该假日的调休时间
		$days = $holiday->getElementsByTagName('overTime')->item(0)->getElementsByTagName('day');
		if($days->length!=0){
			echo '调休日为:';
			foreach ($days as $day) {
			echo $day->nodeValue.'<br/>';
			}
		}
		echo '<br>';
	}

}
?>