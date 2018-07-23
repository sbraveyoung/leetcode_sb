#include<stdio.h>
#include<math.h>
#include<iostream>
#include<stdint.h>
using namespace std;

class Solution{
public:
    uint32_t reverseBits(uint32_t n){
        uint32_t ret=0;
        for(int i = 0; i < 16; i++){
            uint32_t left = n & (0x80000000 >> i);
            uint32_t right = n & (0x1 << i);
            uint32_t tmp = 0x0 | (left >> (31-2*i)) | (right << (31-2*i));
            ret |= tmp;
        }
        return ret;
    }
};
