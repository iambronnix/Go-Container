# Go-Container 🐳

A minimal Linux container runtime built from scratch in Go. This project is a deep dive into Linux kernel primitives, demonstrating how tools like Docker work under the hood without relying on external container libraries.

## ⚠️ Project Status: Educational Prototype
This is an educational project designed to understand container internals. It requires a Linux host with root privileges. **Do not use this in production.**

## 🏗️ Architecture & Primitives Used
- **Process Isolation:** Linux Namespaces (`CLONE_NEWPID`, `CLONE_NEWNS`, `CLONE_NEWUTS`).
- **Resource Limiting:** Cgroups v2 (`memory.max`, `pids.max`).
- **Filesystem Isolation:** `pivot_root` and `chroot` to isolate the container root filesystem.
- **Process Supervision:** A parent process sets up the namespaces, and a child process executes the user command within the isolated environment.

## 🚀 How to Run
1. Build the runtime:
   ```bash
   go build -o gocontainer main.go

## What I Learned
*   How Docker actually works under the hood.
*   System calls: `clone()`, `unshare()`, `pivot_root()`.
*   Memory management in Go and the `syscall` package.

