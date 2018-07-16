# Launcher

Launcher is a manager on commandline apps, for example Elasticsearch, Spark, etc.

# Install

Install binary into `$GOPATH/bin` with `go get -u -v github.com/oscarzhao/launcher`

# Usage

Currently, `launcher` only support one command:

```launcher start <service name>```

When you run this command, `luancher` would read `~/.launcher/config.json` and register all commands in the config.  Then the command with name `<service name>` would be run in the current shell window.

Sample config files can be found under directory `examples`.

