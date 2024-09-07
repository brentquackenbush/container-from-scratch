# Containers From Scratch in Go

**Replicates the core of `docker run` in Go.**

💡 **WHY THIS MATTERS**:
Understanding containers is key for modern software deployment. Containers streamline building, scaling, and running apps, making you a better developer.

> ⚠️ **Linux user?** You can run this directly via CLI. Not on Linux? Here's how to run it with Docker:

### 1. Install Docker 
- Download Docker from [Docker Desktop](https://www.docker.com/products/docker-desktop).
- Install and launch Docker Desktop.

### 2. Build the Docker Image
Run this in your project directory:

```bash
docker build -t container-from-scratch .
```

### 3. Run the Container

After building the image, you can run your Go application inside a Docker container:

>  📌  **Note:** The `--privileged flag` is needed due to Docker’s default security mechanisms.

```bash
docker run --rm -it --privileged container-from-scratch ./main run /bin/bash
```

- `--rm`: Removes the container after stopping.
- `-it`: Opens an interactive terminal.
