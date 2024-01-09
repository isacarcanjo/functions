def fib(n):
    if n<=1:
        return n
    else:
        return fib(n-1) + fib(n-2)

n=int(input("enter a number:"))
print("the fibonacci number of",n,"is",fib(n))




"""@author: github.com/isacarcanjo"""



