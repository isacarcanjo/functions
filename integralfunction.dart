num integralFunct(num a, num b, num n, num(num x) f){
  num h = (b - a) / n;
  num S = 0;
  num x;
  for (var i = 1; i <= n; i++){
    x = a + (i - 0.5) * h;
    S = S + f(x);
  }
  S = h * S;
  return S;
}

main(){
  print(integralFunct(1, 2, 4, (num x) => x*x));
}




"""@author: github.com/isacarcanjo"""



