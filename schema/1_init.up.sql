CREATE TABLE reviews 
( 
    id SERIAL PRIMARY KEY, 
    user_id int not null, 
    product_id int not null, 
    text varchar(255) not null, 
    evaluation BIGINT not null 
); 

