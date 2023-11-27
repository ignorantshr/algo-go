#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define max(x, y) ((x) > (y) ? (x) : (y))

typedef enum {
    RED = 0,
    BLACK,
} nColor;

// 红黑树的结点定义
typedef struct RBnode {
    int key;                // 关键字的值
    struct RBnode* parent;  // 父节点指针
    struct RBnode* lchild;  // 左孩子指针
    struct RBnode* rchild;  // 右孩子指针
    nColor
        color;  // 结点颜色，如：可用 0/1 表示 黑/红，也可使用枚举型enum表示颜色
} RBnode, *RBTree;

RBnode* createNode(int v) {
    RBnode* node = (RBnode*)malloc(sizeof(RBnode));
    node->key = v;
    node->color = RED;
    node->parent = NULL;
    node->lchild = node->rchild = NULL;
    return node;
}

RBnode* leftRotate(RBTree* tree, RBnode* p) {
    RBnode* c = p->rchild;

    p->rchild = c->lchild;
    if (c->lchild != NULL) {
        c->lchild->parent = p;
    }
    c->lchild = p;

    c->parent = p->parent;
    if (p->parent == NULL) {
        *tree = c;
    } else if (p->parent->lchild == p) {
        p->parent->lchild = c;
    } else {
        p->parent->rchild = c;
    }
    p->parent = c;

    return c;
}

RBnode* rightRotate(RBTree* tree, RBnode* p) {
    RBnode* c = p->lchild;

    p->lchild = c->rchild;
    if (c->rchild != NULL) {
        c->rchild->parent = p;
    }
    c->rchild = p;

    c->parent = p->parent;
    if (p->parent == NULL) {
        *tree = c;
    } else if (p->parent->lchild == p) {
        p->parent->lchild = c;
    } else {
        p->parent->rchild = c;
    }
    p->parent = c;

    return c;
}

nColor getColor(RBnode* n) {
    return n == NULL ? BLACK : n->color;
}

/* 黑路同 */
void fixup(RBTree* tree, RBnode* self) {
    if (self == NULL) {
        return;
    }

    /* 根叶黑 */
    if (self->parent == NULL) {  // 根结点
        self->color = BLACK;
        return;
    }

    /* 不红红 */

    if (getColor(self->parent) == BLACK) {
        return;
    }

    int type = 0b00;  // 11: LL; 00: RR; 10: LR ; 01: RL
    RBnode* parent = self->parent;
    RBnode* grandpa = parent->parent;
    RBnode* uncle;
    if (grandpa->lchild == parent) {
        type |= 1 << 1;
        uncle = grandpa->rchild;
    } else {
        // type |= 0 << 1;
        uncle = grandpa->lchild;
    }

    // 红叔叔
    if (getColor(uncle) == RED) {
        parent->color = BLACK;
        uncle->color = BLACK;
        grandpa->color = RED;
        fixup(tree, grandpa);
        return;
    }

    if (self->key < parent->key) {
        type |= 1;
    } else {
        // type |= 0;
    }

    // 黑叔叔
    switch (type) {
        case 0b11:
            rightRotate(tree, grandpa);
            parent->color = BLACK;
            grandpa->color = RED;
            break;
        case 0b00:
            leftRotate(tree, grandpa);
            parent->color = BLACK;
            grandpa->color = RED;
            break;
        case 0b10:
            grandpa->lchild = leftRotate(tree, parent);
            rightRotate(tree, grandpa);
            self->color = BLACK;
            grandpa->color = RED;
            break;
        case 0b01:
            grandpa->rchild = rightRotate(tree, parent);
            leftRotate(tree, grandpa);
            self->color = BLACK;
            grandpa->color = RED;
            break;
        default:
            printf("impossible type: %d", type);
            exit(1);
    }
}

RBTree insert(RBTree root, int v) {
    if (root == NULL) {
        root = createNode(v);
        root->color = BLACK;
        return root;
    }

    RBnode* parent;
    RBnode* cur = root;

    /* 左根右 */

    // 定位插入位置
    while (cur != NULL) {
        parent = cur;
        if (v < cur->key) {
            cur = cur->lchild;
        } else if (v > cur->key) {
            cur = cur->rchild;
        } else {
            return root;
        }
    }

    cur = createNode(v);
    cur->parent = parent;

    if (v < parent->key) {
        parent->lchild = cur;
    } else if (v > parent->key) {
        parent->rchild = cur;
    }

    fixup(&root, cur);
    return root;
}

