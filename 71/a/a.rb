n = gets.to_i
n.times do |k|
    s = gets.chomp
    if s.size <= 10 then
        puts s
    else
        puts s[0,1] + (s.size - 2).to_s + s[-1,1]
    end
end
