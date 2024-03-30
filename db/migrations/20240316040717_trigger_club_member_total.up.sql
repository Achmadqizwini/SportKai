-- THIS IS FOR MYSQL

-- CREATE TRIGGER update_total_joined_member
-- AFTER INSERT ON club_member
-- FOR EACH ROW
-- BEGIN
--     UPDATE club
--     SET joined_member = joined_member + 1
--     WHERE id = NEW.club_id;
-- END; 

-- THIS IS FOR POSTGRES
CREATE OR REPLACE FUNCTION update_total_joined_member_func()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE club
    SET joined_member = joined_member + 1
    WHERE id = NEW.club_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_total_joined_member
AFTER INSERT ON club_member
FOR EACH ROW
EXECUTE FUNCTION update_total_joined_member_func();
