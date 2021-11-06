CREATE OR REPLACE PROCEDURE public.add_arrest(IN _age_at_arrest varchar, IN _gender varchar, IN _date_of_arrest date, IN _arrest_location text, IN _arresting_officer text, IN _arresting_agency text, IN _charge text)
LANGUAGE 'sql'
AS $BODY$
INSERT INTO public.arrest(
	age_at_arrest
	, gender
	, date_of_arrest
	, arrest_location
	, arresting_officer
	, arresting_agency
	, charge
)
	VALUES (
		_age_at_arrest
		, _gender
		, _date_of_arrest
		, _arrest_location
		, _arresting_officer
		, _arresting_agency
		, _charge);
$BODY$;