/*
The user enters two numbers in the variables a and b and then the program must
calculate the value of x according to the second degree equation 
*/

void calculateSecondEquation(double a, double b){
  // x = (-b) / 2a
  if(a != 0){
    double x = (-b) / (2 * a);
    print('x = $x');
  }else{
    print('a cant be 0 in this equation');
  }
}




"""@author: github.com/isacarcanjo"""



