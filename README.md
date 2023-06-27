# Ezra-Assessment
An Elevator API/Application (Ezra technical assessment).

## Requirements task A

Create an Elevator API/Application.
Endpoints have to:
1. Call the elevator from any floor to any other floor.
2. Get real time information about the elevator place, state, direction
Additionally, get log information about events and save them into the database.
Also save every SQL query which gets executed into the database with tracking
information on who/where/what made the call.
Requirements
● The building can have a configurable amount of floors and elevator moves 1 floor
per 5 seconds
● Doors open/close over 2 seconds
● Unit testing is also required.

Important requirement
Elevators must be able to move async - i.e. If I have 5 elevators all of which are moving
in separation directions, I should get records logged about every single action by each
elevator.
These logs must be segregated based on place/state/direction/etc, plus a way to see all
of them in real time.


## Requirements task B

Create an application that simulates a Lending/Repayment API.
The Developed API should have the following requirements met:
1. Receive a Lending request from a user (Loans Should be tied to a subscriber
MSISDN)
2. Receive repayment requests from the user should they top up their loans and update
the relevant tables
3. Add logic to sweep/clear old defaulted loans (The decision to clear the loans should
be configurable as that may vary from market to market. ie : Should the loan be cleared
after a loan age of 6 months?)
4. Simulate generation of dumps from the database to an SFTP server.
5. As you’re designing the API make decisions such as whether a user can have one or
multiple loans. Your database should cater for that.
6. Once a subscriber takes a loan or makes a repayment they should be notified by
SMS of the amount lent if it's a lending operation or the amount recovered if full or
partial.
7. All the endpoints and logic should be tested using a testing library of your choice.
8. Ensure that the endpoints are documented on an openly available

## Project overview
Both solutions are contained in the same workspace and can be run separately using Docker

## Tech-Stack
This project uses a variety of industry standard solutions to achieve the requirements.

 - [Go (Golang)](https://go.dev/)- All of the code + logic is written in Go
 - PostgreSQL- Database operations are done using PostgreSQL
- Github Actions- Automated testsare done using Github actions
- Docker-Virtualization, build and unit tests are done using Docker and Docker compose


