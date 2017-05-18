int trap(int* height, int heightSize) { 
    int l = 0, r = heightSize - 1, water = 0, minHeight = 0;
    while (l < r) { 
        while (l < r && height[l] <= minHeight) 
            water += minHeight - height[l++];
        while (r > l && height[r] <= minHeight)
            water += minHeight - height[r--];
        minHeight = height[l] <= height[r] ? height[l] : height[r];
    }
    return water;
}