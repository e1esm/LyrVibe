FROM postgres:latest
ADD ./internal/repository/userRepository/migrations/000001_users_table.up.sql /docker-entrypoint-initdb.d/
ENTRYPOINT ["docker-entrypoint.sh"]
EXPOSE 5432
CMD ["postgres"]
