/* 1. Dodaj pasazera (klienta) */
CREATE OR REPLACE FUNCTION add_passenger(
    p_firstname VARCHAR(255),
  	p_lastname VARCHAR(255),
    p_email VARCHAR(255),
    p_phone VARCHAR(20),
    p_nationality VARCHAR(20),
  	p_date_of_birth DATE
)
RETURNS INTEGER AS $$
DECLARE
    v_passenger_id INTEGER;
BEGIN
    INSERT INTO Passengers (firstname, lastname, email, phone, nationality, date_of_birth)
    VALUES (p_firstname, p_lastname, p_email, p_phone, p_nationality, p_date_of_birth)
    RETURNING passenger_id INTO v_passenger_id;
    
    RETURN v_passenger_id;
END;
$$ LANGUAGE plpgsql;

/* 2.Zaaktualizuj status lotu */
CREATE OR REPLACE FUNCTION update_flight_status(
    p_flight_id INT,
    p_flight_status VARCHAR(20)
)
RETURNS VOID AS $$
BEGIN
    UPDATE Flights
    SET flight_status = p_flight_status
    WHERE Flights.flight_id = p_flight_id;
END;
$$ LANGUAGE plpgsql;

/* 3. Zwroc informacje o rezerwacji */
CREATE OR REPLACE FUNCTION get_booking_details(
    p_booking_id INT
)
RETURNS TABLE (
    booking_id INT,
    passenger_id INT,
    flight_id INT,
  	seat_number VARCHAR(10),
  	seat_class VARCHAR(10),
    booking_date TIMESTAMP,
    booking_status VARCHAR(20),
    passenger_firstname VARCHAR(255),
    passenger_lastname VARCHAR(255),
    flight_number VARCHAR(20),
    departure_time TIMESTAMP,
    arrival_time TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT b.booking_id, b.passenger_id, b.flight_id, s.seat_number, s.seat_class, 
   		b.booking_date, b.booking_status,
        p.firstname AS passenger_firstname, p.lastname AS passenger_lastname,
        f.flight_number, f.departure_time, f.arrival_time
    FROM Bookings b
    JOIN Passengers p ON b.passenger_id = p.passenger_id
    JOIN Seats s on b.seat_id = s.seat_id
    JOIN Flights f ON b.flight_id = f.flight_id
    WHERE b.booking_id = p_booking_id;
END;
$$ LANGUAGE plpgsql;

/* 4. Znajdz polaczenie danego dnia miedzy danymi lotniskami */
CREATE OR REPLACE FUNCTION search_flights(
    p_departure_airport_name VARCHAR(255),
    p_arrival_airport_name VARCHAR(255),
    p_departure_date DATE
) 
RETURNS TABLE (
    flight_id INT,
    flight_number VARCHAR(20),
    departure_airport_id INT,
    arrival_airport_id INT,
    departure_time TIMESTAMP,
    arrival_time TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT f.flight_id, f.flight_number, f.departure_airport_id, f.arrival_airport_id, f.departure_time, f.arrival_time
    FROM Flights f
    INNER JOIN Airports dep ON f.departure_airport_id = dep.airport_id
    INNER JOIN Airports arr ON f.arrival_airport_id = arr.airport_id
    WHERE dep.airport_name = p_departure_airport_name
        AND arr.airport_name = p_arrival_airport_name
        AND f.departure_time::DATE = p_departure_date
        AND f.flight_id IN (
            SELECT fe.flight_id
            FROM Flights fe
            GROUP BY fe.flight_id
            HAVING SUM(CASE WHEN fe.flight_status = 'cancelled' THEN 1 ELSE 0 END) = 0
        );

    RETURN;
END;
$$ LANGUAGE plpgsql;

/* 5. zwroc przychod */
CREATE OR REPLACE FUNCTION get_total_revenue()
RETURNS DECIMAL AS $$
DECLARE
    total_revenue DECIMAL(10, 2);
BEGIN
    SELECT SUM(p.amount) INTO total_revenue
    FROM Payments p;
    
    RETURN total_revenue;
END;
$$ LANGUAGE plpgsql;

/* 6. anuluj loty z danego lotniska */
CREATE OR REPLACE FUNCTION cancel_flights_by_airport(p_airport_id INT, p_cancellation_date DATE)
RETURNS TABLE (affected_flights INT, affected_bookings INT) AS $$
BEGIN
    UPDATE Flights
    SET flight_status = 'cancelled'
    WHERE departure_airport_id = p_airport_id
    AND departure_time::DATE = p_cancellation_date
    AND flight_status <> 'cancelled';
    
    GET DIAGNOSTICS affected_flights = ROW_COUNT;

    UPDATE Bookings
    SET booking_status = 'cancelled'
    WHERE flight_id IN (
        SELECT flight_id
        FROM Flights
        WHERE departure_airport_id = p_airport_id
        AND flight_status = 'cancelled'
    )
    AND booking_status <> 'cancelled';

    GET DIAGNOSTICS affected_bookings = ROW_COUNT;
    
    CREATE TEMPORARY TABLE temp_counts (affected_f INT, affected_b INT);
    
    INSERT INTO temp_counts (affected_f, affected_b)
    VALUES (affected_flights, affected_bookings);
    
    RETURN QUERY SELECT tmp.affected_f, tmp.affected_b FROM temp_counts as tmp;

EXCEPTION
    WHEN OTHERS THEN
        RAISE;
END;
$$ LANGUAGE plpgsql;

/* 7. wypisz wszystkie rezerwacje dla pasazera */
CREATE OR REPLACE FUNCTION get_bookings_by_passenger(p_passenger_id INTEGER)
RETURNS TABLE (
    booking_id INTEGER,
    booking_date TIMESTAMP,
    booking_status VARCHAR(20),
    flight_number VARCHAR(20),
    departure_time TIMESTAMP,
    arrival_time TIMESTAMP,
    seat_number VARCHAR(10),
    seat_class VARCHAR(10)
)
AS $$
BEGIN
    RETURN QUERY
        SELECT b.booking_id, b.booking_date, b.booking_status,
               f.flight_number, f.departure_time, f.arrival_time,
               s.seat_number, s.seat_class
        FROM Bookings b
        JOIN Flights f ON b.flight_id = f.flight_id
        JOIN Seats s ON b.seat_id = s.seat_id
        WHERE b.passenger_id = p_passenger_id;
END;
$$ LANGUAGE plpgsql;

/* 8. zwroc liste pasezerow dla konkretnego lotu */
CREATE OR REPLACE FUNCTION get_passengers_for_flight(flight_num VARCHAR)
  RETURNS TABLE (passenger_id INT, firstname VARCHAR, lastname VARCHAR, email VARCHAR) AS $$
DECLARE
  passenger_record RECORD;
  passengers_cur CURSOR FOR
    SELECT p.passenger_id, p.firstname, p.lastname, p.email
    FROM Passengers p
    INNER JOIN Bookings b ON p.passenger_id = b.passenger_id
    INNER JOIN Flights f ON b.flight_id = f.flight_id
    WHERE f.flight_number = flight_num;
BEGIN
  FOR passenger_record IN passengers_cur LOOP
    passenger_id := passenger_record.passenger_id;
    firstname := passenger_record.firstname;
    lastname := passenger_record.lastname;
    email := passenger_record.email;
    RETURN NEXT;
  END LOOP;
END;
$$ LANGUAGE plpgsql;