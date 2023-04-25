#include <stdio.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

/* 
    删除链表的节点
    给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

    返回删除后的链表的头节点。
 */

struct ListNode* deleteNode(struct ListNode* head, int val){
    if (head == NULL) {
        return NULL;
    }

    if (head->val == val) {
        return head->next;
    }

    struct ListNode* pre = head;
    struct ListNode* cur = head->next;

    while (cur) {
        if (cur->val == val) {
            pre->next = cur->next;
            break;
        }
        pre = cur;
        cur = cur->next;
    }

    return head;
}