#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

# b.rb

def main()
    n, k = gets.chomp.split(" ").map{|s| next s.to_i }
    m = (n * n / 2) + (n % 2 == 1 ? 1 : 0)
    if m < k then
        puts("NO")
        return
    end

    puts("YES")
    r = 0
    n.times do |i|
        s = "S" * n
        t = i % 2
        while r < k && t < n do
            s[t] = "L"
            t += 2
            r += 1
        end
        puts(s)
    end
end

if __FILE__ == $0 then
    main()
end
