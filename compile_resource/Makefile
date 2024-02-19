PROG = final
OBJS = main.o 
CC = gcc
CFLAGS = -Wall
LDFLAGS = -lm
.PHONY: all
all: $(PROG)
.SUFFIXES: .o .c
.c.o:
	$(CC) $(CFLAGS) -c $<
$(PROG): $(OBJS)
	$(CC) $(CFLAGS) -o $@ $^ $(LDFLAGS)
.PHONY: clean
clean:
	rm -rf $(OBJS) $(PROG)
