struct ListNode {
  int val;
  struct ListNode *next;
};

struct ListNode *reverseBetween(struct ListNode *head, int left, int right);
void walk(struct ListNode *head);
