# Fabric Docker Image

This directory provides a simple Docker setup for running the [Fabric](https://github.com/danielmiessler/fabric) CLI.

## Build

Build the image from the repository root:

```bash
docker build -t fabric -f scripts/docker/Dockerfile .
```

## Persisting configuration

Fabric stores its configuration in `~/.config/fabric/.env`. Mount this path to keep your settings on the host.

### Using a host directory

```bash
mkdir -p $HOME/.fabric-config
# Run setup to create the .env and download patterns
 docker run --rm -it -v $HOME/.fabric-config:/root/.config/fabric fabric --setup
```

Subsequent runs can reuse the same directory:

```bash
docker run --rm -it -v $HOME/.fabric-config:/root/.config/fabric fabric -p your-pattern
```

### Mounting a single .env file

If you only want to persist the `.env` file:

```bash
# assuming .env exists in the current directory
docker run --rm -it -v $PWD/.env:/root/.config/fabric/.env fabric -p your-pattern
```

## Running the server

Expose port 8080 to use Fabric's REST API:

```bash
docker run --rm -it -p 8080:8080 -v $HOME/.fabric-config:/root/.config/fabric fabric --serve
```

The API will be available at `http://localhost:8080`.
