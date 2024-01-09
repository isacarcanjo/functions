<?php 
	function calculate($a,$b){
		if ($a == 0) {
			if ($b == 0) {
				echo "Phương trình có vô số nghiệm";
			}else{
				echo "Phương trình vô nghiệm";
			}
		}else{
			echo "Phương trình có nghiệm: " . -$b/$a;
		}
	}

	calculate(1,1);
	calculate(0,6);
	calculate(0,0);
	calculate(3,3);

 ?>




"""@author: github.com/isacarcanjo"""



