-- drop table if exists persona;

create table persona (
	id				serial,
	identificador 	varchar(255),
	nombre			varchar(255),
	apellido		varchar(255),
	sexo			varchar(1),
	constraint pk_persona primary key (id)
);

create unique index ind_persona
	on persona (identificador);
