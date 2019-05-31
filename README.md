# Journal fullstack exercise

## Description

In this exercise, you'll be filling in the missing piece in a simple 
client/backend/database system. 

The environment is orchestrated using docker compose, which currently has the client and
database configurations built in. 

The client is a simple React app that is expecting to get a list of people from the
backend, with a optional search query parameter.

The database is a MySQL database with 2 tables, `people` and `people_colors`. Here's the 
schemas for those tables:

### people
| **Field** | **Type**    |
| --------- | ----------- |
| id        | Primary Key |
| name      | Varchar     |
| img_url . | Varchar .   |
| location .| Varchar .   |

### people_colors
| **Field** | **Type**    |
| --------- | ----------- |
| id        | Primary Key |
| people_id | Foreign Key |
| color .   | Varchar .   |

The backend should respond to `http://localhost:4000/people` with the optional query parameter `q` 
(eg. `http://localhost:4000/people?q=vin`. 

With no query, this endpoint should return all people in the table. When specified, the query
parameter should be used to search people.

This endpoint should return a json response, with a list of records of the following form:

```json
[
  {
    "colors": [
      "Goldenrod", 
      "Yellow", 
      "Yellow"
    ], 
    "id": 2, 
    "imgUrl": "https://robohash.org/utquamreiciendis.jpg?size=64x64&set=set1", 
    "location": "Iguig", 
    "name": "Rourke Luckin"
  }, 
  {
    "colors": [
      "Yellow", 
      "Puce", 
      "Indigo"
    ], 
    "id": 3, 
    "imgUrl": "https://robohash.org/omnisbeataeest.jpg?size=64x64&set=set1", 
    "location": "Formiga", 
    "name": "Brianna McElree"
  }, 
  ...
]
```

Feel free to create the backend using any language of your choice. The only requirement
is to dockerize the backend, and add it to the docker compose environment. 

## Getting started

To get started, please fork this repository. We'd like you to share this repository
with us once you're done.

To run this environment, the only prerequisite is Docker. Follow the instructions 
[here](https://docs.docker.com/install/) to install docker for your platform.

After that, run the following command in project directory:

```
docker compose up
```

Without an implemented backend, the client and database should launch after a little bit.

If you visit [http://localhost:3000](http://localhost:3000), you should see the following page:

## Expected results

When you've successfully created a backend that's running at port 4000, exposed in your docker
environment, you should see the following list of cards in the client:


And here's what results you should get if you're searching:
