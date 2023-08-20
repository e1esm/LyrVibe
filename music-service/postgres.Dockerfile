FROM postgres:latest
ADD ./music-service/internal/repository/migrations/20230820204714_music_tables.up.sql /docker-entrypoint-initdb.d/
ENTRYPOINT ["docker-entrypoint.sh"]
EXPOSE 5432
CMD ["postgres"]