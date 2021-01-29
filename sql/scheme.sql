create table member (
	userId serial primary key,
	login varchar(30) unique not null,
	password varchar(64) not null,
	firstname varchar(25) not null,
	secondname varchar(25) not null,
	email varchar(50) not null
);

create table roles(
	roleId serial primary key,
	name varchar(20) unique not null
);

create table user_roles (
	role_id integer references roles(roleId) on delete cascade on update cascade,
	user_id integer references member(userId) on delete cascade on update cascade
);

create table news (
	newsId serial primary key,
	header varchar(50) not null,
	body text not null,
	date date not null
);

create table news_author (
	news_id integer references news(newsId) on delete cascade on update cascade,
	author_id integer references member(userId) on delete cascade on update cascade
);

create table class (
	classId serial primary key,
	name varchar(5) unique not null
);

create table class_user (
	student_id integer unique references member(userId) on delete cascade on update cascade,
	class_id integer references class(classId) on delete cascade on update cascade
);

create table headman (
	student_id integer references member(userId) on delete cascade on update cascade,
	class_id integer references class(classId) on delete cascade on update cascade
);

create table class_teacher (
	class_id integer references class(classId) on delete cascade on update cascade,
	teacher_id integer references member(userId) on delete cascade on update cascade
);

create table subject (
	subjectId serial primary key,
	name varchar(25) unique not null,
	course_hours integer
);

create table subject_class (
	class_id integer references class(classId) on delete cascade on update cascade,
	subject_id integer references subject(subjectId) on delete cascade on update cascade
);

create table teacher_subject (
	teacher_id integer references member(userId) on delete cascade on update cascade,
	subject_id integer references subject(subjectId) on delete cascade on update cascade
);

create table student_mark (
	markId serial primary key,
	student_id integer references member(userId) on delete cascade on update cascade,
	teacher_id integer references member(userId) on delete cascade on update cascade check (student_id != teacher_id),
	mark integer not null,
	subject_id integer references subject(subjectId) on delete cascade on update cascade
);

create table schedule (
	scheduleId serial primary key,
	class_id integer references class(classId) on delete cascade on update cascade,
	dayoweek varchar(10) not null,
	time time not null,
	room integer not null,
	mark_date date,
	stud_date date not null
);

create table subject_schedule(
	subject_id integer references subject(subjectId) on delete cascade on update cascade,
	schedule_id integer references schedule(scheduleId) on delete cascade on update cascade
);
