## quickstart

```sh
go run github.com/acarlson99/home-automation -devices=devices.textproto -schedule=schedule.textproto
```

## Extension work

- [X] relative increase/decrease of color/brightness/color-temp
    - this in combination scheduling with crontab would allow for an over-time dimming effect (e.g. brightness -2% every minute for an hour)
- [ ] config variables
- [ ] more comprehensive support for avoiding race conditions
    - this is sort of in place, but elgato lights use `sync.Mutex` instead of batch execution
    - Should migrate towards batch execution
