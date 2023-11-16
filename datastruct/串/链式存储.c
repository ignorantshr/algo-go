#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// 这种存储方式每个字符占1B，每个指针4B（32位下），存储密度较低。因此可以适当改进，让存储密度提高
// typedef struct lnode {
//     char ch;
//     struct lnode* next;
// } lnode, *lstr;

typedef struct lnode {
    char ch[4];  // 每个结点存多个字符，没有字符的位置用'#'或'\0'补足
    struct lnode* next;
} lnode, *lstr;
