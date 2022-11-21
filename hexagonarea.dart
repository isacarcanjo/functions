
import 'dart:io'; // for system input and output

double calculateArea(double s){
  return ((3 * sqrt(3) * pow(s, 2)) / 2).roundToDouble();
}

main() {
  print("Enter the side of the hexagon:");
  double side = double.parse(stdin.readLineSync());

  print("The area of the hexagon is ${calculateArea(side)}");
}
