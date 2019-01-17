BEGIN;

CREATE SEQUENCE bottomline.process_configuration_id_seq;
CREATE TABLE IF NOT EXISTS bottomline.process_configuration (
	id bigint DEFAULT nextval('bottomline.process_configuration_id_seq'::regclass) NOT NULL,
	"name" text NOT NULL,
	description text,
	"configuration" json,
	PRIMARY KEY(id)
);

COMMIT;
