docker-compose up -d

there are 3 folders:
- cmd/api
- cmd/processor
- cmd/reporting


due to a lack of the time I made `cmd/api`  following good structure and `cmd/processor` with `cmd/reporting` have not SOLID rules.

some code was copied between projects, however, it good to have shared libs to work with storage, queues and caching 

