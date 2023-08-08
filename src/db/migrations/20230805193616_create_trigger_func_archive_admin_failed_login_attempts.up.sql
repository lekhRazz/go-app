CREATE OR REPLACE FUNCTION process_admin_failed_login_attempts_archive() RETURNS TRIGGER AS $archive$
    BEGIN
        --
        -- Create a row in emp_audit to reflect the operation performed on emp,
        -- making use of the special variable TG_OP to work out the operation.
        --
       
        IF (TG_OP = 'UPDATE' AND NEW.deleted = true) THEN
            INSERT INTO admin_failed_login_attempts_archive SELECT * FROM admin_failed_login_attempts WHERE deleted = true;
			      DELETE FROM admin_failed_login_attempts WHERE deleted = true;
        END IF;
        RETURN NULL; -- result is ignored since this is an AFTER trigger
    END;
$archive$ LANGUAGE plpgsql;

CREATE TRIGGER archive
AFTER INSERT OR UPDATE OR DELETE ON admin_failed_login_attempts
    FOR EACH ROW EXECUTE FUNCTION process_admin_failed_login_attempts_archive();