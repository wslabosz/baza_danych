CREATE TABLE IF NOT EXISTS Passengers (
    passenger_id SERIAL PRIMARY KEY,
    firstname VARCHAR(255),
  	lastname VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(20),
    nationality VARCHAR(20),
  	date_of_birth DATE
);

CREATE TABLE IF NOT EXISTS Airports (
    airport_id SERIAL PRIMARY KEY,
    airport_code VARCHAR(10),
    airport_name VARCHAR(255),
  	airport_capacity INT,
    country VARCHAR(100),
    city VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS Flights (
    flight_id SERIAL PRIMARY KEY,
    flight_number VARCHAR(20),
    departure_airport_id INT,
    arrival_airport_id INT,
    departure_time TIMESTAMP,
    arrival_time TIMESTAMP,
  	flight_status VARCHAR(20),
  	flight_capacity INT,
    FOREIGN KEY (departure_airport_id) REFERENCES Airports (airport_id),
    FOREIGN KEY (arrival_airport_id) REFERENCES Airports (airport_id)
);

CREATE TABLE IF NOT EXISTS Seats (
    seat_id SERIAL PRIMARY KEY,
    seat_number VARCHAR(10),
    seat_status VARCHAR(20),
  	seat_class VARCHAR(10)
);

CREATE TABLE IF NOT EXISTS Bookings (
    booking_id SERIAL PRIMARY KEY,
    passenger_id INT,
  	flight_id INT,
    seat_id INT,
    booking_date TIMESTAMP,
    booking_status VARCHAR(20),
    FOREIGN KEY (passenger_id) REFERENCES Passengers (passenger_id),
  	FOREIGN KEY (flight_id) REFERENCES Flights (flight_id),
    FOREIGN KEY (seat_id) REFERENCES Seats (seat_id)
);

CREATE TABLE IF NOT EXISTS Payments (
    payment_id SERIAL PRIMARY KEY,
    booking_id INT,
    amount DECIMAL(10, 2),
    payment_date TIMESTAMP,
    payment_status VARCHAR(20),
    FOREIGN KEY (booking_id) REFERENCES Bookings (booking_id)
);

CREATE TABLE IF NOT EXISTS Baggage (
    baggage_id SERIAL PRIMARY KEY,
    passenger_id INT,
    booking_id INT,
  	baggage_type VARCHAR(20),
    weight DECIMAL(5, 2),
    FOREIGN KEY (passenger_id) REFERENCES Passengers (passenger_id),
    FOREIGN KEY (booking_id) REFERENCES Bookings (booking_id)
);