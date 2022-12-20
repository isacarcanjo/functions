
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
        $x += $h;
    }
    return $diff_value;
}

$diff_value = diff(0, 4, 0.1, 'diff_function');

echo $diff_value;
?>

