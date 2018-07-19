#include"hammingWeight.h"
#include<stdio.h>

void test(){
    uint32_t arg []={0,1,2,3,0xffffffff};
    int ret []={0,1,1,2,32};
    for (int i=0;i<sizeof(arg)/sizeof(arg[0]);i++){
        if (ret[i]!=hammingWeight(arg[i])){
            printf("error\n");
            return;
        }
    }
    printf("success\n");
}

int main(){
    test();
    return 0;
}
