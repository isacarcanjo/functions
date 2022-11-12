puts "this is a calculator for integral"

def integral(f, a, b)
    deltax=0.001
    deltay=f.call(0.001)
    area = 0
    while a < b
        area += deltay*deltax
        a += deltax
    end
    return area
end

f = lambda {|x| (x**2-x)/(x**2-2*x+2)}

puts integral(f, 0, 2)




"""@author: github.com/isacarcanjo"""



