#include <stdio.h>

#define buffer_size 3

typedef struct semaphore
{
    int value;
} semaphore;

semaphore *mutex = {1};
semaphore *free = {buffer_size};
semaphore *used = {0};

void producer()
{
    while (1)
    {
        printf("生产产品\n")
            wait(free);
        wait(mutex);
        printf("塞进产品")
            signal(mutex);
        signal(used);
    }
}

void consumer()
{
    while (1)
    {
        wait(used);
        wait(mutex);
        printf("取出一个产品")
            signal(mutex);
        signal(free);
        printf("消耗一个产品")
    }
}

void main()
{
}
