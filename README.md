# LyrVibe - General Information and Whitelist

LyrVibe is an API made for songwriters and music lovers which lets the first to share their song lyrics with the customers and provides the second to share their feelings about the song itself and about precise lyric line. It will help artists to summarize feedback better from the actual audience and let them know towards which direction to move in order to improve their further crafts.


## Gateway: 
- Routing between microservices - 1/4✅❌❌❌

## Auth Service:
- Sign up the service - ✅
- Log in the service - ✅
- Refresh token - ✅
- Store session in the Redis - ✅
- Store user's password, role etc - ✅


## Artist Service:
- Post song information to the service - ✅
    - Lyrics
    - Cover of the song
    - Featuring information
    - Basic info such as:
        - title
        - release year
        - genre
        - album connection(if any)
        - duration
        - country
        - link to the music video
        - Views on the platform
- Add artist profile to the service - ✅
- Summarize users's feedback on their songs based on the comments left - ❌
- Delete song from the service - ❌
- Update artist profile - ❌
- Post whole album - ❌

## Music_service
- Save all the songs in the database - ❌
- Return  songs from the database with multiple filters: -❌
    - By album
    - By artist
    - By genre
    - By country

## User_service:
- Get all songs of an exact artist (basic information): - ❌
    - Title
    - Genre
    - Artist
    - Album
    - Country
- Get all songs from the exact album (basic information): - ❌
    - Album image
    - Album title
    - Album release year
    - Songs list:
        - Title,
        - Co-artists (if any)
- Get information on an exact song: - ❌
    - Song title
    - Artist(s) name
    - Duration
    - Lyrics
    - Release year
    - Comments on each line
- Add comment on a line or lines. (AUTHORISED GUEST ONLY) - ❌
- Rate the song (1-10). (AUTHORISED GUEST ONLY) - ❌
- Get all the comments of the user (AUTHORISED GUEST ONLY) - ❌


## Admin role only:

### Statistics_Service:
- Generate statistics in the form of PDF on the content from the database - ❌
