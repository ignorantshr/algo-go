
typedef struct graph {
    int vexnum;
} graph;

// 求图G中顶点x的第一个邻接点,若有则返回顶点号。若x没有邻接点或图中不存在x则返回-1
int FirstNeighbor(graph g, int x);

// 假设图G中顶点y是顶点x的一个邻接点,返回除y之外顶点x的下一个邻接点的顶点号,若y是x的最后一个邻接点,则返回-1
int NextNeighbor(graph g, int x, int y);

// 返回邻接点的数量
int Neighbors(graph g, int x, int neighbors[]);