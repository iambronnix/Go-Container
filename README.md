# Go-Container 🐳
A minimal container runtime built from scratch in Go, without relying on Docker or external C libraries. This project is a deep dive into Linux kernel primitives.

## Features Implemented
*   **Process Isolation:** Utilizing Linux Namespaces (`CLONE_NEWUTS`, `CLONE_NEWPID`, `CLONE_NEWNS`).
*   **Resource Limiting:** Implementing Cgroups v2 to restrict memory and CPU.
*   **Filesystem Isolation:** Using `pivot_root` or `chroot` to create a custom root filesystem.
*   **Networking:** Creating virtual ethernet (veth) pairs and connecting them to a bridge.

## How to Run
\`\`\`bash
go build -o gocontainer main.go
sudo ./gocontainer run /bin/bash
\`\`\`

## What I Learned
*   How Docker actually works under the hood.
*   System calls: `clone()`, `unshare()`, `pivot_root()`.
*   Memory management in Go and the `syscall` package.

