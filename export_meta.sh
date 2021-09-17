curl 'http://localhost:8888/v1/metadata' -X POST --data-raw '{"type":"export_metadata","version":2,"args":{}}' --output './hasura-meta.json'
git add ./hasura-meta.json