#include "work.h"

void	save_work_session(time_t start_time, time_t elapsed_time)
{
	char	*start_str;
	char	*end_str;
	char	session_log[512];
	char	file_path[512];
	char	*home_path;

	start_str = get_time_format(start_time);
	end_str = get_time_format(start_time + elapsed_time);
	if (!end_str) {
		printf("ERROR: not enough memory\n");
		return;
	}
	sprintf(session_log, "#SESSION %s-%s\n",
		 start_str,
		 end_str
		 );
	free(start_str);
	free(end_str);
	home_path = getenv("HOME");
	if (home_path == NULL)
	{
		printf("ERROR: could not find $HOME variable to append to  history");
		return;
	}
	sprintf(file_path, "%s/%s", home_path, PROGRAM_HISTORY_FILENAME);
	append_str_to_file(session_log, file_path);
}

void	start_work_session(void)
{
	time_t	start_time;
	time_t	elapsed_time;
	char	*finished_str;

	printf("✨ starting work session, good luck !\n");
	start_time = time(0);
	finished_str = get_time_format((time_t)(start_time + SESSION_LENGTH_SECS));
	if (!finished_str) {
		printf("ERROR: not enough memory\n");
		return;
	}
	printf("session will end at: %s\n", finished_str);
	free(finished_str);
	elapsed_time = 0;
	while (elapsed_time < SESSION_LENGTH_SECS) {
		elapsed_time = time(0) - start_time;
	}
	save_work_session((time_t)start_time, (time_t)elapsed_time);
	printf("🎉 work session finished\n");
}
