# HTTP Proxy Example

This example demonstrates how to use the Databricks SDK for Go with an HTTPS proxy. The `proxy` directory contains a simple HTTP proxy server that forwards requests to the Databricks REST API. The `client` directory contains a simple client that uses the Databricks SDK for Go to make requests to the proxy server.

The proxy server generates a self-signed certificate and listens on port 8443 with this certificate. The generated certificate can be found at `proxy/proxy.crt`, and it is cleaned up when the proxy is terminated. The client must trust this certificate to make requests to Databricks through the proxy server.

## Run the Example: Auto-registration of the Proxy Server Certificate

This example demonstrates how to automatically register the proxy server's certificate with the system's certificate store. This allows the client to trust the proxy server's certificate without needing to manually configure the client to trust the certificate. This is the preferred way to configure your system, as it doesn't require application-specific configuration to trust the proxy server's certificate.

Note: this example requires root access to register the proxy server's certificate with the system's certificate store. On Windows, this requires running the proxy server as an administrator.

1. Run the proxy server:

   ```bash
   go run ./proxy --register-certificate
   ```

   This will attempt to register the proxy server's certificate with the system's certificate store. If you're using Windows, you will need to run Powershell as an administrator to register the certificate.

2. In another terminal, run the client:

   ```bash
   HTTPS_PROXY=https://localhost:8443 go run ./client
   ```

   or on Windows:

   ```powershell
   $env:HTTPS_PROXY="https://localhost:8443" go run ./client
   ```

   Note that on Windows, this does not require administrator privileges.

3. When done, terminate the proxy server by typing `CMD+C` or `Ctrl+C` in the terminal where it is running.

## Run the Example: Configure the Proxy Server Certificate at Runtime

This example demonstrates how to configure the client to trust the proxy server's certificate at runtime. This is useful when you don't have root access to register the proxy server's certificate with the system's certificate store, or when you want to avoid modifying the system's certificate store.

1. Run the proxy server:

   ```bash
   go run ./proxy
   ```

2. In another terminal, run the client, this time specifying the path to the CA file:

   ```bash
   DATABRICKS_CA_BUNDLE=proxy/proxy.crt HTTPS_PROXY=https://localhost:8443 go run ./client
   ```

   or on Windows:

   ```powershell
   $env:DATABRICKS_CA_BUNDLE="proxy/proxy.crt" $env:HTTPS_PROXY="https://localhost:8443" go run ./client
   ```

3. When done, terminate the proxy server by typing `CMD+C` or `Ctrl+C` in the terminal where it is running.
