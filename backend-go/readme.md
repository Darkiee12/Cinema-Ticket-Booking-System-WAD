### Dependencies
1. Run docker container with postgresql
```sh
docker run -d -p5432:5432 --name some-postgres -e POSTGRES_PASSWORD=password -e PGDATA=/var/lib/postgresql/data/pgdata -v postgresql:/var/lib/postgresql/data postgres:15.3
```
The database config in `local_db_env` 
2. **You need to create cinema database by yourself** 
3. Initialize database using `cinemav2.sql` file 
4. Download and run app at the release of the repository (**You need to download `local_db_env` and it need to be in the same directory with the app**) 
5. The app will be available at `localhost:8080` 
6. The api documentation will be available at `http://localhost:8080/swagger/index.html` 

