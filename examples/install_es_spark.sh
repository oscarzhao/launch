#!/usr/bin/env bash

# run this script under certain folder

ES_VERSION=5.4.0
SPARK_VERSION=2.3.1

installPython36() {
    brew install pyenv
    echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.bash_profile
    echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.bash_profile
    echo -e 'if command -v pyenv 1>/dev/null 2>&1; then\n  eval "$(pyenv init -)"\nfi' >> ~/.bash_profile
    source ~/.bash_profile
    pyenv install 3.6.4
    pyenv local 3.6.4
}

installDep() {
    pip install elasticsearch numpy jupyterlab -q
    if [ -z "$(which wget)" ]; then brew install wget; fi
    if [ -z "$(which unzip)" ]; then brew install unzip; fi
}

downloadRepoAndData() {
    if [ ! -e elasticsearch-spark-recommender ]; then
        git clone https://github.com/IBM/elasticsearch-spark-recommender.git
        cd elasticsearch-spark-recommender/data
        wget http://files.grouplens.org/datasets/movielens/ml-latest-small.zip
        unzip ml-latest-small.zip
        rm ml-latest-small.zip
        cd ../..
    fi
    echo "Repo (including dataset) elasticsearch-spark-recommender is downloaded."
}

installES() {
    if [ ! -e elasticsearch-$ES_VERSION ]; then
        wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-$ES_VERSION.tar.gz
        tar xfz elasticsearch-$ES_VERSION.tar.gz
        cd elasticsearch-$ES_VERSION
        ./bin/elasticsearch-plugin install https://github.com/MLnick/elasticsearch-vector-scoring/releases/download/v$ES_VERSION/elasticsearch-vector-scoring-$ES_VERSION.zip
        cd ..
    fi
    echo "ElasticSearch is installed."
}

installESSparkConnector() {
    if [ ! -e elasticsearch-hadoop-$ES_VERSION ]; then
        wget http://download.elastic.co/hadoop/elasticsearch-hadoop-$ES_VERSION.zip
        unzip elasticsearch-hadoop-$ES_VERSION.zip
    fi
    echo "ElasticSearch Spark Connector is installed."
}

installSpark() {
    if [ -e spark-$SPARK_VERSION-bin-hadoop2.7.tgz ]; then
        tar zxvf spark-$SPARK_VERSION-bin-hadoop2.7.tgz
    elif [ ! -e spark-$SPARK_VERSION-bin-hadoop2.7 ]; then
        wget http://www-us.apache.org/dist/spark/spark-$SPARK_VERSION/spark-$SPARK_VERSION-bin-hadoop2.7.tgz
        tar zxvf spark-$SPARK_VERSION-bin-hadoop2.7.tgz
    fi
    echo "Spark is installed."
}

installLauncher() {
    go get -u github.com/oscarzhao/launch
    if [ ! -e ~/.launch ]; then mkdir ~/.launch; fi
    cat > ~/.launch/config.json <<EOF
{
    "log": {
        "level": "info"
    },
    "commands": [
        {
            "name": "es5",
            "binaryPath": "`pwd`/elasticsearch-$ES_VERSION/bin/elasticsearch",
            "workingDir": "`pwd`/elasticsearch-$ES_VERSION"
        },
        {
            "name": "pyspark",
            "binaryPath": "`pwd`/spark-$SPARK_VERSION-bin-hadoop2.7/bin/pyspark",
            "args": ["--driver-class-path", "../../elasticsearch-hadoop-$ES_VERSION/dist/elasticsearch-spark-20_2.11-$ES_VERSION.jar"],
            "workingDir": "`pwd`/elasticsearch-spark-recommender",
            "env": [
                "SPARK_HOME=`pwd`/spark-$SPARK_VERSION-bin-hadoop2.7",
                "PYSPARK_DRIVER_PYTHON=jupyter",
                "PYSPARK_DRIVER_PYTHON_OPTS=lab",
                "PYSPARK_PYTHON=python"
            ]
        }
    ]
}
EOF

    cat <<EOF
Configure successfully.
Start ElasticSearch: launch es5
Start Spark:         launch pyspark
EOF
}

installPython36
installDep
downloadRepoAndData
installES
installESSparkConnector
installSpark
installLauncher
