
# Proxy Application

This project is a Go-based proxy application that can handle both TCP and UDP traffic, including HTTPS. It reads configuration from a JSON file and sets up listeners on specified IPs and ports to forward traffic to target IPs and ports.

## Features

- Proxy TCP and UDP traffic
- Forward HTTPS traffic without modification
- Configurable via a JSON file

## Configuration

The configuration file `config.json` should be placed in the same directory as the application or specified via a command-line argument. The file should contain an array of proxy configurations. Each configuration specifies the listen IP, source port, target IP, and target port.

### Example `config.json`

```json
{
  "proxies": [
    {
      "listen_ip": "0.0.0.0",
      "source_port": 8086,
      "target_ip": "10.0.0.1",
      "target_port": 8086
    },
    {
      "listen_ip": "0.0.0.0",
      "source_port": 80,
      "target_ip": "10.0.0.2",
      "target_port": 80
    },
    {
      "listen_ip": "127.0.0.1",
      "source_port": 443,
      "target_ip": "10.0.0.3",
      "target_port": 443
    }
  ]
}
```

## Building the Application

To build the application for the OpenWrt architecture, follow these steps:

```sh
# Set the target architecture for OpenWrt (e.g., MIPS, ARM)
export GOARCH=mipsle
export GOOS=linux

# Build the application
go build -o proxy_app main.go
```

## Deploying to OpenWrt

1. **Prepare the OpenWrt Environment:**

    ```sh
    # SSH into the OpenWrt device
    ssh root@<openwrt-ip>

    # Update package lists
    opkg update

    # Install necessary dependencies (if any)
    opkg install ca-certificates
    ```

2. **Deploy the Application:**

    ```sh
    # Transfer the built application to the OpenWrt device
    scp proxy_app root@<openwrt-ip>:/usr/bin/proxy_app

    # SSH into the OpenWrt device
    ssh root@<openwrt-ip>

    # Make the application executable
    chmod +x /usr/bin/proxy_app

    # Create a startup script
    cat << 'EOF' > /etc/init.d/proxy_app
    #!/bin/sh /etc/rc.common
    START=99
    STOP=10

    start() {
        /usr/bin/proxy_app -config /path/to/config.json &
    }

    stop() {
        killall proxy_app
    }
    EOF

    # Make the startup script executable
    chmod +x /etc/init.d/proxy_app

    # Enable the startup script
    /etc/init.d/proxy_app enable

    # Start the application
    /etc/init.d/proxy_app start
    ```

## Running the Application

To run the application, simply execute the built binary with the `-config` flag:

```sh
./proxy_app -config /path/to/config.json
```

The application will read the configuration from the specified `config.json` file and start listening on the specified IPs and ports.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
 