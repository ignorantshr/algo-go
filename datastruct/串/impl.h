#include <stdarg.h>
#include <stdbool.h>

typedef struct sstr {
} sstr;

void StrAssign(sstr* t, char* s);  // 赋值操作。把串T赋值为chars
void Strcopy(sstr* t, sstr s);     // 复制操作，由串S复制得到串T
bool StrEmpty(sstr s);  // 判空操作。若S为空串，则返回True, 否则返回False
int StrLength(sstr s);        // 求串长。返回串S的元素个数
void ClearString(sstr* s);    // 清空操作。将S清为空串
void DestroyString(sstr* s);  // 销毁串。将串S销毁（回收存储空间）
void Concat(sstr* t,
            sstr s1,
            sstr s2);  // 串联接。用T返回由s1和s2联接而成的新串
bool SubString(sstr* sub,
               sstr s,
               int pos,
               int len);  // 求子串。用Sub返回串S的第pos个字符起长度为len的子串
int Index(
    sstr s,
    sstr
        t);  // 定位操作。若主串S中存在与串T相同的子串，则返回它在主串S中第一次出现的位置；否则函数值为-1
int StrCompare(sstr s,
               sstr t);  // 比较操作。若S > T，则返回值 > 0；若S =
                         // T，则返回值 = 0；若S < T，则返回值 < 0

// void StrAssign(sstr* t, char s[]);
// void Strcopy(sstr* t, sstr s);
// bool StrEmpty(sstr s);
// int StrLength(sstr s);
// void ClearString(sstr* s);
// void DestroyString(sstr* s);
// void Concat(sstr* t, sstr s1, sstr s2);
// bool SubString(sstr* sub, sstr s, int pos, int len);
// int Index(sstr s, sstr t);
// int StrCompare(sstr s, sstr t);
