#pragma once
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define PROGRAM_NAME "work"
#define PROGRAM_HISTORY_FILENAME ".workhistory"
#define SESSION_LENGTH_SECS 1500 // 25 minutes

// src/file.c
void append_str_to_file(char *str, char *path);

// src/time.c
char *get_time_format(time_t time);

// src/work.c
void start_work_session(void);
