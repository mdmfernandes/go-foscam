# go-foscam

[![tests](https://github.com/mdmfernandes/go-foscam/actions/workflows/tests.yml/badge.svg)](https://github.com/mdmfernandes/go-foscam/actions/workflows/tests.yml)
[![codecov](https://codecov.io/gh/mdmfernandes/go-foscam/graph/badge.svg?token=J9CVAXJ6JG)](https://codecov.io/gh/mdmfernandes/go-foscam)

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

$ make run/example
```

## To Do

- [ ] Add more functionalities
- [ ] Support more cameras
