# DataGate

DataGate is a lightweight, no-fuss file-sharing server that I initially built for my own use cases. Instead of dealing with SFTP, SCP, or setting up complex web servers, DataGate provides an easy way to browse, download, and share files over a simple HTTP interface.

## Features

- 🔹 **Minimal Setup** – Just run the binary, and you're good to go.
- 🔹 **Token Authentication** – Secure your file server with an authentication token.
- 🔹 **Download & Browse** – Navigate and download files through a clean web interface.
- 🔹 **Cross-Platform** – Works on Windows, Linux, and macOS.
- 🔹 **No Dependencies** – A single executable, no need to install anything extra.

---

## Getting Started

### Download Prebuilt Binaries

Precompiled binaries are available for Windows, Linux, and macOS.
You can download them from the [Releases](https://github.com/yourusername/datagate/releases) page.

### Running DataGate

Once downloaded, simply run the binary:

```sh
./datagate -d /path/to/serve -p port
```
