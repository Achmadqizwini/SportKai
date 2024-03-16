CREATE TRIGGER update_total_joined_member
AFTER INSERT ON club_member
FOR EACH ROW
BEGIN
    UPDATE club
    SET joined_member = joined_member + 1
    WHERE id = NEW.club_id;
END;