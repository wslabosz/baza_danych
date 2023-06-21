DROP TRIGGER IF EXISTS insert_payment_record_trigger ON Bookings;
DROP TRIGGER IF EXISTS update_booking_status_trigger ON Payments;
DROP TRIGGER IF EXISTS update_booking_status_on_flight_status_change_trigger ON Flights;
DROP TRIGGER IF EXISTS update_payment_status_on_booking_status_change_trigger ON Bookings;
DROP TRIGGER IF EXISTS update_flight_status_on_departure_time_reached_trigger ON Flights;

/* 1. Stworz rekord w tabeli platnosci po rezerwacji */
CREATE OR REPLACE FUNCTION insert_payment_record()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO Payments (booking_id, amount, payment_date, payment_status)
    VALUES (NEW.booking_id, 0.0, CURRENT_TIMESTAMP, 'pending');

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER insert_payment_record_trigger
AFTER INSERT ON Bookings
FOR EACH ROW
EXECUTE FUNCTION insert_payment_record();

/* 2. Potwierdz rezerwacje po potwierdzeniu platnosci */
CREATE OR REPLACE FUNCTION update_booking_status()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.payment_status = 'paid' THEN
        UPDATE Bookings
        SET booking_status = 'confirmed'
        WHERE booking_id = NEW.booking_id;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_booking_status_trigger
AFTER INSERT OR UPDATE OF payment_status ON Payments
FOR EACH ROW
EXECUTE FUNCTION update_booking_status();

/* 3. Anuluj rezerwacje na anulowany lot */
CREATE OR REPLACE FUNCTION update_booking_status_on_flight_status_change()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.flight_status = 'cancelled' THEN
        UPDATE Bookings
        SET booking_status = 'cancelled'
        WHERE flight_id = NEW.flight_id
        AND booking_status <> 'cancelled';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_booking_status_on_flight_status_change_trigger
AFTER UPDATE OF flight_status ON Flights
FOR EACH ROW
EXECUTE FUNCTION update_booking_status_on_flight_status_change();

/* 4. Anuluj platnosc przy anulowanej rezerwacji */
CREATE OR REPLACE FUNCTION update_payment_status_on_booking_status_change()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.booking_status = 'cancelled' THEN
        UPDATE Payments
        SET payment_status = 'cancelled'
        WHERE booking_id = NEW.booking_id;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_payment_status_on_booking_status_change_trigger
AFTER UPDATE OF booking_status ON Bookings
FOR EACH ROW
EXECUTE FUNCTION update_payment_status_on_booking_status_change();

/* 5. Ustaw status lotu na wylot gdy nastepuje czas odlotu */
CREATE OR REPLACE FUNCTION update_flight_status_on_departure_time_reached()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.departure_time <= NOW() THEN
        UPDATE Flights
        SET flight_status = 'departed'
        WHERE flight_id = NEW.flight_id;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_flight_status_on_departure_time_reached_trigger
BEFORE UPDATE OF departure_time ON Flights
FOR EACH ROW
EXECUTE FUNCTION update_flight_status_on_departure_time_reached();