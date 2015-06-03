#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

# c.rb

def main()
    n = gets.to_i

    x0, h0 = gets.chomp.split(" ").map{|s| next s.to_i}
    
    x = [x0-h0-1,x0]
    h = [0,h0]

    (n-1).times do |i|
        xi, hi = gets.chomp.split(" ").map{|s| next s.to_i}
        x << xi
        h << hi
    end

    lm = 0
    rm = 0
    1.upto(n) do |i|
        mm = [lm, rm].max
        ln = ( x[i] - x[i-1] > h0 ) ? (lm + 1) : lm
    end
    
end

if __FILE__ == $0 then
    main()
end
