#define ElemType int

// 二叉树的结点（链式存储）
typedef struct BiTNode {
    ElemType data;                    // 数据域
    struct BiTNode *lchild, *rchild;  // 左右孩子指针
    // struct BitNode* parent;           // 父结点指针
} BiTNode, *BiTree;

// 二叉树的存储——孩子兄弟表示法
typedef struct CSNode {
    ElemType data;                            // 数据域
    struct CSNode *firstchild, *nextsibling;  // 第一个孩子和右兄弟指针
} CSNode, *CSTree;

// 树的先根遍历
// 用孩子兄弟法来存储这棵树，就会发现树的先根遍历序列与这棵树相应二叉树的先序序列相同
// void PreOrder(TreeNode* R) {
//     if (R != NULL) {
//         visit(R);  // 访问根节点
//         while (R还有下一棵子树T) {
//             PreOrder(T);  // 先根遍历下一棵子树
//         }
//     }
// }

// 树的后根遍历
// 树的后根遍历序列与这棵树相应二叉树的中序序列相同
// void PostOrder(TreeNode* R) {
//     if (R != NULL) {
//         while (R还有下一个子树T) {
//             PostOrder(T);  // 后根遍历下一棵子树
//         }
//         visit(R);  // 访问根节点
//     }
// }
