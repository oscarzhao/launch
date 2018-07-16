# Launch

Launch is a manager on commandline apps, for example Elasticsearch, Spark, etc.

## Install

Install binary into `$GOPATH/bin` with `go get -u -v github.com/oscarzhao/launch`

## Usage

Currently, `launch` only support one command (to start a command line process):

```launch <service name>```

When you run this command, `luancher` would read `~/.launch/config.json` and register all commands in the config.  Then the command with name `<service name>` would be run in the current shell window.

Sample config files can be found under directory `examples`.

## Error check

If you encountered an error while running a command, can check `~/.launch/launch.log` for detailed error information.