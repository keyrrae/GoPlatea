total = 0
for i in 0..999
   for j in 1..999
	total += (i / j)
   end
end
puts "s #{total}"
