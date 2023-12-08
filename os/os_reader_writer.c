#include "semaphore.c"
#include <stdlib.h>

semaphore *rw = 1;
semaphore *mutex = 1;
semaphore *write_first = 1;

int count = 0; // reader count

void read()
{
    while (1)
    {
        p(write_first);
        p(mutex);
        if (count == 0)
        {
            p(rw);
        }
        count++;
        v(mutex);
        v(write_first); // wake writer or reader

        // read the file

        p(mutex);
        count--;
        if (count == 0)
        {
            v(rw);
        }
        v(mutex);
    }
}

void write()
{
    while (1)
    {
        p(write_first);
        p(rw);
        // write the file
        v(rw);
        v(write_first);
    }
}