// 将结点u替换为结点v
void RBTransplant(RBTree* T, RBnode* u, RBnode* v) {
    if (u->parent == NULL) {  // u是根节点
        *T = v;
    } else if (u == u->parent->lchild) {  // u是左子节点
        u->parent->lchild = v;
    } else {  // u是右子节点
        u->parent->rchild = v;
    }
    if (v != NULL) {  // 如果v不为空，将其父节点指向u的父节点
        v->parent = u->parent;
    }
}

void removeNode(RBTree* T, int k) {
    if (*T == NULL) {
        return;
    }

    RBnode* parent;
    RBnode* new;  // 接替的顶点
    RBnode* cur = *T;

    // 定位删除的位置
    while (cur != NULL) {
        parent = cur;
        if (k < cur->key) {
            cur = cur->lchild;
        } else if (k > cur->key) {
            cur = cur->rchild;
        } else {
            break;
        }
    }

    if (cur == NULL) {
        return;
    }

    // 被包含在只有一个子树的情况里面了
    // if (cur->lchild == NULL && cur->rchild == NULL) {
    // }

    nColor delete_original_color = cur->color;  // 保存 删除结点 的颜色
    if (cur->lchild == NULL) {
        new = cur->rchild;
        RBTransplant(T, cur, new);
    } else if (cur->rchild == NULL) {
        new = cur->lchild;
        RBTransplant(T, cur, new);
    } else {
        // 找前驱结点
        RBnode* pre = cur->lchild;
        while (pre != NULL && pre->rchild != NULL) {
            pre = pre->rchild;
        }
        delete_original_color = getColor(pre);
        new = pre->lchild;

        if (pre != cur->lchild) {  // 排除特殊情况：pre 是 cur 的左孩子
            RBTransplant(T, pre, new);  // pre 的左孩子顶替 pre 的位置
            pre->lchild = cur->lchild;  // cur 的左孩子挂到 pre 的左孩子
            cur->lchild->parent = pre;  // cur 的左孩子的父节点指向 pre
        }
        pre->rchild = cur->rchild;
        cur->rchild->parent = pre;

        RBTransplant(T, cur, pre);
        pre->color = cur->color;
    }

    free(cur);
    if (delete_original_color ==
        BLACK) {  // 如果删除的是黑色结点，需要修正红黑树性质

        fixup(T, new);  // 从顶替结点的位置开始调整
    }
}

// 中序遍历打印输出
void inorderTraversal(RBnode* node) {
    if (node == NULL) {
        return;
    }
    inorderTraversal(node->lchild);
    printf("%d (%s) ", node->key, (node->color == RED ? "RED" : "BLACK"));
    inorderTraversal(node->rchild);
}

// 横向展开打印
void printTreeHorizontal(RBTree root, int level) {
    // if (root == NULL) {
    //     return;
    // }

    if (root != NULL) {
        printTreeHorizontal(root->rchild, level + 1);
    }

    for (int i = 0; i < level; i++) {
        printf("   ");
    }

    if (root != NULL) {
        printf("%d[%s]\n", root->key, (root->color == RED ? "R" : "B"));
    } else {
        printf("null\n");
    }

    if (root != NULL) {
        printTreeHorizontal(root->lchild, level + 1);
    }
}

void writeToFile(const char* filename, const char* content) {
    FILE* file = fopen(
        filename,
        "a");  // 打开文件，以写入模式打开（如果文件不存在则创建，如果存在则清空内容）
    if (file == NULL) {
        printf("无法打开文件 %s\n", filename);
        return;
    }

    fprintf(file, "%s", content);  // 将内容写入文件

    fclose(file);  // 关闭文件
}

int lvls[10];
char* lines[10];
int height;
int pos[10][100];  // i 层 j 个结点的位置
const int gap = 3;
char null[gap + 1];

