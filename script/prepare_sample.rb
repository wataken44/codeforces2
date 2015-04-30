#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

# prepare_sample.rb

require 'open-uri'

def main()
    arg = ARGV.shift
    url = nil
    mo = /(\d+)\/?([A-Za-z])$/.match(arg)
    if mo then
        url = "http://codeforces.com/problemset/problem/%s/%s" %
              [mo[1], mo[2].upcase]
    else
        url = arg
    end
    fp = open(url)

    fp.each_line do |line|
        next if !line.index("Sample test")
        buf = line.sub(/.*<div class="sample-test">/, "")
        arr = buf.scan(/<pre>(.*?)<\/pre>/)
        arr = arr.map{|it| next it[0].gsub("<br />", "\n")}
        (arr.size / 2).times do |i|
            ofp = open("sample%d.in.txt" % (i+1), "w")
            ofp.puts(arr[2*i])
            ofp.close()
            ofp = open("sample%d.out.txt" % (i+1), "w")
            ofp.puts(arr[2*i+1])
            ofp.close()            
        end
    end
    
    fp.close()
end

if __FILE__ == $0 then
    main()
end
