<?php
	function solve_quad ($a, $b, $c)
	{
		$d = $b*$b - 4*$a*$c;
		if ($d < 0)
			return (array(0));
		else if ($d == 0)
			return (array(-($b/2/$a)));
		else
			return (array(((-$b-sqrt($d))/2/$a), ((-$b+sqrt($d))/2/$a)));
	}
?>




"""@author: github.com/isacarcanjo"""



