CREATE TABLE albums(
    id UUID PRIMARY KEY ,
    artist_id UUID,
    title TEXT,
    cover bytea,
    release_date DATE
);

CREATE TABLE tracks(
    id uuid PRIMARY KEY,
    artist_id uuid,
    album_id uuid,
    cover bytea,
    title varchar(50),
    release_date date,
    genre varchar(15),
    duration varchar(10),
    country varchar(30),
    video_link text,
    feature text[],
    created_at timestamp
);

CREATE TABLE track_statistics(
    id SERIAL PRIMARY KEY,
    track_id uuid REFERENCES tracks(id),
    views int default 0,
    rating float4 default 0
);

CREATE TABLE lyrics(
    id serial primary key,
    song_id uuid references tracks(id) ON DELETE CASCADE,
    line text,
    line_n int
);
