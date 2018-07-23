#ifndef _FIRSTBADVERSION_
#define _FIRSTBADVERSION_

bool isBadVersion(int version);

class Solution{
public:
    int firstBadVersion(int n){
        int left=1,right=n;
        int mid;
        while(left<right){
            mid=left+(right-left)/2;
            if (!isBadVersion(mid)){
                left=mid+1;
            }else{
                right=mid;
            }
        }
        return left;
    }
};

#endif
