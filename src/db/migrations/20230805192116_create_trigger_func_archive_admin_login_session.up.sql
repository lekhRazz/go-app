CREATE OR REPLACE FUNCTION process_archive() RETURNS TRIGGER AS $archive$
    BEGIN
        --
        -- Create a row in emp_audit to reflect the operation performed on emp,
        -- making use of the special variable TG_OP to work out the operation.
        --
       
        IF (TG_OP = 'UPDATE' AND NEW.deleted = true) THEN
            INSERT INTO admin_login_session_archive SELECT * FROM admin_login_session WHERE deleted = true;
			DELETE FROM admin_login_session WHERE deleted = true;
        END IF;
        RETURN NULL; -- result is ignored since this is an AFTER trigger
    END;
$archive$ LANGUAGE plpgsql;

CREATE TRIGGER archive
AFTER INSERT OR UPDATE OR DELETE ON admin_login_session
    FOR EACH ROW EXECUTE FUNCTION process_archive();