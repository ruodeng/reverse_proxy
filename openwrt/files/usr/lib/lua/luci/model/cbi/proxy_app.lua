m = Map("proxy_app", "Proxy App Configuration")

s = m:section(TypedSection, "proxy", "Proxy Settings")
s.addremove = false
s.anonymous = true

o = s:option(Value, "listen_ip", "Listen IP")
o.datatype = "ipaddr"
o.default = "0.0.0.0"

o = s:option(Value, "source_port", "Source Port")
o.datatype = "port"
o.default = "8086"

o = s:option(Value, "target_ip", "Target IP")
o.datatype = "ipaddr"
o.default = "10.0.0.1"

o = s:option(Value, "target_port", "Target Port")
o.datatype = "port"
o.default = "8086"

return m