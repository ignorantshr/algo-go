
// 整型信号量
// 不满足“让权等待”，会发生忙等
typedef struct semaphore {
    int value;
} semaphore;

void wait(semaphore s) {
    while (s.value <= 0) {
    }
    s.value--;
}

void signal(semaphore s) {
    s.value++;
}

// 记录型信号量
typedef struct recordSemaphore {
    int value;
    struct process* l;
} recordSemaphore;

void wait(recordSemaphore s) {
    s.value--;
    if (s.value < 0) {
        block(s.l);  // 将进程挂到阻塞队列，实现了自我阻塞，遵循了“让权等待原则”
    }
}

void signal(recordSemaphore s) {
    s.value++;
    // 释放资源后，说明还有其他进程在等待
    if (s.value <= 0) {
        wakeup(s.l);  // 唤醒一个进程使其从阻塞态变成就绪态
    }
}