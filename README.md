## Blackhole

A tool to dump elastic search.

## Building

```
  glide install
  glide up
  go build
```

## Usage

```
Tool of choice for elasticsearch dump

Usage:
  blackhole [flags]
  blackhole [command]

Available Commands:
  dump        Dump elasticsearch index to json
  export      Export elasticsearch json to index
  help        Help about any command
  version     Print version

Flags:
      --batchsize int   Size for each batch (default 20)
  -h, --help            help for blackhole

```

## License
MIT