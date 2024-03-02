# Requirements
You need **Go environment**, **Node.js**, **npm**, **mongo shell**. 
Below, the way to install above requirements by macports.

`sudo port install go`  
`go version`  

`sudo port install npm10`  
`port installed npm10`  

`sudo port install mongodb`  
`port installed mongodb`  

# Used technologies
**React**, **Gin**, **MongoDB**  

# Getting Started
Firstly, you should clone this repo.  
`git clone git@github.com:LazyJudgements/FunalabJudge.git`  

Next, if you've never use **vite**, you should run following command.
`npm install vite`  

You should run the below bunch of commands.  
`npm install`  

# Launch server
If you wanna launch FLJ, you run following commands as different processes.  
After that, you should go **http://localhost:5173**.  

`cd frontend`
`npm run dev`  
`cd ../backend`
`go run main.go`  
`cd ..`
`sudo mongod --dbpath=/opt/local/var/db/mongodb`  
