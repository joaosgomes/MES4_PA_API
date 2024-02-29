
# MES4 PA

Projeto no Âmbito de Programação Avançada para a Internet

O projeto é composto por:

- **Backend:** Golang + FiberGo
- **Frontend:** React.js (Create React App) + Server-Side Rendering (HTMX)
- **UI Frameworks:** Ant Design (React) + Tailwind CSS (HTMX)

Algumas Tecnologias Frontend:

- WebSockets
- Web Storage API
- Network Information API
- PWA
- Web Share API
- Vibration API
- Navigator Interface
- HTMX (hx-get)

## Stack

| Componente | Tecnologia        |
|------------|-------------------|
| Backend    | Golang + FiberGo  |
| Frontend   | React.js + HTMX   |
| UI         | Ant Design + Tailwind CSS |

## Arquitetura Geral

![](MES4_PA.svg)

## GitHub Repos

- (BE) [https://github.com/joaosgomes/MES4_PA_API](https://github.com/joaosgomes/MES4_PA_API)
- (FE) [https://github.com/joaosgomes/MES4_PA_FE](https://github.com/joaosgomes/MES4_PA_FE)
- (FE PROD DEVMODE = false) [https://github.com/joaosgomes/MES4_PA_FE](https://github.com/joaosgomes/MES4_PA_FE)

## Deploy FE (gh-pages)

- (BE Local) [https://github.com/joaosgomes/MES4_PA_API](https://github.com/joaosgomes/MES4_PA_API)
- (FE MockData) [https://joaosgomes.github.io/MES4_PA_FE/](https://joaosgomes.github.io/MES4_PA_FE/)
- (FE PROD DEVMODE = false) [https://joaosgomes.github.io/MES4_PA_FE_PROD](https://joaosgomes.github.io/MES4_PA_FE_PROD)

## API

![go](https://go.dev/images/go-logo-blue.svg)
![gofiber](https://gofiber.io/assets/images/logo.svg)

### Architecture

#### Project Structure

````console

.
├───components
│   ├───layouts
│   ├───pages
│   ├───ui
│   └───utils
├───config
├───constants
├───controllers
├───services
│   └───db
└───utils

````

#### Go Packeges Used

````console

go mod init api

go get github.com/gofiber/fiber/v2
go get github.com/gofiber/swagger
go get -u github.com/gofiber/template/html/v2
go get github.com/gofiber/contrib/websocket
--go get github.com/gofiber/websocket/v2
--go get github.com/gofiber/websocket
go get github.com/gofiber/contrib/websocket
go get github.com/google/uuid
````

### Swagger

Generate Documentation

#### Enpoint: <http://localhost:5000/swagger>

````console
swag init

````

### Commands

````console

go build main.go

go run main.go  //Check if port is not in use

//or

MAKE build

MAKE run



````

### Routes Backend

````console
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

![React](https://raw.githubusercontent.com/joaosgomes/MES4_PA_FE/main/public/logo512.png)

### Architecture

#### Project Structure

````console

├───components
│   ├───layouts
│   ├───pages
│   └───ui
├───config
├───constants
├───controllers
└───services
````

````console

npx create-react-app app --template cra-template-pwa

npm install axios

npm install antd --save

npm install gh-pages --save-dev

npm install moment

npm install uuid

````

### Commands

````
 npm start
 
 npm run deploy
 
````

### GitHub

````console
git init
git add .
git commit -m "."
git branch -M main
git remote add origin https://github.com/joaosgomes/MES4_PA_API.git
git push -u origin main

````

````console
git init
git add .
git commit -m "."
git branch -M main
git remote add origin https://github.com/joaosgomes/MES4_PA_FE.git
git push -u origin main



git init
git add .
git commit -m "MES4_PA_FE_PROD DEVMODE = false"
git branch -M main
git remote add origin https://github.com/joaosgomes/MES4_PA_FE_PROD.git
git push -u origin main

````

## References

<https://gofiber.io/>
<https://chat.openai.com/>
<https://stackoverflow.com/>
<https://react.dev/>
<https://htmx.org/>
