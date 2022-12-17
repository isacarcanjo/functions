
def calculate_integral(x, y)
	a = Hash.new
	File.open(y,'r').each do |line|
		b = line.split