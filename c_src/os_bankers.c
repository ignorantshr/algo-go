
#define N 5  // 进程数
#define M 3  // 资源类型数量

int Max[N][M];         // 进程声明的最大资源需求数量
int Allocation[N][M];  // 已分配的资源数
int Need[N][M];        // 最多还需要的资源数
int Availiable[M];     // 剩余可用的资源数
int Request[M];        // 一次申请的资源数

int request_source(int i, int Request[M]) {
    // 1. 检查申请的需求数量是不是符合声明
    for (int j = 0; j < M; j++) {
        if (Request[j] > Need[i][j]) {  // 超过最大需求数量
            return -1;
        }
    }

    // 2. 检查剩余可用资源是否满足申请
    // Availiable[i][j] = Availiable[i][j] - Request[j]
    for (int j = 0; j < M; j++) {
        Availiable[j] -= Request[j];
    }

    // 3. 试探着分配，更改各数据结构，并非真正的分配，修改数值只是为了预判
    // Allocation[i][j] = Allocation[i][j] + Request[j]
    for (int j = 0; j < M; j++) {
        Allocation[i][j] += Request[j];
    }

    for (int j = 0; j < M; j++) {
        Need[i][j] -= Request[j];
    }

    // 4. 执行安全性算法，能否把所有进程加入到安全序列
}
