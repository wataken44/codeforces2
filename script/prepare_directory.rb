#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

# prepare_directory.rb

def main()
    argv = ARGV.shift
    mo = /(\d+)\/?([A-Za-z])\/?$/.match(argv)

    return if !mo

    set = mo[1]
    problem = mo[2].downcase
    Dir.mkdir(set) if !FileTest.exists?(set)
    Dir.mkdir(set + "/" + problem) if !FileTest.exists?(set + "/" + problem)
end

if __FILE__ == $0 then
    main()
end
