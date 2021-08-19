- Learn about Go Module; simply put, it is a dependency management tool for Go (just as Compose is for PHP, or NPM for node.js). 
  Read more here: 
  https://golang.org/doc/tutorial/create-module
  https://www.youtube.com/watch?v=AJxKKmzRTUY
  
- Any function/method in Go should return values only if that value is to be used. 
  Your database connection function should not be returning string like "db has been connected". How is that useful to the function caller? 
  
- DO NOT define global error variable. Each function should have error variable defined within and return if necessary (in most cases you would need to return error/nil). 

- It is recommended to treat ID as UNIQUE attribute and use it to search for records in DB tables. 

- Always handle sensitive info/config such as credentials in your application using .env file. So create one and pt the values below in it:
  - DB_NAME=Oluwashola
  - DB_HOST=localhost
  - DB_PORT=5432
  - DB_USERNAME=visasrecords
  - DB_PASSWORD=your-password
  - APP_PORT=:8080
  
You may need to source the .env before starting your application by running: source .env
You can confirm if the .env was properly 'sourced' by running:  echo $DB_NAME


- You can create a citizen table running :
  CREATE TABLE citizens
  (
  id serial PRIMARY KEY,
  name VARCHAR (50) UNIQUE NOT NULL,
  gender VARCHAR (50),
  country VARCHAR (50),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

More info/resource here: https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&ved=2ahUKEwj84p7N9rzyAhUD_RQKHT_GBgAQFnoECA0QAw&url=https%3A%2F%2Fwww.postgresqltutorial.com%2Fpostgresql-show-tables%2F&usg=AOvVaw3BDTL-U_m_7BP4Cwxn5Djc  
PS: Watch out for next blog to learn about database migration