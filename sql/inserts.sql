-- pasazerowie
INSERT INTO Passengers (firstname, lastname, email, phone, nationality, date_of_birth)
VALUES ('Jacek', 'Wisniewski', 'jacek@example.com', '123456789', 'Poland', '1990-05-15');

INSERT INTO Passengers (firstname, lastname, email, phone, nationality, date_of_birth)
VALUES ('Placek', 'Roszkowski', 'placek@example.com', '999777888', 'Norway', '1988-12-10');

-- samoloty
INSERT INTO Airports (airport_code, airport_name, airport_capacity, country, city)
VALUES ('WCH', 'Chopen', 555, 'Poland', 'Warszawa');

INSERT INTO Airports (airport_code, airport_name, airport_capacity, country, city)
VALUES ('LUB', 'Lublinek', 222, 'Poland', 'Lodz');

-- loty
INSERT INTO Flights (flight_number, departure_airport_id, arrival_airport_id, departure_time, arrival_time, flight_status, flight_capacity)
VALUES ('FL001', 1, 2, '2023-06-10 09:00:00', '2023-06-10 12:00:00', 'on-time', 144);

INSERT INTO Flights (flight_number, departure_airport_id, arrival_airport_id, departure_time, arrival_time, flight_status, flight_capacity)
VALUES ('FL002', 2, 1, '2023-06-15 14:30:00', '2023-06-15 17:30:00', 'on-time', 93);

-- miejsca w samolocie
INSERT INTO Seats (seat_number, seat_status, seat_class)
VALUES ('A1', 'Available', 'Economy');

INSERT INTO Seats (seat_number, seat_status, seat_class)
VALUES ('B1', 'Available', 'Economy');

-- rezerwacje
INSERT INTO Bookings (passenger_id, flight_id, seat_id, booking_date, booking_status)
VALUES (1, 1, 1, '2023-06-01 10:00:00', 'Confirmed');

INSERT INTO Bookings (passenger_id, flight_id, seat_id, booking_date, booking_status)
VALUES (2, 2, 2, '2023-06-02 11:30:00', 'Confirmed');

-- platnosci
INSERT INTO Payments (booking_id, amount, payment_date, payment_status)
VALUES (1, 500.00, '2023-06-03 09:00:00', 'Unpaid');

INSERT INTO Payments (booking_id, amount, payment_date, payment_status)
VALUES (2, 750.00, '2023-06-04 10:30:00', 'Paid');

-- bagaze
INSERT INTO Baggage (passenger_id, booking_id, baggage_type, weight)
VALUES (1, 1, 'Checked', 20.5);

INSERT INTO Baggage (passenger_id, booking_id, baggage_type, weight)
VALUES (2, 2, 'Carry-on', 8.2);
