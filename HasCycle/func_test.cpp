#include<iostream>
#include "func.h"
using namespace std;

int test1(){
    Solution s;
    ListNode head(1);
    ListNode first(2);
    ListNode third(3);
    head.next=&first;
    first.next=&third;
    if (s.hasCycle(&head)!=false){
        cout<<"test1 failed."<<endl;
    }
}
int test2(){
    Solution s;
    ListNode head(1);
    ListNode first(2);
    ListNode third(3);
    head.next=&first;
    first.next=&third;
    third.next=&first;
    if (s.hasCycle(&head)!=true){
        cout<<"test2 failed."<<endl;
    }
}

int main(){
    test1();
    test2();
    return 0;
}
