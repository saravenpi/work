#include "work.h"
#include <string.h>

// prints the help message
void	print_help(void)
{
    printf("Usage: %s[options]\n", PROGRAM_NAME);
    printf("Options:\n");
    printf("\t-h, --help        Display this help message\n");
}

int	main(int argc, char *argv[])
{
	if (argc > 1) {
		if (strcmp(argv[1], "-h") == 0 || strcmp(argv[1], "--help") == 0) {
			print_help();
			return (0);
		}
	}
	start_work_session();
	return (0);
}
