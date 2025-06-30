#!/bin/bash

set -euo pipefail

if [[ ! -f ./target/rolldice-0.0.1-SNAPSHOT.jar ]] ; then
    ./mvnw clean package -DskipTests
fi

# Download the OpenTelemetry Java agent, if we haven't done it already.
version=v2.13.0
jar=opentelemetry-javaagent.jar
if [[ ! -f ./${jar} ]] ; then
    curl -sL https://github.com/grafana/grafana-opentelemetry-java/releases/download/${version}/grafana-opentelemetry-java.jar -o ${jar}
fi

# ADD your OpenTelemetry Java agent environment variables here

java -jar ./target/rolldice-0.0.1-SNAPSHOT.jar  

