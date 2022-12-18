
//! Write a function in dart that calculates integral function

import 'dart:math';

double f(double x) => 3.0 * sqrt(x);

double f_int(double x) => (9 / 2) * sqrt(x) - (9 / 2) * x;

double integral(Function f, double a, double b, double dx) {
  double sum = 0;
  while(a < b) {
    a += dx;
    sum += (dx * f(a));
  }
  return sum;
}

void main() {
  print(integral(f, 0, 1, 0.01));
  print(f_int(1));
}


