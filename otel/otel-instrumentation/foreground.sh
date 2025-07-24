echo -e "4 steps setup:\n - Step 1: LGTM + Alloy stack\n - Step 2: Java app\n - Step 3: OpenTelemetry agent\n - Step 4: Pull k6" && \
echo -e "\n\e[1m>> Step 1: Setting up LGTM + Alloy stack...\e[0m" && \
docker-compose -f /root/course/docker-compose.yaml up -d && \
echo -e "\e[92mLGTM ready. You can start the workshop while the terminal finishes up.\e[39m" && \
echo -e "\n\e[1m>> Step 2: Setting up JDK & Java app...\e[0m" && \
apt-get update -qq && \
apt-get install -y openjdk-17-jdk openjdk-17-jre && \
echo -e "\e[92mJDK installed. Building Java app...\e[0m" && \
cd /root/course/rolldice && \
chmod +x ./mvnw && \
chmod +x ./run.sh && \
./mvnw clean package -DskipTests && \
echo -e "\e[92mJava app built successfully\e[0m" && \
echo -e "\n\e[1m>> Step 3: Downloading OpenTelemetry agent...\e[0m" && \
version=v2.13.0 && \
jar=opentelemetry-javaagent.jar && \
curl -sL https://github.com/grafana/grafana-opentelemetry-java/releases/download/${version}/grafana-opentelemetry-java.jar -o ${jar} && \
echo -e "\e[92mOpenTelemetry agent downloaded\e[0m" && \
echo -e "\n\e[1m>> Step 4: Pulling k6 image for load testing...\e[0m" && \
docker pull grafana/k6:latest && \
echo -e "\e[92mSetup completed\e[0m"
