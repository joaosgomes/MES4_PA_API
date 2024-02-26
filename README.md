# MES4 PA

### GitHub Repos:

(BE) https://github.com/joaosgomes/MES4_PA_API

(FE) https://github.com/joaosgomes/MES4_PA_FE

## API



### Architecture


#### Project Structure

````

├───config
├───controllers
├───db
├───docs
├───models
├───public
│   ├───scripts
│   └───styles
├───repository
├───services
├───utils
└───views
````

````

go mod init api

go get github.com/gofiber/fiber/v2
go get github.com/gofiber/swagger
go get -u github.com/gofiber/template/html/v2
go get github.com/gofiber/contrib/websocket
--go get github.com/gofiber/websocket/v2
--go get github.com/gofiber/websocket
go get github.com/gofiber/contrib/websocket
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
http://localhost:5000/swagger
http://localhost:5000/live
http://localhost:5000/ready
http://localhost:5000/metrics
http://localhost:5000/event/:id, GetEvent()
http://localhost:5000/event", GetEvents()
http://localhost:5000/event", PostEvent()
http://localhost:5000/event/:id, PutEvent()
http://localhost:5000/event/:id, DeleteEvent()
http://localhost:5000/ws/:id, WSHandler()
http://localhost:5000/image/:id, .GetImage()
http://localhost:5000/image, .GetImages()
http://localhost:5000/image, .PostImage()
http://localhost:5000/image/:id, .DeleteImage()
````



## FE


### Architecture


#### Project Structure

````

├───components
│   ├───layouts
│   ├───pages
│   └───ui
├───config
├───constants
├───controllers
└───services
````

````

npx create-react-app app --template cra-template-pwa

npm install axios

npm install antd --save

npm install gh-pages --save-dev
````


### Commands:


````
 npm start
 
 npm run deploy
 

````



### Deploy FE (gh-pages)

https://joaosgomes.github.io/MES4_PA_FE/


### GitHub:

````

git init
git add .
git commit -m "."
git branch -M main
git remote add origin https://github.com/joaosgomes/MES4_PA_API.git
git push -u origin main

````


````
git init
git add .
git commit -m "."
git branch -M main
git remote add origin https://github.com/joaosgomes/MES4_PA_FE.git
git push -u origin main

````



## References:

https://gofiber.io/
https://chat.openai.com/
https://stackoverflow.com/
https://react.dev/
https://htmx.org/