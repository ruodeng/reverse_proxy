module("luci.controller.proxy_app", package.seeall)

function index()
    entry({"admin", "services", "proxy_app"}, cbi("proxy_app"), _("Proxy App"), 100).dependent = true
    entry({"admin", "services", "proxy_app", "status"}, call("action_status")).leaf = true
    entry({"admin", "services", "proxy_app", "start"}, call("action_start")).leaf = true
    entry({"admin", "services", "proxy_app", "stop"}, call("action_stop")).leaf = true
    entry({"admin", "services", "proxy_app", "restart"}, call("action_restart")).leaf = true
end

function action_status()
    local status = luci.sys.call("pgrep proxy_app > /dev/null") == 0
    luci.http.prepare_content("application/json")
    luci.http.write_json({ running = status })
end

function action_start()
    luci.sys.call("/etc/init.d/proxy_app start")
    luci.http.redirect(luci.dispatcher.build_url("admin/services/proxy_app"))
end

function action_stop()
    luci.sys.call("/etc/init.d/proxy_app stop")
    luci.http.redirect(luci.dispatcher.build_url("admin/services/proxy_app"))
end

function action_restart()
    luci.sys.call("/etc/init.d/proxy_app restart")
    luci.http.redirect(luci.dispatcher.build_url("admin/services/proxy_app"))
end