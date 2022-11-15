<?php
	function calculate($x, $y, $z)
	{
		if ($x != 0)
			return (($y * -1) + sqrt(($y * $y) - (4 * $x * $z))) / (2 * $x);
		else
			return -1;
	}
	echo calculate(0, 9, 3);
?>




"""@author: github.com/isacarcanjo"""



