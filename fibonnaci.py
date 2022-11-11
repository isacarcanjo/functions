
"""

	Write a function that calculates fibonacci sequence

"""



def fibo(n):
	if n == 0:
		return 0
	elif n == 1:
		return 1
	else:
		return fibo(n-1)+fibo(n-2)



print(fibo(9))




"""@author: github.com/isacarcanjo"""