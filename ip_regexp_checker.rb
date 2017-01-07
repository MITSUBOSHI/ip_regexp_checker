#! /bin/sh
exec ruby -S -x $0 "$@"
#! ruby

target_regexp = ARGV[0]
target_ip_parts = ARGV[1]
exists_until_first_oct = /^\d{1,3}$/
exists_until_second_oct = /^\d{1,3}\.\d{1,3}$/

target_ip_addresses = if Regexp.new(exists_until_first_oct) === target_ip_parts then
  #0.*.*.*
  (0..255).to_a.map{|second_oct|
    (0..255).to_a.map{|third_oct|
      (0..255).to_a.map{|fourth_oct|
        [target_ip_parts, second_oct, third_oct, fourth_oct].join(".")
      }
    }
  }.flatten
elsif Regexp.new(exists_until_second_oct) === target_ip_parts then
  #0.0.*.*
  (0..255).to_a.map{|third_oct|
    (0..255).to_a.map{|fourth_oct|
      [target_ip_parts, third_oct, fourth_oct].join(".")
    }
  }.flatten
else
  #0.0.0.0
  [target_ip_parts]
end

ip_addresses = target_ip_addresses.map{|ip_address|
  Regexp.new(target_regexp) === ip_address ? ip_address : nil
}.compact

puts ip_addresses
