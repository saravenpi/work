#include "work.h"

// creates a string
// that contains the format
// of the given time to
// a human readable date string
// and returns it to
// this format: "YYYY/MM/DD|hh:mm:ss"
char *get_time_format(time_t time) {
	struct tm *tm_time;
	char *time_format_str;

	time_format_str = malloc(20 * sizeof(char));
	if (!time_format_str)
		return (NULL);
	tm_time = localtime(&time);
	sprintf(time_format_str, "%04d/%02d/%02d|%02d:%02d:%02d",
		 tm_time->tm_year + 1900,
		 tm_time->tm_mon + 1,
		 tm_time->tm_mday,
		 tm_time->tm_hour,
		 tm_time->tm_min,
		 tm_time->tm_sec
		);
	return time_format_str;
}
