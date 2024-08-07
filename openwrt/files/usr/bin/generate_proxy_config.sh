#!/bin/sh

config_load proxy_app
config_get listen_ip proxy listen_ip
config_get source_port proxy source_port
config_get target_ip proxy target_ip
config_get target_port proxy target_port

cat <<EOF > /etc/config/proxy_app.json
{
  "proxies": [
    {
      "listen_ip": "$listen_ip",
      "source_port": $source_port,
      "target_ip": "$target_ip",
      "target_port": $target_port
    }
  ]
}
EOF