def integral(a,b,h)
  x = a
  result = 0
  while x < b do
    result += (h*(x*x+x+1))
    x += h
  end
  return result
end




"""@author: github.com/isacarcanjo"""



