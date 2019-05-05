# Fractostatus
Centralized system monitor written in GO. Reports that a system is alive and whether the selected processes is running to the main server. The server is also the location where you will be checking the report.

## Compatibility
#### Linux and most Unix-like
* **Printer mode**
* **Client mode**
* **Server mode**

#### Windows
* **Printer mode**
* _Client mode_ (partial support, main functionality works but process reporting doesn't.)
* **Server mode**

## Usage
Either build it from source or download a prebuilt [release.](https://github.com/bennystarfighter/fractostatus/releases)

#### Printer mode
```bash
./fractostatus
```

#### Client mode
```bash
./fractostatus --client
```

#### Server mode
```bash
./fractostatus --server
```

### Config file
The standard format is in YAML but you can use JSON, TOML etc. THe important thing is that the base name of the file is ```config```

*Paths:* ```$HOME/.config/fractostatus``` , ``` . ```

#### Example
```yaml
# Client mode configuration
# The identifier is what the main server will know this client as.
# Anything but: client-list
identifier: Main-PC
server-address: http://10.0.0.1:8888
server-password: password
# How often the client will send data to the server (in seconds).
# A good value is between 10-600 and not above the servers alive-timeout.
pollrate: 10
# For which processes you want to report status of.
process-watch:
- htop
- gotop

# Server mode configuration
port: 8888
password: password
# The max amount of seconds since a client updated the server until it will be considered dead.
alive-timeout: 60
# Also known as SSL
TLS: false
certfile-path: $HOME/certs/cert
keyfile-path: $HOME/certs/key
```
