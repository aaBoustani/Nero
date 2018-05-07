# Nero
An open source team recognition tool built with Go and Badger

## Usage
In slack:
- `/chi-Nero @user x reason`: give x nero to @user.  
Notes:  
    - The default amount to give is 1 Nero.
    - The reason is optional.
    - Max amount to give in one day is 1
- `/chhal`: get the total number of Nero you have

## Integrate it to your Workspace
1. git clone git@github.com:AhmedBoustani/Nero.git
2. deploy the code to your favorite server
3. Add the app to your workspace.  
&nbsp;&nbsp; The endpoints are `/give` and  `/get-score`
4. Enjoy

## TODO
Split the code into the following packages:  
- router:
    - routes 
    - router
    - logger
- config:
    - db
    - cron
    - env
- api:
    - model (filename should be changed)
    - handlers (should be split into controller and helpers)
- services:
    - nero
