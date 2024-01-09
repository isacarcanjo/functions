class Rectangle

  attr_accessor :width, :height

  def initialize(width, height)
    @width = width
    @height = height
  end

  def area
    @width * @height
  end

end

rec = Rectangle.new(4,4)

puts rec.area




"""@author: github.com/isacarcanjo"""



