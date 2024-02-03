# Go Web App Boilerplare
This is a boilderplate for building a web app with go. This can be used with things such as React to create an awesome and powerfull website.

## How to use
* Clone the repo
* Remove the .git folder
* Create a .env file with a port for the website and your postgres conn string
* go mod init <repo>
* go get <dependency>

## Dependencies
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgresql
go get github.com/gofiber/fiber/v2
go get github.com/gofiber/template/html/v2
go get github.com/joho/godotenv
