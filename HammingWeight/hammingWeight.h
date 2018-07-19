#ifndef _HAMMINGWEIGHT_
#define _HAMMINGWEIGHT_

#include<stdint.h>

int hammingWeight(uint32_t n){
    int count=0;
    while(n){
        count++;
        n&=(n-1);
    }
    return count;
}

#endif
