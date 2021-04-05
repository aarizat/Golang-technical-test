# Golang REST API

REST API to handler the register of employees.

## Installation

Before starting with configurations  in order to run this API, make sure to have `Golang` and `MySQL` in your machine, the versions used in development were `1.16` and `5.7` respectively.

 Now, follow these steps:
 
1. Clone this repository. Execute on your terminal:
```
$ git clone https://github.com/aarizat/Golang-technical-test
```
2. Oncce cloned the repository enter main folder and set up the database. The followings commands will create a Database called `cidenet_db`, a user `user_db` along with password `paswd_db`.
```
$ cd Golang-technical-test
$ cat setup_db.sql | mysql -uroot -p
```
> After executing the last command you must type your password.
3. Install `Go` dependencies to run this API.

Enter to `api` folder and after running `go mod tidy`, like so:
```
$ cd api
$ go mod tidy
```
4. Add this project to `GOROOTH`.
> Make sure that this project this in your `GOROOTH`

6. Inside api folder you can execute: 
```
$ HOST=localhost PORT=3306 USER=user_db DATABASE=cidenet_db PASSWORD=paswd_db go run main.go
```
Look at how is being passed the enviroment variables at the moment run the API.

The above command will execute a server in the port `5000` ready to receive requests.

## Usage
For this API were defined 5 endpoints [CRUD]:
```
* GET     /api/v1/employees             List all of employees
* POST    /api/v1/employees             Create a new employee
* GET     /api/v1/employees/{id_card}   Fetch an employee by his primary key (id_card)
* DELETE  /api/v1/employees/{id_card}   Delete an employee by his primary key (id_card)
* PUT     /api/v1/employees/{id_card}   Update an employee by his primary key (id_card)
```
In order to make a request to any these endpoint you could use a client like `POSTMAN`. Here an example:

#### POST request

<p align="center">
     <p align="center">
          <a href="https://imgur.com/fzIhppI"><img src="https://i.imgur.com/fzIhppI.png" title="source: imgur.com" /></a>
     </p>

#### Response from server

<p align="center">
     <p align="center">
          <a href="https://imgur.com/geAVYNM"><img src="https://i.imgur.com/geAVYNM.png" title="source: imgur.com" /></a>
     </p>

The same way you can make`GET`, `DELETE`, `PUT` request. Make sure pass id_card correct when you want to fetch, delte or update a specified employee.

## Author

<p align="center">
    <h2 align="center">By, Andres Ariza-Triana</h2>
      <p align="center">
        <a href="https://twitter.com/aarizatr" target="_blank">
            <img alt="twitter_page" src="https://raw.githubusercontent.com/EckoJuan/Readme_template/master/images/twitter.png" style="float: center; margin-right: 10px" height="50" width="50">
        </a>
        <a href="https://www.linkedin.com/in/aarizatr/" target="_blank">
            <img alt="linkedin_page" src="https://raw.githubusercontent.com/EckoJuan/Readme_template/master/images/linkedin.png" style="float: center; margin-right: 10px" height="50"  width="50">
        </a>
        <a href="https://medium.com/@aarizatr" target="_blank">
            <img alt="medium_page" src="https://raw.githubusercontent.com/EckoJuan/Readme_template/master/images/medium.png" style="float: center; margin-right: 10px" height="50" width="50">
        </a>
      </p>
</p>
