# qmp-mac-generator

MAC addresses generator for QEMU VMs

This is a simple command line application generating unique MAC address suitable for QEMU instances.

Its main purpose is to use it to pre-generate MAC addresses before launching VMs.

One can use it in CI to ensure no VM would have same address across jobs and projects.

It uses a file lock to ensure that generated MAC address is unique across multiple VM instances.

The locking itself is based on flock sys call which is working fine even from inside docker containers.

By default file lock file is located at `/var/lock/qmp-mac-generator.lock`.

It contains the last generated MAC address.

One can define different lock file by setting `QMP_MAC_GENERATOR_LOCK_FILE` environment variable.

Generated MAC range is `52:54:00:AB:xx:xx` where x is 0x01-0xFF.

The generation starts over from beginning once it reaches the max address.

## Usage from shell

```shell
mac1=$(qmp-mac-generator)
mac2=$(qmp-mac-generator)
...
```

Optionally define different lock file location:

```shell
export QMP_MAC_GENERATOR_LOCK_FILE=/some-path/some-file

mac1=$(qmp-mac-generator)
mac2=$(qmp-mac-generator)
...
```

## Usage inside docker

Mount the lock file from host location, this way all docker instances will generate unique addresses.

```shell
docker ... \
  -v /var/lock/qmp-mac-generator.lock:/var/lock/qmp-mac-generator.lock \
  ...
```

If you leave the file container specific, mac addresses will be unique inside that container only.

## Usage 

```shell
go install github.com/fikin/qmp-mac-generator
...
mac=$(${GOBIN}/qmp-mac-generator)
```

## Building

```shell
go get .
go build -v ./...
go test
```
