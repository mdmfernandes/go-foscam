# go-foscam

Go Library for Foscam IP Cameras

## Supported Cameras

- FI9800P
- FI8919W

## Supported Functionalities

- ChangeMotionStatus
- GetMotionStatus (WIP)

## Run the example

```bash
$ cd go-foscam

$ cat << EOF > .envrc
export FOSCAM_URL=http(s)://<host>:<port>
export FOSCAM_USER=<user>
export FOSCAM_PASSWORD=<password>
EOF

make run/example
```

## To Do

- [ ] `staticcheck` *vs.* `golangci-lint` - to replace `go vet`
- [ ] Run audit and coverage in CI
- [ ] Add unit tests for current cameras
- [ ] Add more Functionalities
- [ ] Support more cameras
- [ ] Complete TODOs from code
