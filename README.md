A CRUD application with `gorm.io/gorm`.

### References
- [CRUD REST API GOLANG GORM AND MYSQL
  ](https://golangexample.com/crud-rest-api-golang-gorm-and-mysql/)

Notes: Follow the steps in [Tutorial: Accessing a relational database - Set up a database](https://go.dev/doc/tutorial/database-access#set_up_database) to set up the database. Change the table name `album` to `albums` as `gorm.DB.AutoMigrate()` create table name in plurals. 
