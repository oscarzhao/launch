# Launch

Launch is a manager on commandline apps, for example Elasticsearch, Spark, etc.

### Why I wrote this tool?

Recently I have been learning Collaborative Filtering implemented with Spark and ElasticSearch.
The [tutorial](https://github.com/IBM/elasticsearch-spark-recommender "tutorial spark es") requires a lot of preparations.
In my dev machine, I need to start ElasticSearch and Spark, and the Spark needs to be binded with jupyter notebook.

Here are a few problems I encountered (under windows):

1. Every time I start ElasticSearch, I need to `cd <somepath>/cf/elasticsearch-5.4.0/bin` directory, and run `elasticsearch.bat`
2. As to Spark, `pyspark` is `<somepath>/cf/spark-2.3.1-bin-hadoop2.7/bin`, but I need to run it under code directory `<code path>/elasticsearch-spark-recommender`
3. As to Spark, I need to configure four environment variables (to bind with jupyter lab, etc.)
4. As to Spark, it has some start args (need to include jar file to bind ES and Spark)

With `launch`, I can start with one command `launch pyspark` with the following setup:

```{json}
{
    "name": "pyspark",
    "binaryPath": "D:/Software/cf/spark-2.3.1-bin-hadoop2.7/bin/pyspark.cmd",
    "args": ["--driver-class-path", "../../elasticsearch-hadoop-5.4.0/dist/elasticsearch-spark-20_2.11-5.4.0.jar"],
    "workingDir": "D:/code/src/github.com/IBM/elasticsearch-spark-recommender",
    "env": [
        "SPARK_HOME=D:/Software/cf/spark-2.3.1-bin-hadoop2.7",
        "PYSPARK_DRIVER_PYTHON=jupyter",
        "PYSPARK_DRIVER_PYTHON_OPTS=lab",
        "PYSPARK_PYTHON=python"
    ]
}
```

## Install

Install binary into `$GOPATH/bin` with `go install github.com/oscarzhao/launch`

## Usage

Currently, `launch` only support one command (to start a command line process):

```launch <service name>```

When you run this command, `luancher` would read `~/.launch/config.json` and register all commands in the config.  Then the command with name `<service name>` would be run in the current shell window.

Sample config files can be found under directory `examples`.

## Error check

If you encountered an error while running a command, can check `~/.launch/launch.log` for detailed error information.