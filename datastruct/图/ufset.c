#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAXSIZE 4

typedef struct ufset {
    int parent[MAXSIZE];
} ufset;

// 初始化并查集
ufset* initial() {
    ufset* uf = malloc(sizeof(ufset));
    for (int i = 0; i < MAXSIZE; i++) {
        uf->parent[i] = -1;
    }
    return uf;
}

int Find(ufset* u, int x) {
    int root = x;
    while (u->parent[root] != -1) {
        root = u->parent[root];
    }

    while (x != root) {
        int tmp = u->parent[x];
        u->parent[x] = root;
        x = tmp;
    }

    return root;
}

void Union(ufset* u, int x, int y) {
    int xr = Find(u, x);
    int yr = Find(u, y);
    if (xr == yr) {
        return;
    }

    u->parent[xr] = yr;
}
