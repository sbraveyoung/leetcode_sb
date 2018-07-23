#include"func.h"
#include<stdio.h>
#include<vector>
using namespace std;

struct testCase{
    int versions;
    int firstErrorVersion;
    testCase(int v,int fev):versions(v),firstErrorVersion(fev){}
    bool isBadVersion(int version){
        if (version>versions || version<=0){
            return false;
        }
        if (version>=firstErrorVersion){
            return true;
        }
        return false;
    }
};

testCase *p=NULL;

bool isBadVersion(int version){
    if (p!=NULL){
        return p->isBadVersion(version);
    }
    //should not run here.
}

void test(){
    vector<testCase> tcv;
    tcv.push_back(testCase(5,4));
    tcv.push_back(testCase(6,4));
    tcv.push_back(testCase(5,1));
    tcv.push_back(testCase(6,1));
    tcv.push_back(testCase(5,5));
    tcv.push_back(testCase(6,6));
    tcv.push_back(testCase(3,2));
    Solution s;
    for (vector<testCase>::iterator it=tcv.begin();it!=tcv.end();it++){
        p=&*it;
        int ret=s.firstBadVersion(it->versions);
        if (ret!=it->firstErrorVersion){
            printf("failed. versions:%d, firstErrorVersion:%d\n",it->versions,it->firstErrorVersion);
        }
    }
}

int main(){
    test();
}
