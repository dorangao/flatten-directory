# Project: Directory Flattening with Docker

This project provides a set of Docker images to flatten directory structures for various programming languages. The process involves copying all files from a source directory to a target directory without preserving the folder structure.

## Prerequisites

- Docker installed on your system
- Source and target directories prepared

## Usage

### Bash

1. Navigate to the `bash` directory:
    ```sh
    cd bash
    ```

2. Build the Docker image:
    ```sh
    docker build -t bash-flatten .
    ```

3. Run the Docker container:
    ```sh
    docker run --rm -v "$(pwd)/source_dir":/source_dir -v "$(pwd)/target_dir":/target_dir bash-flatten /source_dir /target_dir
    ```

### TypeScript

1. Navigate to the `typescript` directory:
    ```sh
    cd ../typescript
    ```

2. Prepare test data:
    ```sh
    ../prepare_testdata.sh
    ```

3. Build the Docker image:
    ```sh
    docker build -t ts-flatten .
    ```

4. Run the Docker container:
    ```sh
    docker run --rm -v "$(pwd)/source_dir":/source_dir -v "$(pwd)/target_dir":/target_dir ts-flatten /source_dir /target_dir
    ```

### Python

1. Navigate to the `python` directory:
    ```sh
    cd ../python
    ```

2. Prepare test data:
    ```sh
    ../prepare_testdata.sh
    ```

3. Build the Docker image:
    ```sh
    docker build -t python-flatten .
    ```

4. Run the Docker container:
    ```sh
    docker run --rm -v "$(pwd)/source_dir":/source_dir -v "$(pwd)/target_dir":/target_dir python-flatten /source_dir /target_dir
    ```

### Java

1. Navigate to the `Java` directory:
    ```sh
    cd ../Java
    ```

2. Prepare test data:
    ```sh
    ../prepare_testdata.sh
    ```

3. Build the Docker image:
    ```sh
    docker build -t java-flatten .
    ```

4. Run the Docker container:
    ```sh
    docker run --rm -v "$(pwd)/source_dir":/source_dir -v "$(pwd)/target_dir":/target_dir java-flatten /source_dir /target_dir
    ```

### Go

1. Navigate to the `go` directory:
    ```sh
    cd ../go
    ```

2. Prepare test data:
    ```sh
    ../prepare_testdata.sh
    ```

3. Build the Docker image:
    ```sh
    docker build -t go-flatten .
    ```

4. Run the Docker container:
    ```sh
    docker run --rm -v "$(pwd)/source_dir":/source_dir -v "$(pwd)/target_dir":/target_dir go-flatten /source_dir /target_dir
    ```

### Rust

1. Navigate to the `rust` directory:
    ```sh
    cd ../rust
    ```

2. Prepare test data:
    ```sh
    ../prepare_testdata.sh
    ```

3. Build the Docker image:
    ```sh
    docker build -t rust-flatten .
    ```

4. Run the Docker container:
    ```sh
    docker run --rm -v "$(pwd)/source_dir":/source_dir -v "$(pwd)/target_dir":/target_dir rust-flatten /source_dir /target_dir
    ```

### PHP

1. Navigate to the `php` directory:
    ```sh
    cd ../php
    ```

2. Prepare test data:
    ```sh
    ../prepare_testdata.sh
    ```

3. Build the Docker image:
    ```sh
    docker build -t php-flatten .
    ```

4. Run the Docker container:
    ```sh
    docker run --rm -v "$(pwd)/source_dir":/source_dir -v "$(pwd)/target_dir":/target_dir php-flatten /source_dir /target_dir
    ```

### C\#

1. Navigate to the `cs` directory:
    ```sh
    cd ../cs
    ```

2. Prepare test data:
    ```sh
    ../prepare_testdata.sh
    ```

3. Build the Docker image:
    ```sh
    docker build -t csharp-flatten .
    ```

4. Run the Docker container:
    ```sh
    docker run --rm -v "$(pwd)/source_dir":/source_dir -v "$(pwd)/target_dir":/target_dir csharp-flatten /source_dir /target_dir
    ```

## Conclusion

This project demonstrates how to use Docker to flatten directory structures across multiple programming languages. Each language-specific directory contains a `Dockerfile` and necessary scripts to build and run the Docker containers.

## References

For more information, refer to [this article](https://medium.com/@dorangao/docker-to-verify-multi-language-implementations-a-journey-of-learning-and-applying-generative-ai-804e04efb48a).