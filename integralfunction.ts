function integral(func: (x:number)=>number, a:number, b:number, dx:number) {
    let sum:number = 0;
    for(let i = a; i < b; i+= dx) {
        sum += (func(i)+func(i+dx))/2*dx;
    }
    console.log(sum);
}

integral(function(x) {return x*x}, 0, 1, 0.0001)




"""@author: github.com/isacarcanjo"""



