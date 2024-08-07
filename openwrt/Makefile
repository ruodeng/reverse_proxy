include $(TOPDIR)/rules.mk

PKG_NAME:=proxy_app
PKG_VERSION:=1.0
PKG_RELEASE:=1

include $(INCLUDE_DIR)/package.mk

define Package/proxy_app
  SECTION:=utils
  CATEGORY:=Utilities
  TITLE:=Proxy App
  DEPENDS:=+luci
endef

define Package/proxy_app/description
  A proxy application with LuCI interface for OpenWrt.
endef

define Build/Compile
  # Nothing to compile
endef

define Package/proxy_app/install
  $(INSTALL_DIR) $(1)/etc/config
  $(INSTALL_CONF) ./files/etc/config/proxy_app $(1)/etc/config/proxy_app

  $(INSTALL_DIR) $(1)/etc/init.d
  $(INSTALL_BIN) ./files/etc/init.d/proxy_app $(1)/etc/init.d/proxy_app

  $(INSTALL_DIR) $(1)/usr/bin
  $(INSTALL_BIN) ./files/usr/bin/generate_proxy_config.sh $(1)/usr/bin/generate_proxy_config.sh

  $(INSTALL_DIR) $(1)/usr/lib/lua/luci/controller
  $(INSTALL_DATA) ./files/usr/lib/lua/luci/controller/proxy_app.lua $(1)/usr/lib/lua/luci/controller/proxy_app.lua

  $(INSTALL_DIR) $(1)/usr/lib/lua/luci/model/cbi
  $(INSTALL_DATA) ./files/usr/lib/lua/luci/model/cbi/proxy_app.lua $(1)/usr/lib/lua/luci/model/cbi/proxy_app.lua
endef

$(eval $(call BuildPackage,proxy_app))