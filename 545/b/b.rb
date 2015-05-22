#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

# b.rb

def main()
    s = gets.chomp
    t = gets.chomp
    v = 0
    s.size.times do |i|
        if s[i] != t[i] then
            s[i] = [s[i],t[i]][v % 2]
            v += 1
        end
    end
    if (v % 2) == 0 then
        puts s
    else
        puts "impossible"
    end
end

if __FILE__ == $0 then
    main()
end
