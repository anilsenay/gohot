# Go Hot Reload Example

I implemented a basic hot reloading tool for working with static html pages in Golang. It's better to use some live reloading tool such as [air](https://github.com/cosmtrek/air) currently. I am planning to implement a file watcher as well.

![ScreenRecording2024-01-15at00 53 35-ezgif com-video-to-gif-converter](https://github.com/anilsenay/go-hot-reload/assets/1047345/32120213-cd52-4328-9d45-a7cf05d804af)

## Installation
```sh
go install github.com/anilsenay/gohot
```

## Usage
### Start hot reload server:

Default:
```sh
gohot start
```
With options:
```sh
gohot start --port 3000 --proxy http://localhost:8080
```

### Refresh the page:

Default:
```sh
gohot refresh
```
With options:
```sh
gohot refresh --port 3000
```
