CC = gcc
CFLAGS = -Wall -Wextra -I./includes
LDFLAGS =

SRC := $(wildcard src/*.c) $(wildcard src/*/*.c)
OBJ := $(SRC:.c=.o)

EXEC = work

.PHONY: all clean fclean re

all: $(EXEC)

$(EXEC): $(OBJ)
	$(CC) $(LDFLAGS) -o $@ $^

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

clean:
	rm -f $(OBJ)

fclean: clean
	rm -f $(EXEC)

re: fclean all
