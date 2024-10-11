FROM alpine:latest

COPY flatten_files.sh /flatten_files.sh
RUN chmod +x /flatten_files.sh

ENTRYPOINT ["sh", "/flatten_files.sh"]