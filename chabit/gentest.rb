#!/usr/bin/env ruby


ops = ["set_a", "set_b", "get_c"]

n = ARGV[0].to_i
q = ARGV[1].to_i

puts "#{n} #{q}"
a = ""
b = ""
(0...n).each do
  a << (rand(2) + 48)
  b << (rand(2) + 48)
end
puts a
puts b

(0...q).each do
  op = ops[rand(3)]
  if op == "get_c"
    puts "#{op} #{rand(n)}"
  else
    puts "#{op} #{rand(n)} #{rand(2)}"
  end
end
