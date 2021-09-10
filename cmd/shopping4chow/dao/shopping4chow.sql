-- Database: shopping4chow

-- DROP DATABASE shopping4chow;
DROP TABLE IF EXISTS meal CASCADE;
DROP TABLE IF EXISTS ingredient CASCADE;
DROP TABLE IF EXISTS recipe CASCADE;
DROP TABLE IF EXISTS shoppinglist CASCADE;
DROP TABLE IF EXISTS meal_pics CASCADE;
DROP TYPE IF EXISts unit;
DROP TABLE IF EXISTS user_join_meal;
DROP TABLE IF EXISTS user_profile;

CREATE TABLE user_profile (
	user_name TEXT,
	email TEXT,
	PRIMARY KEY(user_name)
);

-- For Development Environment--
INSERT INTO user_profile(user_name,email) VALUES ('DEV','test@gmail.com');


CREATE TABLE meal (
	id SERIAL,
	name	TEXT NOT NULL,
	pic_id	INTEGER,
	description	TEXT,
	website	TEXT,
	PRIMARY KEY(id)
);

CREATE TABLE user_join_meal (
	user_name TEXT REFERENCES user_profile (user_name) ON DELETE CASCADE,
	meal_id int REFERENCES meal (id) ON DELETE CASCADE,
	CONSTRAINT id PRIMARY KEY (user_name,meal_id)
);

CREATE TABLE ingredient (
	id	SERIAL,
	name	VARCHAR(50) NOT NULL,
	s3key	VARCHAR(50),
	preferred_store VARCHAR(50),
	PRIMARY KEY(id)
);

CREATE TYPE unit AS ENUM ('none', 'each', 'peice', 'bag', 'bottle', 'box', 'case', 'pack','jar', 'can', 'bunch', 'roll',
	'dozen', 'small', 'large', 'lbs', 'qt', 'oz', 'cup', 'gallon', 'tbsp', 'tsp', 'g', 'kg','liter', 'milliliter','pis');

CREATE TABLE recipe (
	id	SERIAL,
	meal_id	INTEGER NOT NULL,
	name	TEXT NOT NULL,
	ingredient_id	INTEGER,
	amount	INTEGER NOT NULL,
	units	unit NOT NULL,
	PRIMARY KEY(id),
		CONSTRAINT fk_meal_id
		FOREIGN KEY(meal_id)
			REFERENCES meal(id),
	CONSTRAINT fk_ingredient_id
		FOREIGN KEY(ingredient_id)
			REFERENCES ingredient(id)
);

CREATE TABLE shoppinglist (
	id INTEGER,
	username VARCHAR(50),
	ingredient_id	INTEGER,
	PRIMARY KEY(id),
	CONSTRAINT fk_ingredient_id
		FOREIGN KEY(ingredient_id)
			REFERENCES ingredient(id)
);

CREATE TABLE meal_pics (
	id INTEGER,
	meal_id INTEGER NOT NULL,
	s3key VARCHAR(50) NOT NULL,
	PRIMARY KEY(id),
	CONSTRAINT fk_meal_id
		FOREIGN KEY(meal_id)
			REFERENCES meal(id)	
);
