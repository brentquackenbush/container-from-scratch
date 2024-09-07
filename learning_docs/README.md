# Container From Scratch in Go

Welcome to the **Container from Scratch** project! This repository gives you a hands-on, in-depth understanding of how containers work at a low level using Go. Inspired by [Liz Rice's *Containers From Scratch*](https://www.youtube.com/watch?v=8fi7uSYlOdc), this project guides you through building a minimal container using Linux namespaces and cgroups.

By the end of this guide, you'll have built a working, container-like environment from scratch, learning to isolate processes and manage system resources with direct system calls.

## Table of Contents
- [Introduction](#introduction)
- [What Are Containers?](#what-are-containers)
- [Namespaces: Process Isolation](#namespaces-process-isolation)
  - [UTS Namespace](#uts-namespace)
  - [PID Namespace](#pid-namespace)
  - [Mount Namespace](#mount-namespace)
  - [Chroot and Chdir for Filesystem Isolation](#chroot-and-chdir-for-filesystem-isolation)
- [Control Groups (cgroups): Resource Management](#control-groups-cgroups-resource-management)
  - [Limiting Processes with the PIDs Subsystem](#limiting-processes-with-the-pids-subsystem)
  - [Automatic Cleanup Using notify_on_release](#automatic-cleanup-using-notify_on_release)
- [Proc Filesystem](#proc-filesystem)
- [Conclusion](#conclusion)
- [How to Run the Code](#how-to-run-the-code)

## Introduction
Containers power cloud computing and microservices by providing isolated environments for applications. But at its core, a container is just a process running in an isolated environment, with limited access to system resources.

While tools like Docker abstract these details, this project strips away those layers to show you how containers are built using Linux namespaces and cgroups.

## What Are Containers?
A container is a lightweight, executable package that includes everything needed to run an app: code, runtime, libraries, and system tools. Containers isolate apps from the host system, offering a consistent environment.

At a low level, containers use **namespaces** (for isolation) and **cgroups** (for resource control).

## Namespaces: Process Isolation
Namespaces isolate system resources like process IDs and filesystems, making a container appear as if it's running on a separate system.

### UTS Namespace
The UTS namespace gives the container its own hostname, isolated from the host.

```go
syscall.Sethostname([]byte("container"))
```

### PID Namespace
The PID namespace isolates the container's process ID space. From inside, the container only sees its own processes, starting with PID 1.

```go
syscall.CLONE_NEWPID
```

### Mount Namespace
The Mount namespace isolates filesystem mount points, giving each container its own view of the filesystem.

```go
syscall.CLONE_NEWNS
```

### Chroot and Chdir for Filesystem Isolation
To isolate the container's filesystem, we use `chroot` and `chdir`.

```go
syscall.Chroot("/path/to/container/root")
syscall.Chdir("/")
```

## Control Groups (cgroups): Resource Management
Cgroups limit and monitor resources (CPU, memory, processes) a process can use. This prevents a container from consuming excessive resources.

### Limiting Processes with the PIDs Subsystem
We use the PIDs subsystem of cgroups to limit the number of processes inside the container.

```go
cgroups := "/sys/fs/cgroup/"
pids := filepath.Join(cgroups, "pids")
os.Mkdir(filepath.Join(pids, "liz"), 0755)
ioutil.WriteFile(filepath.Join(pids, "liz/pids.max"), []byte("20"), 0700)
```

### Automatic Cleanup Using notify_on_release
We enable automatic cgroup cleanup by setting the `notify_on_release` flag.

```go
ioutil.WriteFile(filepath.Join(pids, "liz/notify_on_release"), []byte("1"), 0700)
```

## Proc Filesystem
The `/proc` filesystem provides information about running processes. To simulate a complete environment, we mount `/proc` inside the container.

```go
syscall.Mount("proc", "proc", "proc", 0, "")
syscall.Unmount("proc", 0)
```

## Conclusion

Building a container from scratch in Go teaches the fundamental concepts behind containerization:

- **Process Isolation**: Use namespaces to isolate PIDs, hostnames, and filesystems.
- **Resource Management**: Use cgroups to control resources like CPU and memory.
- **Filesystem Isolation**: Use chroot to restrict access to a specific filesystem.
- **Proc Filesystem**: Simulate a Linux environment by mounting /proc.
