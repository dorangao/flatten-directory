FROM openjdk:11-jdk-slim

WORKDIR /app

COPY FlattenDirectory.java ./

RUN javac FlattenDirectory.java

ENTRYPOINT ["java", "FlattenDirectory"]
