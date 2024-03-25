#include "medium.h"
#include <stdio.h>
#include <stdlib.h>

const int MAX_NODE_NUM = 550;
int arr[MAX_NODE_NUM];

int main(void) {
  int n;
  scanf("%d", &n);
  for (int i = 0; i < n; i++) {
    scanf("%d", &arr[i]);
  }
  struct ListNode *head =
      (struct ListNode *)malloc(sizeof(struct ListNode) * 1);
  head->val = arr[0];

  struct ListNode *current_node = head;
  for (int i = 1; i < n; i++) {
    struct ListNode *new_node =
        (struct ListNode *)malloc(sizeof(struct ListNode) * 1);
    current_node->next = new_node;
    new_node->val = arr[i];
    current_node = new_node;
  }
  int left, right;
  scanf("%d %d", &left, &right);
  struct ListNode *new_head = reverseBetween(head, left, right);
  walk(new_head);
}
