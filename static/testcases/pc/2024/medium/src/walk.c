#include "medium.h"
#include <stdio.h>

void walk(struct ListNode *head) {
  struct ListNode *current = head;
  while (current != NULL) {
    printf("%d\n", current->val);
    current = current->next;
  }
}
