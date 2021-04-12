godiscover
==========

Tool for discovering go code.

Given a go package, godiscover tries to find "most popular" github repos using
that package. It uses godoc API for fetching package usages and then uses
github API for fetching stars of those repos.

# Pre-requisites

You can generate Github Access Token at https://github.com/settings/tokens.
You need to generate a "Personal Access Token".

# Build

Run `make build`

# Usage

`bin/godiscover <package_name> <github_access_token>`

Example usage:
```
bin/godiscover -pkg=go.mongodb.org/mongo-driver/mongo -token=<token>

2021/04/12 16:47:30 Fetching importer of pkg: go.mongodb.org/mongo-driver/mongo
2021/04/12 16:47:32 Fetched 379 github repos which import pkg go.mongodb.org/mongo-driver/mongo.
2021/04/12 16:47:34 Fetched stars of 100 repos.
2021/04/12 16:47:37 Fetched stars of 200 repos.
2021/04/12 16:47:40 Fetched stars of 300 repos.
2021/04/12 16:47:42 Repo: https://github.com/hashicorp/vault, Stars: 20504, Search query: https://github.com/search?type=code&q=go+mongodb+org+mongo-driver+mongo+in:file+repo:hashicorp/vault.
2021/04/12 16:47:42 Repo: https://github.com/chrislusf/seaweedfs, Stars: 11772, Search query: https://github.com/search?type=code&q=go+mongodb+org+mongo-driver+mongo+in:file+repo:chrislusf/seaweedfs.
2021/04/12 16:47:42 Repo: https://github.com/golang-migrate/migrate, Stars: 6220, Search query: https://github.com/search?type=code&q=go+mongodb+org+mongo-driver+mongo+in:file+repo:golang-migrate/migrate.
2021/04/12 16:47:42 Repo: https://github.com/mongodb/mongo-go-driver, Stars: 5602, Search query: https://github.com/search?type=code&q=go+mongodb+org+mongo-driver+mongo+in:file+repo:mongodb/mongo-go-driver.
2021/04/12 16:47:42 Repo: https://github.com/RichardKnop/machinery, Stars: 5093, Search query: https://github.com/search?type=code&q=go+mongodb+org+mongo-driver+mongo+in:file+repo:RichardKnop/machinery.
2021/04/12 16:47:42 Repo: https://github.com/gomods/athens, Stars: 3494, Search query: https://github.com/search?type=code&q=go+mongodb+org+mongo-driver+mongo+in:file+repo:gomods/athens.
2021/04/12 16:47:42 Repo: https://github.com/kedacore/keda, Stars: 3046, Search query: https://github.com/search?type=code&q=go+mongodb+org+mongo-driver+mongo+in:file+repo:kedacore/keda.
2021/04/12 16:47:42 Repo: https://github.com/spaceuptech/space-cloud, Stars: 2056, Search query: https://github.com/search?type=code&q=go+mongodb+org+mongo-driver+mongo+in:file+repo:spaceuptech/space-cloud.
2021/04/12 16:47:42 Repo: https://github.com/wal-g/wal-g, Stars: 1721, Search query: https://github.com/search?type=code&q=go+mongodb+org+mongo-driver+mongo+in:file+repo:wal-g/wal-g.
2021/04/12 16:47:42 Repo: https://github.com/mainflux/mainflux, Stars: 1366, Search query: https://github.com/search?type=code&q=go+mongodb+org+mongo-driver+mongo+in:file+repo:mainflux/mainflux.
2021/04/12 16:47:42 Repo: https://github.com/apache/servicecomb-service-center, Stars: 1224, Search query: https://github.com/search?type=code&q=go+mongodb+org+mongo-driver+mongo+in:file+repo:apache/servicecomb-service-center.
2021/04/12 16:47:42 Repo: https://github.com/zhenghaoz/gorse, Stars: 1176, Search query: https://github.com/search?type=code&q=go+mongodb+org+mongo-driver+mongo+in:file+repo:zhenghaoz/gorse.
2021/04/12 16:47:42 Repo: https://github.com/looplab/eventhorizon, Stars: 980, Search query: https://github.com/search?type=code&q=go+mongodb+org+mongo-driver+mongo+in:file+repo:looplab/eventhorizon.
...
```

You can use the URL in "Search query" for looking at code examples of package.
