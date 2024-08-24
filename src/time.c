#include "work.h"

char	*get_time_format(time_t time)
{
	struct tm	*tm_time;
	char		*time_format_str;

	time_format_str = malloc(9 * sizeof(char));
	if (!time_format_str)
		return (NULL);
	tm_time = localtime(&time);
	sprintf(time_format_str, "%02d:%02d:%02d",
		 tm_time->tm_hour,
		 tm_time->tm_min,
		 tm_time->tm_sec
	);
	return (time_format_str);
}

time_t	get_elapsed_time_from_date_string(const char *date_str)
{
    struct tm input_time = {0};
    time_t input_time_t;
	time_t current_time;
    time_t elapsed_time;

    if (strptime(date_str, "%Y:%m:%d", &input_time) == NULL)
        return -1;
    input_time_t = mktime(&input_time);
    current_time = time(NULL);
	elapsed_time = difftime(current_time, input_time_t);
    return (elapsed_time);
}
