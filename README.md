# Requirements
You need **Go environment**, **Node.js**, **npm**, **mongodb**. 
Below, the way to install above requirements by macports.  

```
% sudo port install go
% port installed go
The following ports are currently installed:
  go @1.22.0_0 (active)
```

```
% sudo port install npm10
% port installed npm10
The following ports are currently installed:
  npm10 @10.4.0_0 (active)
```

```
% sudo port install mongodb
% port installed mongodb 
The following ports are currently installed:
  mongodb @6.0.7_0 (active)
```

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
`% sudo mongod --dbpath=/opt/local/var/db/mongodb`  
`% cd ../backend`  
`% go run seeds/delete/delete.go -c users`  
`% go run seeds/delete/delete.go -c problems`  
`% go run seeds/delete/delete.go -c submission`  
`% go run seeds/insert/insert.go -c users -f users.json`  
`% go run seeds/insert/insert.go -c problems -f problems.json`  
`% go run seeds/insert/insert.go -c users -f users.json`  
`% rm -rf ../compile_resource/*`  

# Launch server
If you wanna launch FLJ, you run following commands as different processes.  
After that, you should go **http://localhost:5173**.  

`% cd frontend`  
`% npm run build`  
`% npm run preview`  
`% cd ../backend`  
`% go run main.go -release`  
