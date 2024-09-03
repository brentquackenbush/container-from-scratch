# Container From Scratch Project Using Go

**This project replicates the core functionality of `docker run` in a minimalistic way.**

💡 **WHY THIS MATTERS**:

If you’re a software engineer, understanding containers is essential. As of September 2024, containers are the standard for deploying public applications. Mastering this technology enables you to build, deploy, and scale applications efficiently, meeting the demands of modern software development.

> ⚠️ **Note:** On a Linux machine? You can run this locally via the CLI without Docker. If you aren't, follow these steps to run this application:

### 1. Install Docker on macOS

First, you need to install Docker on your MacBook:

1. Go to the [Docker Desktop for Mac website](https://www.docker.com/products/docker-desktop).
2. Download and install Docker Desktop by following the instructions.
3. Once installed, launch Docker Desktop, and ensure it’s running (you should see the Docker whale icon in the menu bar).

### 2. Build the Docker Image

Navigate to your project directory in the terminal and run the following command to build the Docker image:

```bash
docker build -t container-from-scratch .
```

This command tells Docker to build an image named `container-from-scratch` using the Dockerfile in the current directory.

### 3. Run the Docker Container

After building the image, you can run your Go application inside a Docker container:

>  📌  **Note:** The `--privileged flag` is needed due to Docker’s default security mechanisms.

```bash
docker run --rm -it --privileged container-from-scratch ./main run /bin/bash
```

- The `--rm` flag ensures the container is removed after it stops.
- The `-it` flags provide an interactive terminal so you can see the output.