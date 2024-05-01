### Dependencies
1. Run docker container with postgresql
```sh
docker run -d -p5432:5432 --name some-postgres -e POSTGRES_PASSWORD=password -e PGDATA=/var/lib/postgresql/data/pgdata -v postgresql:/var/lib/postgresql/data postgres:15.3
```
The database config in `local_db_env`  
2. **You need to create cinema database by yourself** \
3. Initialize database using `cinemav2.sql` file \
4. Download and run app at the release of the repository (**You need to download `local_db_env` and it need to be in the same directory with the app**) \
5. The app will be available at `localhost:8080` \
6. The api documentation will be available at `http://localhost:8080/swagger/index.html` 


```
cinema-golang
├─ backend-go
│  ├─ Dockerfile
│  ├─ Makefile
│  ├─ common
│  │  ├─ app_err.go
│  │  ├─ app_response.go
│  │  ├─ const.go
│  │  ├─ paging.go
│  │  ├─ random.go
│  │  ├─ simple_cinema.go
│  │  ├─ simple_user.go
│  │  ├─ sql_model.go
│  │  ├─ time.go
│  │  ├─ transaction.go
│  │  └─ uid.go
│  ├─ component
│  │  ├─ appctx
│  │  │  └─ appctx.go
│  │  ├─ hasher
│  │  │  └─ md5.go
│  │  ├─ tokenprovider
│  │  │  ├─ jwt
│  │  │  │  └─ jwt.go
│  │  │  ├─ provider.go
│  │  │  └─ token.go
│  │  └─ transactor
│  │     ├─ example
│  │     │  └─ gorm_transactor.go
│  │     └─ transactor.go
│  ├─ docs
│  │  ├─ docs.go
│  │  ├─ swagger.json
│  │  └─ swagger.yaml
│  ├─ go.mod
│  ├─ go.sum
│  ├─ img.png
│  ├─ local_db_env
│  ├─ main.go
│  ├─ middleware
│  │  ├─ authorize.go
│  │  ├─ check_role.go
│  │  └─ recover.go
│  ├─ module
│  │  ├─ auditorium
│  │  │  ├─ biz
│  │  │  │  ├─ create.go
│  │  │  │  ├─ find.go
│  │  │  │  └─ list.go
│  │  │  ├─ model
│  │  │  │  ├─ auditorium.go
│  │  │  │  └─ filter.go
│  │  │  ├─ repository
│  │  │  │  ├─ create.go
│  │  │  │  └─ list_auditoriums_with_cinema_name.go
│  │  │  ├─ store
│  │  │  │  ├─ create.go
│  │  │  │  ├─ find.go
│  │  │  │  ├─ list.go
│  │  │  │  ├─ store.go
│  │  │  │  └─ transaction.go
│  │  │  └─ transport
│  │  │     └─ gin
│  │  │        ├─ create.go
│  │  │        └─ list_auditorium_with_cinema.go
│  │  ├─ auditorium_seat
│  │  │  ├─ model
│  │  │  │  ├─ filter.go
│  │  │  │  └─ seat.go
│  │  │  └─ store
│  │  │     ├─ create.go
│  │  │     ├─ list.go
│  │  │     ├─ store.go
│  │  │     └─ transaction.go
│  │  ├─ cinema
│  │  │  ├─ biz
│  │  │  │  ├─ create_cinema.go
│  │  │  │  ├─ get_cinema.go
│  │  │  │  └─ list_cinema.go
│  │  │  ├─ model
│  │  │  │  ├─ cinema.go
│  │  │  │  └─ filter.go
│  │  │  ├─ store
│  │  │  │  ├─ create.go
│  │  │  │  ├─ find.go
│  │  │  │  ├─ list.go
│  │  │  │  ├─ store.go
│  │  │  │  └─ update.go
│  │  │  └─ transport
│  │  │     ├─ endpoint
│  │  │     │  ├─ endpoints.go
│  │  │     │  ├─ request.go
│  │  │     │  └─ response.go
│  │  │     └─ gin
│  │  │        ├─ create_cinema.go
│  │  │        ├─ get_cinema_from_id.go
│  │  │        ├─ get_cinema_from_name.go
│  │  │        └─ list_cinema.go
│  │  ├─ company
│  │  │  ├─ biz
│  │  │  │  ├─ create.go
│  │  │  │  ├─ get_company.go
│  │  │  │  └─ list.go
│  │  │  ├─ model
│  │  │  │  ├─ company.go
│  │  │  │  └─ filter.go
│  │  │  ├─ store
│  │  │  │  ├─ create.go
│  │  │  │  ├─ find.go
│  │  │  │  ├─ list.go
│  │  │  │  └─ store.go
│  │  │  └─ transport
│  │  │     └─ gin
│  │  │        ├─ create.go
│  │  │        ├─ get_company_from_id.go
│  │  │        └─ list.go
│  │  ├─ country
│  │  │  ├─ biz
│  │  │  │  ├─ create.go
│  │  │  │  └─ list.go
│  │  │  ├─ model
│  │  │  │  └─ country.go
│  │  │  ├─ store
│  │  │  │  ├─ create.go
│  │  │  │  ├─ list.go
│  │  │  │  └─ store.go
│  │  │  └─ transport
│  │  │     └─ gin
│  │  │        ├─ create.go
│  │  │        └─ list.go
│  │  ├─ movie
│  │  │  ├─ biz
│  │  │  │  ├─ create.go
│  │  │  │  ├─ find_company.go
│  │  │  │  ├─ get.go
│  │  │  │  └─ list.go
│  │  │  ├─ model
│  │  │  │  ├─ filter.go
│  │  │  │  └─ movie.go
│  │  │  ├─ store
│  │  │  │  ├─ create.go
│  │  │  │  ├─ find.go
│  │  │  │  ├─ get.go
│  │  │  │  ├─ list.go
│  │  │  │  └─ store.go
│  │  │  └─ transport
│  │  │     └─ gin
│  │  │        ├─ create.go
│  │  │        ├─ get_movie_from_id.go
│  │  │        └─ list.go
│  │  ├─ show
│  │  │  ├─ biz
│  │  │  │  ├─ create.go
│  │  │  │  ├─ find_show.go
│  │  │  │  ├─ get.go
│  │  │  │  ├─ list.go
│  │  │  │  └─ update.go
│  │  │  ├─ model
│  │  │  │  ├─ filter.go
│  │  │  │  └─ show.go
│  │  │  ├─ repository
│  │  │  │  └─ create.go
│  │  │  ├─ store
│  │  │  │  ├─ create.go
│  │  │  │  ├─ find.go
│  │  │  │  ├─ get.go
│  │  │  │  ├─ list.go
│  │  │  │  ├─ store.go
│  │  │  │  ├─ transaction.go
│  │  │  │  └─ update.go
│  │  │  └─ transport
│  │  │     └─ gin
│  │  │        ├─ create.go
│  │  │        ├─ get_show_from_id.go
│  │  │        └─ list.go
│  │  ├─ ticket
│  │  │  ├─ biz
│  │  │  │  ├─ create.go
│  │  │  │  ├─ find_ticket.go
│  │  │  │  ├─ get.go
│  │  │  │  ├─ list.go
│  │  │  │  └─ update.go
│  │  │  ├─ model
│  │  │  │  ├─ filter.go
│  │  │  │  └─ ticket.go
│  │  │  ├─ store
│  │  │  │  ├─ create.go
│  │  │  │  ├─ find.go
│  │  │  │  ├─ get.go
│  │  │  │  ├─ list.go
│  │  │  │  ├─ store.go
│  │  │  │  ├─ transaction.go
│  │  │  │  └─ update.go
│  │  │  └─ transport
│  │  │     └─ gin
│  │  │        ├─ create.go
│  │  │        ├─ get_ticket.go
│  │  │        ├─ list.go
│  │  │        └─ update.go
│  │  └─ user
│  │     ├─ business
│  │     │  ├─ login.go
│  │     │  └─ register.go
│  │     ├─ transport
│  │     │  └─ ginuser
│  │     │     ├─ get_profile.go
│  │     │     ├─ login.go
│  │     │     └─ register.go
│  │     ├─ usermodel
│  │     │  ├─ user.go
│  │     │  └─ user_create.go
│  │     └─ userstore
│  │        ├─ create.go
│  │        ├─ find.go
│  │        └─ store.go
```
