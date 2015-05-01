#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

# c.rb

def main()
    n, m = gets.chomp.split(" ").map{|s| next s.to_i }
    d = [0] * m
    h = [0] * m

    m.times do |i|
        d[i], h[i] = gets.chomp.split(" ").map{|s| next s.to_i }
    end

    r = h[0]
    1.upto(m-1) do |i|
        dd = d[i] - d[i-1]
        dh = h[i] - h[i-1]
        if dd.abs < dh.abs then
            puts("IMPOSSIBLE")
            return
        else
            r = [r, h[i]].max
            r = [r, [h[i], h[i-1]].min + (dd.abs + dh.abs) / 2 ].max
        end
    end
    r = [r, h[-1] + (n - d[-1])].max
    r = [r, h[0] + (d[0] - 1)].max
    puts r
end

if __FILE__ == $0 then
    main()
end
