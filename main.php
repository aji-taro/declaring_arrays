<?php
echo "# php\n";

// 基本
$array01 = [1, 2, 3, 4];
echo "array01 ", var_export($array01, true), " count:", count($array01), "\n";

// 二次元配列
$array02 = [[10, 20, 30], [100, 200, 300],];
// あえてループしてみる
for ($i=0; $i<count($array02); $i++) {
    echo "array02[{$i}]";
    for ($k=0; $k<count($array02[$i]); $k++) {
        echo " " . $array02[$i][$k]; 
    }
    echo "\n";
}

// 要素の追加
$array03 = [11, 22, 33];
array_push($array03, 44, 55);
echo "array03 ", var_export($array03, true), " count:", count($array03), "\n";

// 連想配列
$array04 = [
    "goals" => 3,
    "assists" => 5,
    "duels" => 7,
    ];
echo "array04 ", var_export($array04, true), " count:", count($array04), "\n";

