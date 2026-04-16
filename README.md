# Compatdata Humanizer

Utility to automatically create human-readable symlinks to Steam compatdata folders.

## License

GPL-3.0-only (see [LICENSE](LICENSE)).

## Installation

1. Download the latest release from the [Releases](https://github.com/birdhimself/compatdata-humanizer/releases) page.
2. Copy the downloaded `compatdata-humanizer` to your preferred install location (`~/.local/bin` recommended)
3. Make sure the file is executable (`chmod +x ~/.local/bin/compatdata-humanizer`)

## Usage

```shell
compatdata-humanizer
```

## Configuration

On first run, a configuration file with comments will be created in your XDG configuration directory, for example
`~/.config/compatdata-humanizer/config.toml`.

Alternatively, you can use command line arguments, use `-h` for help.

Configuration precedence (per value): command line arguments > configuration file > defaults

## Development

### Requirements

- Go 1.26
- Make

### Commands

Run:

```shell
make run
```

Build for production:

```shell
make build
```

Build + install:

```shell
make install
```

Uninstall:

```shell
make uninstall
```

---

This project was created **without** the use of AI agents/"vibe coding". Suggested completions provided by JetBrains'
local inline completion model were sometimes applied.
