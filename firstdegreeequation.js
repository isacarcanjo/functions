/*Write a javascript function that calculates the roots of a first degree equation
When you pass three parameters a, b, c to this function, it will calculate the roots of the equation ax + b = c.

For example, when you run the function with a = 3, b = 2, and c = 5, it will return the array [-1, -5/3]:*/



const root = (a, b, c) => {
    let x = -(b / a);
    let y = -(c / a);
    return [x, y];
};

console.log(root(5, 2, 0));




"""@author: github.com/isacarcanjo"""



