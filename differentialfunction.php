
<?php
function diff_function($x, $y) {
    return $y * ($x - $y);
}

function diff($a, $b, $h, $dif_function) {
    $diff_value = 0;
    $x = $a;
    $result = [];
    while ($x <= $b) {
        $y = call_user_func($dif_function, $x, $h);
        if (abs($y) < $h) {
            $diff_value += ($y * $h);
        }