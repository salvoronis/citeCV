create or replace function get_schedule_marks(classid int, studentid int, from_date date, to_date date) returns table(dayoweek text, les_time text, subject text, mark int) as $$
	select sch.dayoweek, sch.time, sch.room, sub.name, sm.mark from schedule sch
		inner join subject_schedule ss on ss.schedule_id = sch.scheduleid
		inner join subject sub on sub.subjectid = ss.subject_id
		inner join student_mark sm on sm.subject_id = sub.subjectid
	where sch.class_id = $1 and sm.student_id = $2 and sm.stud_date >= $3 and sm.stud_date <= $4
	order by sch.time
$$ language sql;

create or replace function get_schedule_by_class(classid int) returns table(classname text, dayoweek text, lestime time, room int, subject text, t_login text, t_fname text, t_sname text) as $$
	select class.name, sch.dayoweek, sch.time, sch.room, sub.name, teacher.login, teacher.firstname, teacher.secondname from class
		inner join schedule sch on sch.class_id = class.classid
		inner join subject_schedule ss on ss.schedule_id = sch.scheduleid
		inner join subject sub on sub.subjectid = ss.subject_id
		inner join teacher_subject ts on ts.subject_id = sub.subjectid
		inner join member teacher on teacher.userid = ts.teacher_id
	where class.classid = $1
	order by sch.time
$$ language sql;
