
/*
===============
Second Degree
===============

Instruksi
=========
Buatlah sebuah function yang menerima parameter dengan tipe data number bernama a, b, dan c.
Function secondDegreeEquation akan menghitung akar dari persamaan kuadrat dengan parameter tersebut.

Persamaan kuadrat adalah:
a * x^2 + b * x + c = 0

Jika tidak ada penyelesaian, maka function akan mereturn "No Solution".
Jika ada 2 penyelesaian, maka function akan mereturn array multidimensi yang berisi 2 penyelesaian.
Jika ada 1 penyelesaian, maka function akan mereturn array yang berisi 1 penyelesaian.

Contoh ada 2 penyelesaian:
[ [x1, x2], [x1, x2] ]

Contoh ada 1 penyelesaian:
[ [x] ]

Contoh tidak ada penyelesaian:
No Solution

RULE:
- WAJIB menggunakan algoritma/pseudocode (boleh tidak menggunakan algoritma/pseudocode jika sudah mengerti algoritma/pseudocode)
- Dilarang menggunakan Math.pow, Math.sqrt, atau built-in function yang lain yang mungkin dapat menyelesaikan masalah ini (seperti yang ada di bilangan.js).

*/

function secondDegreeEquation(a, b, c) {
  var arr = []
  var arr2 = []