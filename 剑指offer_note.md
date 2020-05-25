1、在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
```c++
class Solution {
public:
    bool Find(int target, vector<vector<int> > array) {
        int row = array.size();
        if(row <= 0) return false;
        int col = array[0].size();
        if(col <= 0) return false;
        int sx = 0, sy = col - 1;
        while(sx < row && sy >= 0){
            if(array[sx][sy] == target) return true;
            else if(array[sx][sy] < target) sx++;
            else sy--;
        }
        return false;
    }
};
```
2、请实现一个函数，将一个字符串中的每个空格替换成“%20”。例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
```c++
class Solution {
public:
	void replaceSpace(char *str,int length) {
        int spnums = 0;
        for(int i = 0; i < length; i++){
            if(str[i] == ' ') spnums++;
        }
        int j = length + 2 * spnums - 1;
        length--;
        while(length < j){
            if(str[length] != ' ') str[j--] = str[length--];
            else{
                str[j--] = '0';
                str[j--] = '2';
                str[j--] = '%';
                length--;
            }
        }
	}
};
```
3、输入一个链表，按链表从尾到头的顺序返回一个ArrayList。
```c++
class Solution {
public:
    vector<int> printListFromTailToHead(ListNode* head) {
        vector<int> rres;
        for(;head != nullptr ; rres.push_back(head->val), head = head->next);
        return vector<int> (rres.rbegin(), rres.rend());
    }
};
```
4、输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
```c++
class Solution {

public:
    TreeNode* reConstructBinaryTree(vector<int> pre,vector<int> vin) {
       if(pre.size() == 0 || vin.size() == 0) return NULL;
        
        TreeNode* root = new TreeNode(pre[0]);
        int num = find(vin.begin(), vin.end(), pre[0]) - vin.begin();
        
        vector<int> temp1(pre.begin()+1, pre.begin()+1+num);
        vector<int> temp2(vin.begin(),vin.begin()+num);
        root->left = reConstructBinaryTree(temp1, temp2);
        
        vector<int> temp3(pre.begin()+1+num, pre.end());
        vector<int> temp4(vin.begin()+1+num,vin.end());
        root->right = reConstructBinaryTree(temp3, temp4);
        
        return root;
    }
};
```
5、用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。
```c++
class Solution
{//队列是先进先出。
public:
    void push(int node) {
        stack1.push(node);
    }

    int pop() {
        if(stack2.empty()){
            while(!stack1.empty()){
            stack2.push(stack1.top());
            stack1.pop();
            }
        }
        int res = stack2.top();
        stack2.pop();
        return res;
    }

private:
    stack<int> stack1;
    stack<int> stack2;
};
```
6、把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个非递减排序的数组的一个旋转，输出旋转数组的最小元素。
例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。
NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。
```c++
class Solution {
private: 
    int minNumberInRotateArray(vector<int> rotateArray, int left, int right){
        if(left == right) return rotateArray[left];
        else{
            int mid = left + (right - left) / 2;
            if(rotateArray[left] < rotateArray[mid]){
                return minNumberInRotateArray(rotateArray, mid, right);
            }
            else if(rotateArray[mid] < rotateArray[right]){
                return minNumberInRotateArray(rotateArray, left, mid);
            }
            else{
                return minNumberInRotateArray(rotateArray, left+1, right);
            }
        }
    }
public:
    int minNumberInRotateArray(vector<int> rotateArray) {
        //二分
        int n = rotateArray.size();
        //return minNumberInRotateArray(rotateArray, 0, n - 1);
        int left = 0, right = n - 1;
        while(left <= right){
            if(left == right) return rotateArray[left];
            else{
                int mid = left + (right - left) / 2;
                if(rotateArray[left] < rotateArray[mid]){
                    left = mid;
                }
                else if(rotateArray[mid] < rotateArray[right]){
                    right = mid;
                }
                else{
                    left++;
                }
            }
        }
        return -1;
    }
};
```
7、大家都知道斐波那契数列，现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0，第1项是1）。
n<=39
```c++
class Solution {
public:
    int Fibonacci(int n) {
        vector<int> dp;
 
        dp.push_back(0);
        dp.push_back(1);
        dp.push_back(1);
        for(int i = 3; i <= n; i++){
            dp.push_back(dp[i-1] + dp[i-2]);
        }
        return dp[n];
    }
};
```
8、一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
```c++
class Solution {
public:
    int jumpFloor(int number) {
        int a = 1, b = 1;
        while(number--){
            a = a + b;
            b = a - b;
        }
        return b;
    }
};
```
9、一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
```c++
class Solution {
public:
    int jumpFloorII(int number) {
        return pow(2, number - 1);
    }
};
```
10、我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？
```c++
class Solution {
public:
    int rectCover(int number) {
        if(number == 0) return 0;
        int a = 1, b = 1;
        while(number--){
            a = a + b;
            b = a - b;
        }
        return b;
    }
};
```
11、输入一个整数，输出该数二进制表示中1的个数。其中负数用补码表示。
```c++
class Solution {
public:
     int  NumberOf1(int n) {
         //减1会使得最右边的1变0 且其右边会出来一串1，与掉就可以了。
         int res = 0;
         while(n){
             n = n & (n - 1);
             res++;
         }
         return res;
     }
};
```
12、给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。
保证base和exponent不同时为0
```c++
class Solution {
public:
    double Power(double base, int exponent) {
        int n = exponent >= 0 ? exponent : -exponent;
        double res = 0;
        if(exponent == 0) return 1.0;
        if((n & 1) == 0){//偶数
            res = Power(base*base, n >> 1);
        }
        else{
            res = base * Power(base*base, n >> 1);
        }
        return exponent >= 0? res : 1/res;
    }
};
```
13、输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。
```c++
//插入排序的思路
class Solution {
public:
    void reOrderArray(vector<int> &array) {
        int n = array.size();
        int k = 0;
        for(int i = 0; i < n; i++){
            if((array[i] & 0x0001) == 1){
                if(i > k){
                    int j = i;
                    int temp = array[i];
                    while(j > k){
                        array[j] = array[j-1];
                        j--;
                    }
                    array[k] = temp;
                }
                k++;
            }
            
        }
    }
};
```
14、输入一个链表，输出该链表中倒数第k个结点。
```c++
//双指针法
class Solution {
public:
    ListNode* FindKthToTail(ListNode* pListHead, unsigned int k) {
        ListNode* p1 = pListHead;
        ListNode* p2 = pListHead;
        while(k--){
            if(!p1) return nullptr;
            p1 = p1->next;
        }
        while(p1){
            p1 = p1->next;
            p2 = p2->next;
        }
        return p2;
    }
};
```
15、输入一个链表，反转链表后，输出新链表的表头。
```c++
//头插法
class Solution {
public:
    ListNode* ReverseList(ListNode* pHead) {
        ListNode* cur = pHead;
        ListNode* res = NULL;
        while(cur){
            ListNode* temp = cur->next;
            cur->next = res;
            res    = cur;
            cur = temp;
        }
        return res;
    }
};
```
16、输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。
```c++
class Solution {
public:
    ListNode* Merge(ListNode* pHead1, ListNode* pHead2)
    {
        ListNode* p = new ListNode(-1);
        ListNode* res = p;
        int cur = 0;
        while(pHead1 || pHead2){
            if(pHead1 && pHead2){//两个链表都不为空，拿到当前小值并步进
                if(pHead1->val > pHead2->val) {
                    cur = pHead2->val;
                    pHead2 = pHead2->next;
                } 
                else {
                    cur = pHead1->val;
                    pHead1 = pHead1->next;
                }
                
            }
            //反之拿非空，并步进
            else if(!pHead1) {
                cur = pHead2->val;
                pHead2 = pHead2->next;
            }
            else {
                cur = pHead1->val;
                pHead1 = pHead1->next;
            }
            p->next = new ListNode(cur);
            p = p->next;
        }
        p->next = NULL;
        return res->next;
    }
};
```
17、输入两棵二叉树A，B，判断B是不是A的子结构。（ps：我们约定空树不是任意一个树的子结构）
```c++
class Solution {
private:
    bool dfs(TreeNode* p1, TreeNode* p2){
        //子结构即包含完全一致的值
        if(!p2) return true;
        if(!p1) return false;
        if(p1->val == p2->val) {
            return dfs(p1->left, p2->left) && dfs(p1->right, p2->right);
        }
        else {
            return false;
        }
    }
public:
    bool HasSubtree(TreeNode* pRoot1, TreeNode* pRoot2)
    {
        if(!pRoot1) return false;//空树特殊处理
        if(!pRoot2) return false;//空树特殊处理
        return dfs(pRoot1, pRoot2) || HasSubtree(pRoot1->left, pRoot2) || HasSubtree(pRoot1->right, pRoot2);
    }
};
```
18、操作给定的二叉树，将其变换为源二叉树的镜像。
```c++
class Solution {
public:
    void Mirror(TreeNode *pRoot) {
        if(!pRoot) return ;
        TreeNode * temp = pRoot->left;
        pRoot->left = pRoot->right;
        pRoot->right = temp;
        //左右交换，左镜像，右镜像
        Mirror(pRoot->left);
        Mirror(pRoot->right);
    }
};
```
19、输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字，
例如，如果输入如下4 X 4矩阵： 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 则依次打印出数字1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10.
```c++
class Solution {
public:
    vector<int> printMatrix(vector<vector<int> > matrix) {
        int h = matrix.size(), l = matrix[0].size();
        vector<int> res;
        vector<vector<bool> > cango(h , vector<bool>(l, true));//构建可走地图
        // 列加一，行加一，列减一，行减一
        vector<int> dpx = {0, 1, 0, -1};
        vector<int> dpy = {1, 0, -1, 0};
        int curdir = 0;
        int curx = 0, cury = 0;
        int nextx = 0, nexty = 0;
        while(curx >= 0 && curx < h && cury >= 0 && cury < l && cango[curx][cury]){
            res.push_back(matrix[curx][cury]);
            cango[curx][cury] = false;
            nextx = curx + dpx[curdir];
            nexty = cury + dpy[curdir];
            //信值是否可走
            if(nextx >= 0 && nextx < h && nexty >=0 && nexty < l && cango[nextx][nexty]){
                curx = nextx;
                cury = nexty;
            }
            else{
                //不可走的话说明需要转弯了（顺时针转90度）
                curdir = (curdir + 1) % 4;
                curx += dpx[curdir];
                cury += dpy[curdir];
            }
            
        }
        return res;
    }
};
```
20、定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的min函数（时间复杂度应为O（1））。
注意：保证测试中不会当栈为空的时候，对栈调用pop()或者min()或者top()方法。
```c++
class Solution {
private:
    stack<int> theStack;
    stack<int> minStack;
public:
    void push(int value) {
        if(minStack.empty() || value <= minStack.top()){
            minStack.push(value);
        }
        theStack.push(value);
    }
    void pop() {
        if(theStack.top() == minStack.top()){
            minStack.pop();
        }
        theStack.pop();
    }
    int top() {
        return theStack.top();
    }
    int min() {
        return minStack.top();
    }
};
```
21、输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。
假设压入栈的所有数字均不相等。例如序列1,2,3,4,5是某栈的压入顺序，序列4,5,3,2,1是该压栈序列对应的一个弹出序列，但4,3,5,1,2就不可能是该压栈序列的弹出序列。
（注意：这两个序列的长度是相等的）
```c++
class Solution {
public:
    bool IsPopOrder(vector<int> pushV,vector<int> popV) {
        stack<int> intstack;
        int i = 0;
        for(auto c : pushV){
            intstack.push(c);
            while(!intstack.empty() && intstack.top() == popV[i]){
                intstack.pop();
                i++;
            }
        }
        return intstack.empty();
    }
};
```
22、从上往下打印出二叉树的每个节点，同层节点从左至右打印。
```c++
class Solution {
public:
    vector<int> PrintFromTopToBottom(TreeNode* root) {
        vector<int> res;
        if(!root) return res;
        queue<TreeNode*> q;
        q.push(root);
        while(!q.empty()){
            TreeNode* cur = q.front();
            q.pop();
            res.push_back(cur->val);
            if(cur->left) q.push(cur->left);
            if(cur->right) q.push(cur->right);
        }
        return res;
    }
};
```
23、输入一个非空整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则输出Yes,否则输出No。假设输入的数组的任意两个数字都互不相同。
```c++
//后序遍历，左右中，小大中。
class Solution {
public:
    bool VerifySquenceOfBST(vector<int> sequence) {
        int length = sequence.size();
        if(length == 0) return false;
        int i = 0;
        while(--length){
            while(sequence[i] < sequence[length]) i++;
            while(sequence[i] > sequence[length]) i++;
            if(i < length) return false;
            i = 0;
        }
        return true;
    }
};
```
24、输入一颗二叉树的根节点和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。路径定义为从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。
```c++
class Solution {
private:
    void find(TreeNode* root,int expectNumber, vector<int> path, int sum){
        int val = root->val;
        path.push_back(val);
        sum += val;
        if(sum == expectNumber && !root->left && !root->right) allpath.push_back(path); 
        if(root->left) find(root->left,expectNumber, path, sum);
        if(root->right) find(root->right,expectNumber, path, sum);
    }
public:
    vector<vector <int>> allpath;
    vector<vector<int> > FindPath(TreeNode* root,int expectNumber) {
        if(!root) return allpath;
        vector<int> path;
        int sum = 0;
        find(root,expectNumber,path,sum);
        return allpath;
    }
};
```
25、输入一个复杂链表（每个节点中有节点值，以及两个指针，一个指向下一个节点，另一个特殊指针random指向一个随机节点），
请对此链表进行深拷贝，并返回拷贝后的头结点。（注意，输出结果中请不要返回参数中的节点引用，否则判题程序会直接返回空）
```c++
class Solution {
public:
    RandomListNode* Clone(RandomListNode* pHead)
    {
        if(!pHead) return nullptr;
        RandomListNode* cur = pHead, *res = nullptr;
        RandomListNode* temp;
        while(cur){
            //复制节点
            temp = new RandomListNode(cur->label);
            temp->next = cur->next;
            cur->next = temp;
            cur = cur->next->next;
        }
        cur = pHead;
        while(cur){
            //给复制的节点增加random
            if(cur->random){
                cur->next->random = cur->random->next;
            }
            cur = cur->next->next;
        }
        cur = pHead;
        res = cur->next;
        temp = res;
        while(cur){
            //拆分
            cur->next = cur->next->next;
            if(!temp->next) break;
            temp->next = temp->next->next;
            cur = cur->next;
            temp = temp->next;
        }
        return res;
    }
};
```
26、输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求不能创建任何新的结点，只能调整树中结点指针的指向。
```c++
    // 记录子树链表的最后一个节点，终结点只可能为只含左子树的非叶节点与叶节点
    // 最后的一个节点可能为最右侧的叶节点
       
        // 1.将左子树构造成双链表，并返回链表头节点

        // 3.如果左子树链表不为空的话，将当前root追加到左子树链表
// 当根节点只含左子树时，则该根节点为最后一个节点
        // 4.将右子树构造成双链表，并返回链表头节点

        // 5.如果右子树链表不为空的话，将该链表追加到root节点之后


class Solution {
public:
    TreeNode* leftLast = nullptr;
    TreeNode* Convert(TreeNode* pRootOfTree)
    {
        if(!pRootOfTree) return nullptr;
        if(!pRootOfTree->left && !pRootOfTree->right){
            leftLast = pRootOfTree;
            return pRootOfTree;
        }
        TreeNode* pRootleft = Convert(pRootOfTree->left);
        if(pRootleft){
            leftLast->right = pRootOfTree;
            pRootOfTree->left = leftLast;
            leftLast = pRootOfTree;
        }
        TreeNode* pRootright = Convert(pRootOfTree->right);
        if(pRootright){
            pRootOfTree->right = pRootright;
            pRootright->left = pRootOfTree;
        }
        return pRootleft != nullptr ? pRootleft : pRootOfTree;
    }
};
```
27、输入一个字符串,按字典序打印出该字符串中字符的所有排列。例如输入字符串abc,则打印出由字符a,b,c所能排列出来的所有字符串abc,acb,bac,bca,cab和cba。
输入描述:
输入一个字符串,长度不超过9(可能有字符重复),字符只包括大小写字母。
```c++
class Solution {
public:
    vector<string> Permutation(string str) {
        if (str.size() != 0){
            int nSize = str.size();
            Permutation(str, nSize, 0);
        }
        for (set<string>::iterator iter = permutationSet.begin(); iter != permutationSet.end(); ++iter){
            permutation.push_back(*iter);
        }
        return permutation;
    }
    //固定前n的全排列，固定前n后的全排列来自于固定前n+1个（最后那个遍历一边）的全排列。
    void Permutation(string str,int nSize, int n){
        if (n == str.size()){
            permutationSet.insert(str);
        }
        else{
            Permutation(str, nSize, n + 1);//遍历的第一个不需要作交换，后边的按序交换
            for (int i = n+1; i < str.size(); ++i){
                if (str[n] != str[i]){
                    char tmp = str[n];
                    str[n] = str[i];
                    str[i] = tmp;
                    Permutation(str, nSize, n + 1);
                    tmp = str[n];
                    str[n] = str[i];
                    str[i] = tmp;
                }
            }
        }
    }
public:
    vector<string> permutation;
    set<string> permutationSet;
};
```
28、数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。例如输入一个长度为9的数组{1,2,3,2,2,2,5,4,2}。由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。如果不存在则输出0。
```c++
class Solution {
public:
    int MoreThanHalfNum_Solution(vector<int> numbers) {
        int n = numbers.size();
        unordered_map<int, int> thenums;
        for(auto c : numbers){
            if(thenums.find(c) != thenums.end()){
                thenums[c]++;
            }
            else{
                thenums[c] = 1;
            }
            if(thenums[c] > n/2) return c;
        }
       
        return 0;
    }
};
```
29、输入n个整数，找出其中最小的K个数。例如输入4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4,。
```c++
class Solution {
public:
    vector<int> GetLeastNumbers_Solution(vector<int> input, int k) {
        priority_queue<int> Q;//默认大顶堆
        vector<int> res;
        if((k<=0) || input.size() < k) return res;
        for(auto c : input){
            if(Q.size() < k) Q.push(c);
            else if(c < Q.top()){
                Q.pop();
                Q.push(c);
            }
        }
        while(!Q.empty()){
            res.push_back(Q.top());
            Q.pop();
        }
        return res;
    }
};
```
30、计算连续子向量的最大和
```c++
class Solution {
public:
    int FindGreatestSumOfSubArray(vector<int> array) {
        int n = array.size();
        vector<int> curmax(array);//这个数组用来记录以当前元素结尾时的最大和值
        int res = array[0];
        for(int i = 1; i < n; i++){
            if(curmax[i-1] + array[i] >= array[i]){
                curmax[i] = curmax[i-1] + array[i];
            }
             res = curmax[i] > res ? curmax[i] : res;
        }
        return res;
    }
};
```
31、求出任意非负整数区间中1出现的次数（从1 到 n 中1出现的次数）。
```c++
class Solution {
public:
    int NumberOf1Between1AndN_Solution(int n)
    {
        int nums = 0;
        int m = 1;
        int lnum = 0, rnum = 0, mnum = 0;
        //按某位填1来算，如果那位真实值大于1，那么填1的所有情况都满足小于n
        //m为1,10,100...
        while(m <= n)
        {
            rnum = n % m;//当前位置右侧的数值
            mnum = n / m % 10;//当前位置数值
            lnum = n / m / 10;//当前位左侧数值

            /*
            若当前位的值大于1，那么左侧的数从 0 到 lnum（lnum+1） 个都可取， 当前位置取1， 右侧（m个）也随便取。
            若当前位的值==1，左侧则不能随便取（仅能取0到lnum-1, lnum个配上右侧随便，或左侧取lnum右侧取rnum内（rnum+1）个）
            若当前位的值小于1，当前置为 1 时，左侧只能取0到lnum-1，右侧随便取 m 个（0 ~ 999...）
            */

            if(mnum > 1) nums += (lnum + 1 ) * m;
            else if(mnum == 1) nums += (lnum * m) + (rnum + 1);
            else{
                nums += lnum * m;
            }
            m *= 10;
        }
        return nums;
    }
};
```
32、输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。例如输入数组{3，32，321}，则打印出这三个数字能排成的最小数字为321323。
```c++
//构造vector排序，比较的cmp
class Solution {
public:
    string PrintMinNumber(vector<int> numbers) {
        int n = numbers.size();
        string res = "";
        sort(numbers.begin(), numbers.end(), cmp);
        for(int i = 0; i < n; i++){
            res += to_string(numbers[i]);
        }
        return res;
    }
static bool cmp(int a, int b){
        string A = to_string(a) + to_string(b);
        string B = to_string(b) + to_string(a);
        return A < B;
    }
};
```
33、把只包含质因子2、3和5的数称作丑数（Ugly Number）。例如6、8都是丑数，但14不是，因为它包含质因子7。 习惯上我们把1当做是第一个丑数。求按从小到大的顺序的第N个丑数。
```c++
//典型的动态规划
/**
新的丑数必来自前面某一位的2,3,5倍，分别用t2位的2倍和t3位的3倍和t5位的5倍。
那么一开始t2 == t3 == t5 == 0
若新的丑数来自于t2位的2倍，那么下个新丑数只能在（t2+1）位的两倍和原来t3 t5倍数里选。
同理若来自于t3的3倍，那新的丑数还是应来自于t2位的两倍，（t3+1）位3倍，t5 5倍里。
1 (2 3 5)
1 2 (3 5 4)
1 2 3 (5 4 6)
1 2 3 4 (5 6 8)
...
*/
class Solution {
public:
    int GetUglyNumber_Solution(int index) {
        if(index < 7) return index;
        vector<int> res(index);
        res[0] = 1;
        int t2 = 0, t3 = 0, t5 = 0;
        for(int i = 1; i < index ; i++){
            res[i] = min(res[t2]*2, min(res[t3]*3, res[t5]*5));
            if(res[i] == res[t2]*2) t2++;
            if(res[i] == res[t3]*3) t3++;
            if(res[i] == res[t5]*5) t5++;
        }
        return res[index - 1];
    }
};
```
34、在一个字符串(0<=字符串长度<=10000，全部由字母组成)中找到第一个只出现一次的字符,并返回它的位置, 如果没有则返回 -1（需要区分大小写）.（从0开始计数）
```c++
//两次遍历，第一次对字母个数统计，第二次按str里的顺序得到位1的返回。
class Solution {
public:
    int FirstNotRepeatingChar(string str) {
       if (str.size() <= 0) return -1;
        int _hash[256] = {0};
        for(int i = 0; i < str.size(); i++){
            _hash[str[i]]++;
        }
        for(int i = 0; i < str.size(); i++){
            if(_hash[str[i]] == 1) return i;
        }
        return -1;
    }
};
```
35、
