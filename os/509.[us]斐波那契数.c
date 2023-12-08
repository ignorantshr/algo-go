/*
 * @lc app=leetcode.cn id=509 lang=c
 *
 * [509] 斐波那契数
 *
 * https://leetcode.cn/problems/fibonacci-number/description/
 *
 * algorithms
 * Easy (66.13%)
 * Likes:    707
 * Dislikes: 0
 * Total Accepted:    607.5K
 * Total Submissions: 920.2K
 * Testcase Example:  '2'
 *
 * 斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1
 * 开始，后面的每一项数字都是前面两项数字的和。也就是：
 *
 *
 * F(0) = 0，F(1) = 1
 * F(n) = F(n - 1) + F(n - 2)，其中 n > 1
 *
 *
 * 给定 n ，请计算 F(n) 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 2
 * 输出：1
 * 解释：F(2) = F(1) + F(0) = 1 + 0 = 1
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 3
 * 输出：2
 * 解释：F(3) = F(2) + F(1) = 1 + 1 = 2
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 4
 * 输出：3
 * 解释：F(4) = F(3) + F(2) = 2 + 1 = 3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 30
 *
 *
 */
#include <memory.h>
#include <stdio.h>
#include <stdlib.h>

// @lc code=start

int fib(int n) {
    return fib4(n);
}

// 初级版本
int fib1(int n) {
    if (n <= 1) {
        return n;
    }
    return fib1(n - 1) + fib1(n - 2);
}

int fib2_body(int n, int mem[]);
// 备忘录版本
int fib2(int n) {
    int* mem = (int*)calloc(n + 1, sizeof(int));

    return fib2_body(n, mem);
}

int fib2_body(int n, int mem[]) {
    if (n <= 1) {
        return n;
    }

    if (mem[n] != 0) {
        return mem[n];
    }
    mem[n] = fib2_body(n - 1, mem) + fib2_body(n - 2, mem);
    return mem[n];
}

// 从底向上递推版本
int fib3(int n) {
    int* mem = (int*)calloc(n + 1, sizeof(int));

    mem[0] = 0, mem[1] = 1;

    for (int i = 2; i <= n; i++) {
        mem[i] = mem[i - 1] + mem[i - 2];
    }

    return mem[n];
}

// 从底向上递推优化版本
int fib4(int n) {
    if (n <= 1) {
        return n;
    }

    int a, b, res;

    a = 0, b = 1;
    for (int i = 2; i <= n; i++) {
        res = a + b;
        a = b;
        b = res;
    }

    return res;
}

// @lc code=end

int test(int(func)(int)) {
    int result[] = {0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377};
    for (int i = 0; i < sizeof(result) / sizeof(int); i++) {
        int got = func(i);
        if (got != result[i]) {
            printf("%d got: %d, want: %d\n", i, got, result[i]);
            return -1;
        }
    }

    return 0;
}

int main(int argc, char const* argv[]) {
    int (*fibs[])(int) = {fib1, fib2, fib3, fib4};
    int size = sizeof(fibs) / sizeof(fibs[0]);

    for (int i = 0; i < size; i++) {
        if (test(fibs[i]) != 0) {
            printf("func %d failed\n", i + 1);
        } else {
            printf("func %d succeed\n", i + 1);
        }
    }

    return 0;
}
