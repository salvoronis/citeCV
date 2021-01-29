drop table if exists 	subject_schedule,
			schedule,
			student_mark,
			teacher_subject,
			subject_class,
			subject,
			class_teacher,
			headman,
			class_user,
			class,
			news_author,
			news,
			user_roles,
			roles,
			member cascade;

drop function if exists get_schedule_marks(classid int, studentid int) cascade;
