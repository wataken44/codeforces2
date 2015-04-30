#!/usr/bin/env ruby
# -*- coding: utf-8 -*-

# prepare_problem.rb

def main()
    argv = ARGV.shift
    mo = /(\d+)\/?([A-Z])\/?$/.match(argv)

    return if !mo

    set = mo[1]
    problem = mo[2].downcase

    script_dir = File.expand_path(File.dirname(__FILE__))

    system("ruby %s/prepare_directory.rb %s%s" % [script_dir, set, problem])
    pwd = Dir.pwd()
    Dir.chdir("%s/%s" % [set, problem])
    system("ruby %s/prepare_sample.rb %s%s" % [script_dir, set, problem])
    Dir.chdir(pwd)
end

if __FILE__ == $0 then
    main()
end
