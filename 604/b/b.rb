#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

# b.rb

def main()
    n, k = gets.chomp().split(" ").map{|si| next si.to_i() }
    s = gets.chomp().split(" ").map{|si| next si.to_i() }

    m = s[-1]
    t = n - k
    t.times do |tt|
        mm = s[tt] + s[2 * t - 1 - tt]
        m = mm if m < mm
    end
    puts m
end

if __FILE__ == $0 then
    main()
end
