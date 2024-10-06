# fsort

`fsort` is a command-line utility written in Go for sorting large text files efficiently. It supports various sorting options and integrates seamlessly with other command-line tools via piping.

## Features

- **External Sorting**: Handles very large files that do not fit into memory.
- **Flag Support**:
  - `-n, --numeric`: Sort numerically.
  - `-r, --reverse`: Reverse the sort order.
  - `-u, --unique`: Remove duplicate lines.
  - `-f, --ignore-case`: Case-insensitive sorting.
  - `-c, --chunk-size`: Number of lines per chunk (default: 100000).
- **Piping Support**: Can be used with other command-line utilities like `head`, `tail`, etc.
