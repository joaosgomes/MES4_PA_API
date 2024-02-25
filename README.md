# MES4 PA



## API

````

go mod init api

go get github.com/gofiber/fiber/v2
go get github.com/gofiber/swagger
go get -u github.com/gofiber/template/html/v2

````

### Swagger

Generate Documentation

````
swag init

````

### Commands


````

go build main.go

go run main.go  //Check if port is not in use

````

### Routes

````
http://localhost:5000/
http://localhost:5000/live
http://localhost:5000/ready

````



## FE

````

npx create-react-app app --template cra-template-pwa

npm install axios

npm install antd --save

npm install gh-pages --save-dev
````


### Commands:


````
 npm start

````


### GitHub:

````

git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/joaosgomes/MES4_PA.git
git push -u origin main

````