int treeHeight(RBTree t) {
    if (t == NULL) {
        return 0;
    }

    return max(treeHeight(t->lchild), treeHeight(t->rchild)) + 1;
}

int bottomNum(RBTree node, int level) {
    int n = pow(2, height - 1);  // 底层结点总数
    int splitn = pow(2, level);  // 当前层结点总数
    int per = n / splitn;        // 划分
    int i = lvls[level];         // 当前层第几个结点
    return i * per +
           per / 2;  // 计算底层结点比自己小的结点数 i * per + per/2, i from 0
}

int brotherPos(int level) {
    if (lvls[level] == 0) {
        return 0;
    }
    return pos[level][lvls[level] - 1];
}

// 横向展开打印
void printTreeVertical(RBTree root, int level, int h) {
    int n = powl(2, level);
    char pre[200];
    pre[0] = '\0';
    int bnn = bottomNum(root, level);
    bnn = level == height - 1 ? bnn * 2 : bnn * 2 - 1;
    int preCharNum = gap * bnn;

    // 找同层左兄弟结点的尾巴的位置
    int bropos = brotherPos(level);
    preCharNum -= bropos;
    // 防止过长异常
    if (preCharNum <= 0 && lvls[level] != 0) {
        preCharNum = gap / 2;
    }

    for (int i = 0; i < preCharNum; i++) {
        strcat(pre, " ");
    }

    char tmp[300];
    int l;
    if (root == NULL) {
        l = sprintf(tmp, "%s%s", pre, null);
    } else {
        l = sprintf(tmp, "%s%d[%s]", pre, root->key,
                    (root->color == RED ? "R" : "B"));
    }
    tmp[l] = '\0';
    pos[level][lvls[level]] = bropos + l;

    // printf("[%d] line[%d] space:%d, l:%d\n", root == NULL ? -1 : root->key,
    //        level, preCharNum, l);

    if (lines[level] == NULL) {
        lines[level] = (char*)malloc(preCharNum + l);
        strcpy(lines[level], tmp);
    } else {
        size_t newSize =
            strlen(lines[level]) + strlen(tmp) + 1;  // 计算新的内存大小
        char* newBuffer =
            (char*)realloc(lines[level], newSize);  // 重新分配内存
        if (newBuffer != NULL) {
            lines[level] = newBuffer;
            strcat(lines[level], tmp);
        } else {
            // 内存分配失败的处理逻辑
            printf("realloc error\n");
            exit(1);
        }
    }

    lvls[level]++;
    if (lvls[level] == n) {
        // writeToFile("tree.log", strcat(lines[level], "\n"));
        printf("%s\n", lines[level]);
    }

    if (h == height) {
        return;
    }
    if (root == NULL) {
        printTreeVertical(NULL, level + 1, h + 1);
        printTreeVertical(NULL, level + 1, h + 1);
    } else {
        printTreeVertical(root->lchild, level + 1, h + 1);
        printTreeVertical(root->rchild, level + 1, h + 1);
    }
}

void initPrint(RBTree root) {
    height = treeHeight(root);

    for (int i = 0; i < gap; i++) {
        null[i] = 'n';
    }
    null[gap] = '\0';
}

void testing() {
    RBTree t = NULL;
    t = insert(t, 54);
    t = insert(t, 12);
    t = insert(t, 78);
    t = insert(t, 45);
    t = insert(t, 23);
    t = insert(t, 67);
    t = insert(t, 89);
    t = insert(t, 35);
    t = insert(t, 16);
    t = insert(t, 92);
    t = insert(t, 28);
    t = insert(t, 73);
    t = insert(t, 59);
    t = insert(t, 30);
    t = insert(t, 81);
    t = insert(t, 47);
    t = insert(t, 63);
    t = insert(t, 19);
    t = insert(t, 72);
    t = insert(t, 41);
    printf("------------\n");
    // printTreeVertical(t, 0, 1);
    printf("------------\n");
    initPrint(t);
    printTreeVertical(t, 0, 1);
    // removeNode(&t, 5);
    // printf("------------\n");
    // printTreeVertical(t, 0, 1);
}

int main(int argc, char const* argv[]) {
    testing();
    return 0;
}
