#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

# a.rb

def main()
    n = gets.to_i
    r = []
    n.times do |i|
        a = gets.chomp.split(" ").map{|s| next s.to_i }
        r << (i + 1) if !(a.index(1) || a.index(3))
    end
    puts r.size
    puts r * " " if r.size > 0
end

if __FILE__ == $0 then
    main()
end
