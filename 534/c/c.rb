#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

# c.rb

def main()
    n, a = gets.chomp.split(" ").map{|s| next s.to_i }
    d = gets.chomp.split(" ").map{|s| next s.to_i }
    s = d.inject{|ret, i| next ret + i }
    r = []
    d.each_with_index do |di, i|
        ss = s - di
        least = [1, a - ss].max
        most = [di, a - (n - 1)].min
        cand = most - least + 1
        r << di - cand
    end
    
    puts(r * " ")
end
    
if __FILE__ == $0 then
    main()
end
