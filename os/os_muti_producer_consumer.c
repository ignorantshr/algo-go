
typedef struct semaphore
{
    int value;
} semaphore;

semaphore *apple = 0;
semaphore *orange = 0;
semaphore *plate = 1;
semaphore *mutext = 1;

void dad()
{
    while (1)
    {

        p(plate);
        plate_op("put an apple");
        v(apple);
    }
}

void daughter()
{
    while (1)
    {
        p(apple);
        plate_op("eat an apple");
        v(plate);
    }
}

void mom()
{
    while (1)
    {
        p(plate);
        plate_op("put an orange");
        v(orange);
    }
}

void son()
{
    while (1)
    {
        p(orange);
        plate_op("eat an orange");
        v(plate);
    }
}

// 盘子容量为 1 的写法
// 如果大于1，则必须在访问盘子的时候加上一个互斥信号量
void plate_op(char *op)
{
    p(mutext);
    printf("%s\n", op);
    v(mutext);
}