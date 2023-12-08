#include "semaphore.c"

enum status {
    thinking,
    hungry,
    eating,
};

semaphore all_lock = {1};
semaphore* chopsticks[5] = {1, 1, 1, 1, 1};

typedef struct philosopher {
    int status;
    semaphore* left_cs;
    semaphore* right_cs;

    philosopher* left;
    philosopher* right;
} philosopher;

// philosopher *thinker1 = {
//     chopstick1,
//     chopstick2};

// v1 不要全都从一个方向拿筷子，比如让偶数的号的哲学家按照
// “左-》右”的顺序拿筷子，让奇数号的按照“右-》左”的顺序拿筷子。这样保证同时拿起筷子时不会造成死锁
// 问题：同时拿起筷子时只有一个人可以进食

// v2 全局加锁，保证只有一人能同时拿起筷子
// 问题：一人进食，旁边的人可能拿到全局锁，造成对面的两人无法进食
void activityV2(philosopher* thinker) {
    while (1) {
        // thinking...
        // hungry
        p(all_lock);
        // pick up chopsticks
        p(thinker->left_cs);
        p(thinker->right_cs);
        v(all_lock);
        // eating
        v(thinker->left_cs);
        v(thinker->right_cs);
    }
}

// v3 两边人都在非进食的情况下才能进食，否则主动释放全局锁
// 改变及查询状态时也需要加锁
void activityV3(philosopher* thinker) {
    while (1) {
        if (thinker->status != hungry) {
            /* thinking */
            p(all_lock);
            thinker->status = hungry;
            v(all_lock);
        }
        // hungry
        p(all_lock);
        if (thinker->left->status != eating &&
            thinker->right->status != eating) {
            thinker->status = eating;
            // pick up chopsticks
            p(thinker->left_cs);
            p(thinker->right_cs);
        }
        v(all_lock);
        if (thinker->status == eating) {
            /* eating */
            v(thinker->left_cs);
            v(thinker->right_cs);
            p(all_lock);
            thinker->status = thinking;
            v(all_lock);
        }
    }
}
