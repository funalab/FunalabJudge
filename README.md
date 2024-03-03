# Requirements
You need **Go environment**, **Node.js**, **npm**, **mongo shell**. 
Below, the way to install above requirements by macports.  

`% sudo port install go`  
`% go version`  

`% sudo port install npm10`  
`% port installed npm10`  

`% sudo port install mongodb`  
`% port installed mongodb`  

# Used technologies
**React**, **Gin**, **MongoDB**  

# Getting Started
Firstly, you should clone this repo.  
`% git clone git@github.com:LazyJudgements/FunalabJudge.git`  
`% cd FunalabJudge`  

You should run the below bunch of commands.  
`% cd frontend`  
`% npm install`  

You should add seed data into db by running following commands.  
`% cd ../backend`  
`% sudo mongod --dbpath=/opt/local/var/db/mongodb`  
`% go run seeds/delete/delete.go -c problems`  
`% go run seeds/delete/delete.go -c submission`  
`% go run seeds/insert/insert.go -c problems -f problems.json`  
`% rm -rf ../compile_resource/*`  

# Launch server
If you wanna launch FLJ, you run following commands as different processes.  
After that, you should go **http://localhost:5173**.  

`% cd ../frontend`  
`% npm run build`  
`% cd ../backend`  
`% go run main.go -release`  
`% cd ../frontend`  
`% npm run preview` (for check on local)  
