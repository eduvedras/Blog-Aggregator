#! /usr/bin/bash

source ../.env

cd ../sql/schema/
goose postgres $CONN_STRING down-to 0

goose postgres $CONN_STRING up

host=localhost:8080/v1

result=$(curl -X POST -d '{"name": "Jose"}' "$host/users")
apiKey=$(echo $result | jq .api_key)
result=$(curl -X POST -d "{\"name\": \"BBC News\", \"url\": \"http://feeds.bbci.co.uk/news/world/rss.xml\"}" "$host/feeds" -H "Authorization: ApiKey ${apiKey:1:(-1)}")
feedId=$(echo $result | jq .feed.id)
curl -X POST -d "{\"name\": \"CNN Top Stories\", \"url\": \"http://rss.cnn.com/rss/edition.rss\"}" "$host/feeds" -H "Authorization: ApiKey ${apiKey:1:(-1)}"

result=$(curl -X POST -d '{"name": "Mariana"}' "$host/users")
apiKey=$(echo $result | jq .api_key)
curl -X POST -d "{\"feed_id\": $feedId}" "$host/feed_follows" -H "Authorization: ApiKey ${apiKey:1:(-1)}"

result=$(curl -X POST -d '{"name": "Joao"}' "$host/users")
apiKey=$(echo $result | jq .api_key)
result=$(curl -X POST -d "{\"name\": \"Tech Crunch\", \"url\": \"http://feeds.feedburner.com/TechCrunch/\"}" "$host/feeds" -H "Authorization: ApiKey ${apiKey:1:(-1)}")
feedId=$(echo $result | jq .feed.id)

curl -X POST -d '{"name": "Maria"}' "$host/users"

result=$(curl -X POST -d '{"name": "Tiago"}' "$host/users")
apiKey=$(echo $result | jq .api_key)
curl -X POST -d "{\"feed_id\": $feedId}" "$host/feed_follows" -H "Authorization: ApiKey ${apiKey:1:(-1)}"
curl -X POST -d "{\"name\": \"Forbes\", \"url\": \"https://www.forbes.com/most-popular/feed/\"}" "$host/feeds" -H "Authorization: ApiKey ${apiKey:1:(-1)}"
