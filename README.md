# Musicfreeier

Musicfreeier is a CLI application written in Go that aims to make music more accessible.

It acts as a layer over existing music downloading tools, exposing a simple and powerful command for retrieving your favorite music:

```bash
musicfreeier <url>
```

## Requirements

- [yt-dlp](https://github.com/yt-dlp/yt-dlp#installation) — must be installed and available in your `PATH`

```bash
# Linux (Debian/Ubuntu)
sudo apt install yt-dlp ffmpeg
```
