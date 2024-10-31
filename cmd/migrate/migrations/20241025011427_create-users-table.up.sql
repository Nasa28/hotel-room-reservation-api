CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    firstName VARCHAR(255) NOT NULL,
    lastName VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    phoneNumber VARCHAR(255) ,
    streetName VARCHAR(255),      
    streetNumber VARCHAR(255),     
    city VARCHAR(255),             
    state VARCHAR(255),            
    country VARCHAR(255),          
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
