FROM rust:1.52

WORKDIR /app

COPY flatten_directory.rs ./

RUN rustc flatten_directory.rs

ENTRYPOINT ["./flatten_directory"]
