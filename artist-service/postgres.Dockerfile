FROM postgres:latest
ADD ./artist-service/internal/repository/migrations/000001_artists_table.up.sql /docker-entrypoint-initdb.d/
ENTRYPOINT ["docker-entrypoint.sh"]
EXPOSE 5432
CMD ["postgres", "-c", "log_statement=all"]