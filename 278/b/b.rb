#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

# b.rb

def main()
    n = gets.to_i
    s = []
    n.times do |k|
        s << gets.chomp
    end

    l = 1
    while true do
        m = []
        n.times do |k|
            next if s[k].size < l
            (s[k].size - l + 1).times do |a|
                m << s[k][a, l]
            end
        end
        m = m.sort.uniq
        if m.size == 26 ** l then
            l += 1
            next
        end
        c = [*"a".."z"]
        (26 ** l).times do |i|
            ii = i
            ss = ""
            l.times do |b|
                ss = c[ii % 26] + ss
                ii = (ii - ii % 26) / 26
            end
            if !m.index(ss) then
                puts ss
                return
            end
        end
    end
end

if __FILE__ == $0 then
    main()
end
