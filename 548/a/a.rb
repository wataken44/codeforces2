#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

# a.rb

def is_palindrome(s)
    s.size.times do |i|
        return false if s[i] != s[s.size - i - 1]
    end
    return true
end

def ok(s, k)
    return false if s.size % k != 0
    n = s.size / k
    k.times do |i|
        return false if !is_palindrome(s[n * i, n])
    end
    return true
end

def main()
    s = gets.chomp
    k = gets.to_i
    
    puts ok(s, k) ? "YES" : "NO"
end


if __FILE__ == $0 then
    main()
end
