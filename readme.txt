First run the docker-compose.yml file in order to setup a docker container for postgresql
docker-compose up --build

Then run go run main.go to inititate the go lang project

Postgresql port : 5432
Golang Project port : 3000


API CALLS:

GET ALL USERS: http://localhost:3000/GetAllUser
ADD A USER: http://localhost:3000/AddUser   -> Accepts Json
{
    "name":"Kritik",
    "email":"kritikmanral1@gmail.com",
    "course":"Btech",
    "Phone Number":8755130890
}
DELTE A USER: http://localhost:3000/DeleteUser/id
FIND USERS CORRESPONDING TO A PARTICULAR NAME: http://localhost:3000/FindUserByName/"name"

GET ALL BOOKS: http://localhost:3000/GetAllBooks
DELETE A BOOK: http://localhost:3000/DeleteBook/id
ADD A BOOK: http://localhost:3000/AddBook -> Accepts Json
{
  "Book Name":"DSA",
  "Course":"Btech",
  "Author":"XYZ",
  "Serial Number":1234
}
