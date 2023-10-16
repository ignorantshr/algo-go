#include "semaphore.c"

semaphore plate = {1};

semaphore finish = {0};

semaphore m1 = {0};
semaphore m2 = {0};
semaphore m3 = {0};

void provider()
{
    int i = 0;
    while (1)
    {
        switch (i)
        {
        case 0:
            p(plate);
            v(m1); // 从提供第一种组合材料开始
            v(plate);
            break;
        case 1:
            p(plate);
            v(m2);
            v(plate);
            break;
        case 2:
            p(plate);
            v(m3);
            v(plate);
            break;
        default:
            break;
        }
        i = (++i) % 3;
        p(finish);
    }
}

void smoker1()
{
    while (1)
    {
        p(m1);
        // get m1
        v(finish);
    }
}

void smoker2()
{
    while (1)
    {
        p(m2);
        // get m2
        v(finish);
    }
}

void smoker3()
{
    while (1)
    {
        p(m3);
        // get m3
        v(finish);
    }
}
