# labcheck

A CLI tool for checking your self-hosted homelab health over SSH.

## Commands

- `labcheck init` — configure your homelab connection
- `labcheck ping` — test SSH connectivity
- `labcheck status` — interactive container status table

## Installation

```bash
go install github.com/rjdailey/labcheck@latest
```

## Usage

```bash
labcheck init     # first time setup
labcheck status   # view container status
```

### status keybindings

| Key   | Action         |
| ----- | -------------- |
| `↑/↓` | navigate       |
| `n`   | sort by name   |
| `s`   | sort by stack  |
| `h`   | sort by health |
| `u`   | sort by uptime |
| `q`   | quit           |

## Configuration

Stored at `~/.labcheck/config.yaml`.

## Roadmap

- [x] SSH connectivity check
- [x] Interactive container status TUI
- [ ] Disk usage
- [ ] SnapRAID sync status
- [ ] UniFi network stats
