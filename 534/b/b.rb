#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

# b.rb

def main()
    v1, v2 = gets.chomp.split(" ").map{|s| next s.to_i }
    t, d = gets.chomp.split(" ").map{|s| next s.to_i }

    r = 0
    t.times do |tt|
        t1 = tt
        t2 = t - 1 - tt
        r += [v1 + d * t1, v2 + d * t2].min
    end
    puts r
    
end
    
if __FILE__ == $0 then
    main()
end
