-- THIS IS FOR MYSQL
-- DROP TRIGGER IF EXISTS update_total_joined_member;

-- THIS IS FOR POSTGRES
DROP TRIGGER update_total_joined_member ON club_member;
DROP FUNCTION update_total_joined_member_func();
