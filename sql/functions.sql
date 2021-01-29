create or replace function get_schedule_marks(studentid int, from_date date, to_date date) returns table(dayoweek varchar, les_time time, room int, subject varchar, mark int) as
$$
declare
	stud_class_id int;
begin
	stud_class_id = (select cu.class_id from member inner join class_user cu on member.userid = cu.student_id where member.userid = $1 limit 1);
	return query select sch.dayoweek, sch.time, sch.room, sub.name, sm.mark from schedule sch
		inner join subject_schedule ss on ss.schedule_id = sch.scheduleid
		inner join subject sub on sub.subjectid = ss.subject_id
		inner join student_mark sm on sm.subject_id = sub.subjectid
	where sch.class_id = stud_class_id and sm.student_id = $1 and sm.stud_date >= $2 and sm.stud_date <= $3
	order by sch.time;
end;
$$ language plpgsql;

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
