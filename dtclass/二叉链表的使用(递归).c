#include <stdio.h>
#include <stdlib.h>
#include <malloc.h>
#define MaxSize 100
/*将二叉链表的结点结构定义和各个函数定义放到这里*/
typedef char DataType; /*定义二叉树结点的数据类型，假设为char型*/
typedef struct BiNode  /*定义二叉链表的结点类型*/
{
	DataType data;
	struct BiNode *lchild, *rchild;
} BiNode;
void PreOrder(BiNode *root)
{
	if (root == NULL)
		return; /*递归调用的结束条件*/
	else
	{
		printf("%c ", root->data); /*访问根结点的数据域，为char型*/
		PreOrder(root->lchild);	   /*前序递归遍历root的左子树*/
		PreOrder(root->rchild);	   /*前序递归遍历root的右子树*/
	}
}
void InOrder(BiNode *root)
{
	if (root == NULL)
		return; /*递归调用的结束条件*/
	else
	{
		InOrder(root->lchild);	   /*中序递归遍历root的左子树*/
		printf("%c ", root->data); /*访问根结点的数据域，为char型*/
		InOrder(root->rchild);	   /*中序递归遍历root的右子树*/
	}
}

void PostOrder(BiNode *root)
{
	if (root == NULL)
		return; /*递归调用的结束条件*/
	else
	{
		PostOrder(root->lchild);   /*后序递归遍历root的左子树*/
		PostOrder(root->rchild);   /*后序递归遍历root的右子树*/
		printf("%c ", root->data); /*访问根结点的数据域，为char型*/
	}
}
void LeverOrder(BiNode *root)
{
	int rear;
	BiNode *q = NULL, *Q[MaxSize]; /*采用顺序队列*/
	int front = rear = -1;		   /*初始化顺序队列*/
	if (root == NULL)
		return;			  /*二叉树为空，算法结束*/
	Q[++rear] = root;	  /*根指针入队*/
	while (front != rear) /*当队列非空时*/
	{
		q = Q[++front];			/*出队*/
		printf("%c ", q->data); /*访问结点，为char型*/
		if (q->lchild != NULL)
			Q[++rear] = q->lchild;
		if (q->rchild != NULL)
			Q[++rear] = q->rchild;
	}
}
BiNode *CreatBiTree(BiNode *root)
{
	char ch;
	scanf("%c", &ch); /*输入结点的数据信息*/
	if (ch == '#' || ch == '\n')
		root = NULL; /*递归结束，建立一棵空树*/
	else
	{
		root = (BiNode *)malloc(sizeof(BiNode));  /*生成新结点*/
		root->data = ch;						  /*新结点的数据域为ch*/
		root->lchild = CreatBiTree(root->lchild); /*递归建立左子树*/
		root->rchild = CreatBiTree(root->rchild); /*递归建立右子树*/
	}
	return root;
}
void DestroyBiTree(BiNode *root)
{
	if (root == NULL)
		return;
	DestroyBiTree(root->lchild);
	DestroyBiTree(root->rchild);
	free(root);
}

int main()
{
	BiNode *root = NULL;	  /*定义二叉树的根指针变量*/
	root = CreatBiTree(root); /*建立一棵二叉树*/
	printf("该二叉树的根结点是：%c\n", root->data);
	printf("\n该二叉树的前序遍历序列是：");
	PreOrder(root);
	printf("\n该二叉树的中序遍历序列是：");
	InOrder(root);
	printf("\n该二叉树的后序遍历序列是：");
	PostOrder(root);
	printf("\n该二叉树的层序遍历序列是：");
	LeverOrder(root);
	DestroyBiTree(root);
	return 0;
}
