#include"func.h"

void test(){
    Solution s;
    if (s.reverseBits(0)!=0){
        cout<<"s.reverseBits(0)="<<s.reverseBits(0)<<" not 0"<<endl;
    }
    if (s.reverseBits(0xffffffff)!=0xffffffff){
        cout<<"s.reverseBits(0xffffffff)="<<s.reverseBits(0xffffffff)<<" not 0xffffffff"<<endl;
    }
    if (s.reverseBits(43261596)!=964176192){
        cout<<"s.reverseBits(43261596)="<<s.reverseBits(43261596)<<" not 964176192"<<endl;
    }
}

int main(){
    test();
}
