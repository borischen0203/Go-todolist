<img src="https://raw.githubusercontent.com/scraly/gophers/main/chandleur-gopher.png" alt="chandleur-gopher">


# Go-todolist
This is a todo list service. You can get, create, update and delete todo item.

Frontend side fork from:
https://github.com/sdil/todo


## Features
- Get All todo items
- Create todo item
- Update todo item status
- Delete todo items

# How to use

## Run in Local:
Required
<!-- - Install go(version >= 1.7) -->
- Install docker
- Install `make` cli(https://formulae.brew.sh/formula/make)
```bash
brew install make
```

### Run steps
Step1: Clone the repo
```bash
git clone https://github.com/borischen0203/Go-todolist.git
```
Step2: Use `make` to execute makefile run docker-compose
```bash
make docker-up
```

If you don't have make command, use below command
```bash
docker-compose -f docker-compose.yml up --build
```

Step3:
Open browse:http://localhost:8080/

Demo:
[img]https://upload.cc/i1/2022/03/08/0emvHs.png[/img]


## Tech Stack
- Golang
- Gin framework
- Gorm
- RESTful API
- MySQL
<!-- - Swagger -->
- Docker-compose
- Github action(CI)
- Ngnix


## Todo:
- [ ] Add more status codes
- [ ] Improve services
- [ ] Improve handlers

