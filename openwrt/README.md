Building the Package
Place the Makefile in the openwrt/package/proxy_app/ directory.
Place the configuration, init script, and LuCI files in the appropriate directories under openwrt/package/proxy_app/files/.
Navigate to the OpenWrt build root directory.
Run the following commands to build the package:
make package/proxy_app/compile V=s
This will generate a .ipk package in the bin/packages directory, which can be installed on an OpenWrt device using opkg.