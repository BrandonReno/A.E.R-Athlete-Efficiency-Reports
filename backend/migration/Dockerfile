FROM gomicro/goose

ADD migration/scripts/ /migrations/
ADD migration/entrypoint.sh /migrations/
RUN chmod +x /migrations/entrypoint.sh
