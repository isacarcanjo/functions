function integral(a:number,b:number,f:any) {
    let sum=0;
    for(let i=a;i<=b;i++){
        sum+=f(i);
    }
    return sum;
}

console.log(integral(2,3,x=>x*x));




"""@author: github.com/isacarcanjo"""



