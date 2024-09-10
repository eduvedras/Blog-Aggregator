# Blog-Aggregator

An API that allows users to aggregate all their favorite RSS blogs and feeds in one place. All you need to do is create an account and follow the feeds you like. If the feed you want does not exist in the database, you can create it by providing a name and a URL.

Periodically, a scraper will retrieve posts from each feed URL and add them to the database, making them accessible to users who follow the feed. A command-line interface for the API is available at [Blog-Aggregator-CLI](https://github.com/eduvedras/Blog-Aggregator-CLI) and a graphical user interface is under development, you can also make HTTP requests directly to the API endpoints.

## Motivation

This project was born out of my personal needs. I like to follow several blogs on topics ranging from technology to finance, but doing so required visiting each individual blog to check for new posts.

By using this tool, all the new posts from the feeds I follow are aggregated in one place, solving this problem. While other tools exist to address this issue, most are paid or lack certain features, so I decided to build one myself. Another motivation was to hone my programming skills by working on a project that involves several key concepts such as authentication, parallel programming, and more.

### Goal

The goal of this project is to provide a platform where users can aggregate posts from their favorite blogs. Specifically, the system interacts with four main tables via API endpoints:

* users: Stores user account information.
* feeds: Stores feed data.
* feed_follows: Represents the relationship between users and the feeds they follow. If a user unfollows a feed, this object is deleted.
* posts: Stores posts from all feeds, each with a feedId to identify the source feed.

## Endpoints
The API is hosted at:
*https://blog-aggregator-158858990102.europe-southwest1.run.app*

To interact with the API, prepend this URL to the endpoints. For example, to check the health of the service, use:
`https://blog-aggregator-158858990102.europe-southwest1.run.app/v1/healthz`.

### Check Server Health

**URL**: `/v1/healthz`

**Method**: `GET`

**Response**:
```json
{
  "status": "ok"
}
```

### Check if Error function is working

**URL**: `/v1/err`

**Method**: `GET`

**Response**:
```json
{
  "error": "Internal Server Error"
}
```

### Create user

**URL**: `/v1/users`

**Method**: `POST`

**Response**:
```json
{
  "id": "3f8805e3-634c-49dd-a347-ab36479f3f83",
  "created_at": "2022-09-01T00:00:00Z",
  "updated_at": "2022-09-01T00:00:00Z",
  "name": "Jose",
  "api_key": "5493b19da20c48a9bc1c260cecd85a61ebad5da74967d6066574f3ac28aa8c59"
}
```

### Get user

**URL**: `/v1/users`

**Method**: `GET`

**Headers**:
- `Authorization: ApiKey <key>`

**Response**:
```json
{
  "id": "3f8805e3-634c-49dd-a347-ab36479f3f83",
  "created_at": "2022-09-01T00:00:00Z",
  "updated_at": "2022-09-01T00:00:00Z",
  "name": "Jose",
  "api_key": "5493b19da20c48a9bc1c260cecd85a61ebad5da74967d6066574f3ac28aa8c59"
}
```

### Create feed

**URL**: `/v1/feeds`

**Method**: `POST`

**Headers**:
- `Authorization: ApiKey <key>`

**Request Body**:
```json
{
  "name": "BBC News",
  "url": "http://feeds.bbci.co.uk/news/world/rss.xml",
}
```

**Response**:
```json
{
  "feed": {
    "id": "4a82b372-b0e2-45e3-956a-b9b83358f86b",
    "created_at": "2021-05-01T00:00:00Z",
    "updated_at": "2021-05-01T00:00:00Z",
    "name": "BBC News",
    "url": "http://feeds.bbci.co.uk/news/world/rss.xml",
    "user_id": "d6962597-f316-4306-a929-fe8c8651671e"
  },
  "feed_follow": {
    "id": "c834c69e-ee26-4c63-a677-a977432f9cfa",
    "feed_id": "4a82b372-b0e2-45e3-956a-b9b83358f86b",
    "user_id": "d6962597-f316-4306-a929-fe8c8651671e",
    "created_at": "2021-05-01T00:00:00Z",
    "updated_at": "2021-05-01T00:00:00Z"
  } 
} 
```

### Get feeds

**Url**: `/v1/feeds`

**Method**: `GET`

**Query Parameters**:
- `offset` (integer, optional): Starting position to list the feeds. Defaults to `0`.
- `limit` (integer, optional): The number of feeds you want to list. Defaults to `20`.

**Response**:
```json
[
  {
    "id": "4a82b372-b0e2-45e3-956a-b9b83358f86b",
    "created_at": "2021-05-01T00:00:00Z",
    "updated_at": "2021-05-01T00:00:00Z",
    "name": "BBC News",
    "url": "http://feeds.bbci.co.uk/news/world/rss.xml",
    "user_id": "d6962597-f316-4306-a929-fe8c8651671e"
  },
  {
    "id": "db72557-b0e2-45e3-956a-b9b83358f86b",
    "created_at": "2021-06-01T00:00:00Z",
    "updated_at": "2021-06-01T00:00:00Z",
    "name": "CNN Top Stories",
    "url": "http://rss.cnn.com/rss/edition.rss",
    "user_id": "b4820sk9-f316-4306-a929-fe8c8651671e"
  }
]
```

### Create feed_follow

**URL**: `/v1/feed_follows`

**Method**: `POST`

**Headers**:
- `Authorization: ApiKey <key>`

**Request Body**:
```json
{
  "feed_id": "4a82b372-b0e2-45e3-956a-b9b83358f86b",
}
```

**Response**:
```json
{
  "id": "c834c69e-ee26-4c63-a677-a977432f9cfa",
  "feed_id": "4a82b372-b0e2-45e3-956a-b9b83358f86b",
  "user_id": "d6962597-f316-4306-a929-fe8c8651671e",
  "created_at": "2021-05-01T00:00:00Z",
  "updated_at": "2021-05-01T00:00:00Z"
} 
```

### Delete a feed_follow

**URL**: `/v1/feed_follows/{feedFollowId}`

**Method**: `DELETE`

**Headers**:
- `Authorization: ApiKey <key>`

**Response**:
- `200 OK`: The feed_follow was successfully deleted.
- `400 Bad Request`: Invalid feed_follow id.
- `500 Internal Server Error`: Something went wrong in the server while trying to delete the feed_follow.

### Get feed_follows of a user

**URL**: `/v1/feed_follows`

**Method**: `GET` 

**Headers**:
- `Authorization: ApiKey <key>`

**Response**:
```json
[
  {
    "id": "c834c69e-ee26-4c63-a677-a977432f9cfa",
    "feed_id": "4a82b372-b0e2-45e3-956a-b9b83358f86b",
    "user_id": "0e4fecc6-1354-47b8-8336-2077b307b20e",
    "created_at": "2018-01-01T00:00:00Z",
    "updated_at": "2018-01-01T00:00:00Z"
  },
  {
    "id": "ad752167-f509-4ff3-8425-7781090b5c8f",
    "feed_id": "f71b842d-9fd1-4bc0-9913-dd96ba33bb15",
    "user_id": "0e4fecc6-1354-47b8-8336-2077b307b20e",
    "created_at": "2018-02-01T00:00:00Z",
    "updated_at": "2018-02-01T00:00:00Z"
  }
]
```

### Get posts of feeds followed by user

**URL**: `/v1/posts`

**Method**: `GET`

**Query Parameters**:
- `limit` (integer, optional): The number of posts you want to list. Defaults to `10`.

**Headers**:
- `Authorization: ApiKey <key>`

**Response**:
```json
[
  {
    "id": "e20ff4f6-f5a7-4ae5-95a4-3d78b1b8c2df",
    "created_at": "2024-09-05 14:50:38.940043481+00:00",
    "updated_at": "2024-09-06 11:33:32.909614169+00:00",
    "title": "Michel Barnier named by Macron as new French PM",
    "url": "https://www.bbc.com/news/articles/cqjlxvg2gj7o",
    "description": "The French president names the EU's former Brexit negotiator almost two months after snap elections.",
    "published_at": "2022-03-13 15:04:00+00:00",
    "feed_id": "ea9ba4e4-2025-428c-b778-3d3dbc180510",
  },
  {
    "id": "f673cfde-12cc-4f31-a986-6feda5b0d2c3",
    "created_at": "2024-09-05 14:50:39.477619485+00:00",
    "updated_at": "2024-09-09 16:54:59.009103884+00:00",
    "title": "Shania Twain calls for equal pay and more diversity in country music",
    "url": "https://www.cnn.com/2023/04/03/entertainment/shania-twain-equal-diversity/index.html",
    "description": "Shania Twian is standing up for others in country music.",
    "published_at": "2022-01-31 10:56:00+00:00",
    "feed_id": "03f07bd0-eaf8-4366-8d1a-d9115fca0ad8",
  }
]
```

## Examples

### cURL Quickstart Example

```bash
$ host="https://blog-aggregator-158858990102.europe-southwest1.run.app"
$ curl -X POST -d '{"name": "Jose"}' "$host/v1/users"
{
  "id":"00e98306-2d72-4287-878b-3607820cd987",
  "created_at":"2024-09-05T15:14:05.340563676Z",
  "updated_at":"2024-09-05T15:14:05.340563798Z",
  "name":"Jose",
  "api_key":"18b21948f3ae444685442ce9901369d3698aff2cf45411c20a168e93bf5c0433"
}
$ apiKey="18b21948f3ae444685442ce9901369d3698aff2cf45411c20a168e93bf5c0433"
$ curl -X GET "$host/v1/users" -H "Authorization: ApiKey $apiKey"
{
  "id":"00e98306-2d72-4287-878b-3607820cd987",
  "created_at":"2024-09-05T15:14:05.340563676Z",
  "updated_at":"2024-09-05T15:14:05.340563798Z",
  "name":"Jose",
  "api_key":"18b21948f3ae444685442ce9901369d3698aff2cf45411c20a168e93bf5c0433"
}
$ curl -X POST -d "{\"name\": \"BBC News\", \"url\": \"http://feeds.bbci.co.uk/news/world/rss.xml\"}" "$host/v1/feeds" -H "Authorization: ApiKey $apiKey"
{
  "feed":{
    "id":"e148217f-174c-4ed2-bf21-de080bd204c6",
    "created_at":"2024-09-05T18:05:43.886873582Z",
    "updated_at":"2024-09-05T18:05:43.886873668Z",
    "name":"BBC News",
    "url":"http://feeds.bbci.co.uk/news/world/rss.xml",
    "user_id":"00e98306-2d72-4287-878b-3607820cd987",
    "last_fetched_at":null
  },
  "feed_follow":{
    "id":"0c71a2c6-c47c-4750-ba65-0e7f02d9ca85",
    "feed_id":"e148217f-174c-4ed2-bf21-de080bd204c6",
    "user_id":"00e98306-2d72-4287-878b-3607820cd987",
    "created_at":"2024-09-05T18:05:43.906955319Z",
    "updated_at":"2024-09-05T18:05:43.906955399Z"
  }
}
$ curl -X GET "$host/v1/feeds?offset=0&limit=2"
[
  {
    "id":"80614b4d-482c-4189-9a16-f62c9a2a0e69",
    "created_at":"2024-09-05T18:05:43.886873582Z",
    "updated_at":"2024-09-05T18:05:43.886873668Z",
    "name":"BBC News",
    "url":"http://feeds.bbci.co.uk/news/world/rss.xml",
    "user_id":"00e98306-2d72-4287-878b-3607820cd987",
    "last_fetched_at":"2024-09-05T18:19:30Z"
  },
  {
    "id":"ea9ba4e4-2025-428c-b778-3d3dbc180510",
    "created_at":"2024-09-05T14:45:48.25644086Z",
    "updated_at":"2024-09-05T14:45:48.25644176Z",
    "name":"CNN Top Stories",
    "url":"http://rss.cnn.com/rss/edition.rss",
    "user_id":"4dea0dbc-fb49-4358-ad08-93410cb94b37",
    "last_fetched_at":"2024-09-05T18:19:30Z"
  }
]
$ feedYouWantToFollow="ea9ba4e4-2025-428c-b778-3d3dbc180510"
$ curl -X POST -d "{\"feed_id\": \"$feedYouWantToFollow\"}" "$host/v1/feed_follows" -H "Authorization: ApiKey $apiKey"
{
  "id":"4f68f751-13b2-44db-8bf2-9477b6dfbe15",
  "feed_id":"ea9ba4e4-2025-428c-b778-3d3dbc180510",
  "user_id":"00e98306-2d72-4287-878b-3607820cd987",
  "created_at":"2024-09-05T19:30:27.373188578Z",
  "updated_at":"2024-09-05T19:30:27.373188751Z"
}
$ curl -X DELETE "$host/v1/feed_follows/4f68f751-13b2-44db-8bf2-9477b6dfbe15" -H "Authorization: ApiKey $apiKey"
{}
$ curl -X GET "$host/v1/feed_follows" -H "Authorization: ApiKey $apiKey"
[
  {
    "id":"0c71a2c6-c47c-4750-ba65-0e7f02d9ca85",
    "feed_id":"e148217f-174c-4ed2-bf21-de080bd204c6",
    "user_id":"00e98306-2d72-4287-878b-3607820cd987",
    "created_at":"2024-09-05T18:05:43.906955319Z",
    "updated_at":"2024-09-05T18:05:43.906955399Z"
  }
]
$ curl -X GET "$host/v1/posts?limit=1" -H "Authorization: ApiKey $apiKey"
[
  {
    "id":"01eb942e-2539-4872-8149-1ec90bfce770",
    "created_at":"2024-09-05T14:50:24.676224257Z",
    "updated_at":"2024-09-05T14:50:24.676224531Z",
    "title":"Russian authorities detain suspect over St. Petersburg cafe blast",
    "url":"https://edition.cnn.com/webview/europe/live-news/russia-ukraine-war-news-04-03-23/index.html",
    "description":"",
    "published_at":null,
    "feed_id":"ea9ba4e4-2025-428c-b778-3d3dbc180510"
  }
]
```
