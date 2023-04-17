#include <stdio.h>
#include <stdlib.h>
#include <malloc.h>
#define MaxSize 100

typedef char DataType;

typedef struct BiNode {
    DataType data;
    struct BiNode *lchild, *rchild;
} BiNode;

BiNode * Construct(BiNode *root) {
    char ch;
    // 在输入的字符串中只取一个
    scanf("%c", &ch);
    
    if (ch == '#' || ch == '\n') {
        root = NULL;
    } else {
        root = (BiNode *)malloc(sizeof(BiNode));
        root->data = ch;
        root->lchild = Construct(root->lchild);
        root->rchild = Construct(root->rchild);
    }
    
    return root;
}

void Destroy(BiNode *root) {
    if (root == NULL) 
        return;
    Destroy(root->lchild);
    Destroy(root->rchild);
    free(root);
}

void PreOrder (BiNode *root) {
    if (root == NULL)
        return;
    else {
        printf("%c ", root->data);
        PreOrder(root->lchild);
        PreOrder(root->rchild);
    }
}

void InOrder (BiNode *root) {
    if (root == NULL)
        return;
    else {
        InOrder(root->lchild);
        printf("%c ", root->data);
        InOrder(root->rchild);
    }
}

void PostOrder (BiNode *root) {
    if (root == NULL)
        return;
    else {
        PostOrder(root->lchild);
        PostOrder(root->rchild);
        printf("%c ", root->data);
    }
}

int Deep(BiNode *root) {
    if (root == NULL)
        return 0;
    int left = Deep(root->lchild);
    int right = Deep(root->rchild);
    return left > right ? ++left : ++right;
}

int CountNode (BiNode *root) {
    if (root == NULL)
        return 0;
    return CountNode(root->lchild) + CountNode(root->rchild) + 1; 
}

void PrintLeaf(BiNode *root) {
    if (root == NULL)
        return;
    if (root->lchild == NULL && root->rchild == NULL){
        printf("%c ", root->data);
    } else {
        PrintLeaf(root->lchild);
        PrintLeaf(root->rchild);
    }
}

int main() {
    BiNode *tree = NULL;
    tree = Construct(tree);
    printf("\n前序遍历为：");
    PreOrder(tree);
    printf("\n后序遍历为：");
    PostOrder(tree);
    printf("\n输出所有叶子节点：");
    PrintLeaf(tree);
    printf("\n");
}