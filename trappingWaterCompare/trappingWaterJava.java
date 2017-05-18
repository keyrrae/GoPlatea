public class Solution {
    public int trap(int[] height) {
        int n = height.length, l = 0, r = n - 1, water = 0, minHeight = 0;
        while (l < r) {
            while (l < r && height[l] <= minHeight)
                water += minHeight - height[l++];
            while (r > l && height[r] <= minHeight)
                water += minHeight - height[r--];
            minHeight = Math.min(height[l], height[r]);
        }
        return water;
    }
}