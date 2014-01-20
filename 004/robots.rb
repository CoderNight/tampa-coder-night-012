while t = ARGF.gets
  x = ARGF.gets.chomp.partition("X")[0].size
  b = ARGF.gets
  ARGF.gets
  puts ["GO WEST","GO WEST","GO EAST"][t.chomp.chars.map.with_index { |c,i| (i+x+1)%2*(i<=>x)*(c<=>'#') }.reduce(:+)+b.chomp.chars.map.with_index { |c,i| (i+x)%2*(i<=>x)*(c<=>'#') }.reduce(:+) <=> 0]
end
