
def calculate_integral(x, y)
	a = Hash.new
	File.open(y,'r').each do |line|
		b = line.split
		a[b[0].to_sym] = b[1].to_f
	end	

	answer = 0.0
	a.each do |key, value|
		answer += value * x ** key
	end	

	return answer
end	
