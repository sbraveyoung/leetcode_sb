#ifndef __HAS_CYCLE__
#define __HAS_CYCLE__

#include<stdio.h>

struct ListNode{
    int val;
    ListNode *next;
    ListNode(int x):val(x),next(NULL){}
};

class Solution{
public:
    bool hasCycle(ListNode *head){
        if (head==NULL || head->next==NULL){
            return false;
        }
        ListNode *fast=head->next->next,*slow=head;
        while(fast!=NULL){
            if (fast==slow){
                return true;
            }
            if (fast->next==NULL){
                return false;
            }
            if (fast->next==slow){
                return true;
            }
            fast=fast->next->next;
            slow=slow->next;
        }
        return false;
    }
};

#endif
