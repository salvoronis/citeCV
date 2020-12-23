create table class (
	class_id	serial primary key,
	name		varchar(3) unique not null
);

create table student (
	student_id	serial primary key,
	class		integer references class(class_id) on delete cascade on update cascade,
	nickname	varchar(30) unique not null,
	password	varchar(64) not null,
	name		varchar(35) not null
);

create table headman (
	student_id	integer references student(student_id) on delete cascade on update cascade,
	class_id	integer references class(class_id) on delete cascade on update cascade
);

create table teacher (
	teacher_id	serial primary key,
	name		varchar(35) not null,
	nickname	varchar(30) unique not null,
	password	varchar(64) not null
);

create table class_teacher (
	class_id	integer references class(class_id) on delete cascade on update cascade,
	teacher_id	integer references teacher(teacher_id) on delete cascade on update cascade
);

create table roles (
	role_id		serial primary key,
	name		varchar(20) unique not null
);

create table teacher_roles (
	roles_id	integer references roles(role_id) on delete cascade on update cascade,
	teacher_id	integer references teacher(teacher_id) on delete cascade on update cascade
);

create table news (
	news_id		serial primary key,
	header		varchar(50) not null,
	tags		text,
	body		text not null,
	date		date not null
);

create table news_author (
	teacher_id	integer references teacher(teacher_id) on delete cascade on update cascade,
	news_id		integer references news(news_id) on delete cascade on update cascade
);

create table subject (
	subject_id	serial primary key,
	name		varchar(50) unique not null,
	course_hours	integer,
	credits		integer
);

create table teacher_subject (
	teacher_id	integer references teacher(teacher_id) on delete cascade on update cascade,
	subject_id	integer references subject(subject_id) on delete cascade on update cascade
);

create table schedule (
	schedule_id	serial primary key,
	dayoweek	varchar(3) not null,
	time		varchar(5) not null,
	room		integer,
	class_id	integer references class(class_id) on delete cascade on update cascade
);

create table schedule_of_subject (
	schedule_id	integer references schedule(schedule_id) on delete cascade on update cascade,
	subject_id	integer references subject(subject_id) on delete cascade on update cascade
);

create table subject_of_class (
	class_id	integer references class(class_id) on delete cascade on update cascade,
	subject_id	integer references subject(subject_id) on delete cascade on update cascade
);

create table students_mark (
	mark_id		serial primary key,
	value		integer,
	student		integer references student(student_id) on delete cascade on update cascade,
	teacher		integer references teacher(teacher_id) on delete cascade on update cascade,
	subject		integer references subject(subject_id) on delete cascade on update cascade
);
