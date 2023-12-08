#include <stdio.h>

#define buffer_size 3

typedef struct semaphore {
    int value;
} semaphore;

semaphore* mutex = {1};
semaphore* free = {buffer_size};
semaphore* used = {0};

void producer() {
    while (1) {
        printf("生产产品\n");
        p(free);
        p(mutex);
        printf("塞进产品");
        v(mutex);
        v(used);
    }
}

void consumer() {
    while (1) {
        p(used);
        p(mutex);
        printf("取出一个产品");
        v(mutex);
        v(free);
        printf("消耗一个产品");
    }
}

// ===========使用管程==============
typedef semaphore condition;
typedef struct monitor {
    condition full, empty;  // 条件变量
    int count;
} monitor;
typedef struct Item {
    /* data */
} Item;

void insert(monitor m, Item item) {
    if (m.count == buffer_size) {
        wait(m.full);
    }
    m.count++;
    insert_item(item);
    if (m.count == 1) {
        signal(m.empty);
    }
}

void remove(monitor m) {
    if (m.count == 0) {
        wait(m.empty);
    }
    m.count--;
    if (m.count == buffer_size - 1) {
        signal(m.full);
    }
}
