public class CachedEvaluation {
    public static String exampleChromosome = "011001110000011100100001110110000000010100110010010111110111010100111111110100011001000010011001111100101011111000001010011101101110100101011001100110001101101000000011000010110100011111001110100010101011001110010001110000111101000101010111100100000101110111101101011010000011000110001101001110110110111001011001101011110100011111011010001100010100010100101011110101100010001100101111011011000010110000101001001010101000101110110111110011101100001000100011111011101001010101111010101110011010001111000110111010011011001001011001100000000100110010111111100010010001011010110011010101001110010100111001001011100100100100100010110011110000100101010111110101101000001111001111101100011111001010101000010011001110100100011100101011000011100010101110011101111000101000110001010100100111100101111011010100001100010010101011000100010101110101010111101001110000110110101001000010011110111001100101111011100001010";

    private int n;
    private int k;
    private int[][] chooseCache;
    private int[][] edges;
    private int upperBound;
    private static char[] charBits;
    private static boolean[] bits;
    private static boolean[][] adj;

    /**
     * We are using R(5,5) for 43 nodes by default.
     */
    public CachedEvaluation() {
        this.n = 43;
        this.k = 5;
        this.chooseCache = new int[n][k];
        this.edges = getEdgePossibilities();
        this.upperBound = choose(n, k);
    }

    /**
     * I'm assuming that this will work for any number of nodes (n),
     * and any size clique (k), but I haven't tested it. Use at your own risk.
     */
    public CachedEvaluation(int n, int k) {
        this.n = n;
        this.k = k;
        this.chooseCache = new int[n][k];
        this.edges = getEdgePossibilities();
        this.upperBound = choose(n, k);
    }

    public int evaluate(String bitString) {
        int[] arr = { 0, 0, 0, 0, 0 };
        int numCliques = 0;
        int result = 0;

        charBits = bitString.toCharArray();
        bits = charToBoolean(charBits);
        adj = getAdjMatrixFromBoolArray(bits);
			
		/* Evaluate every possible clique */
        for (int i = 0; i < upperBound; i++) {
            getElement(i, arr);

            result = evalEdges(arr);

            if (result == 0 || result == 10) {
                numCliques++;
            }
        }

        return numCliques;
    }
    /**
     * Returns the number of edges in arr that are red (1).
     *
     * If this returns kC2 (10 for R(5, 5)), it is a red clique,
     * 				   0,   				 it is a blue clique
     */
    private int evalEdges(int[] arr) {
        int result = 0;

        for (int i = 0; i < edges.length; i++) {
            result += adj[arr[edges[i][0]]][arr[edges[i][1]]] ? 1 : 0;
        }

        return result;
    }

    /**
     * Returns the number nCk
     *
     * @param n
     * @param k
     * @return nCk
     *
     * @precondition k < n
     */
    private int choose(int n, int k) {
        if (n < k) {
            return 0;
        } if (n == k) {
            return 1;
        }
		
		/* Check to see if it's already cached */
        if (chooseCache[n - 1][k - 1] != 0) {
            return chooseCache[n - 1][k - 1];
        }

		/* Take advantage of the fact that nCk == nC(n-k) to do faster computation */
        int diff;
        int max;

        if (k < n - k) {
            diff = n - k;
            max = k;
        } else {
            diff = k;
            max = n - k;
        }

        int ans = diff + 1;

        for (int i = 2; i <= max; i++) {
            ans = (ans * (diff + i)) / i;
        }
		
		/* Cache answer before returning */
        chooseCache[n - 1][k - 1] = ans;

        return ans;
    }

    /**
     * Populates ans with the mth lexicographic subset of size k from n vertices (defined at top)
     */
    private void getElement(int m, int[] ans) {
        int a = n;
        int b = k;
        int x = (choose(n, k) - 1) - m; // x is the "dual" of m

        for (int i = 0; i < k; i++) {
            ans[i] = getLargestV(a, b, x); // largest value v, where v < a and vCb < x
            x = x - choose(ans[i], b);
            a = ans[i];
            b--;
        }

        for (int i = 0; i < k; i++) {
            ans[i] = (n - 1) - ans[i];
        }
    }

    /**
     * Same as above, but you can specify n and k.
     * I use this for getEdgePossibilities()
     */
    private int[] getElement(int m, int n, int k) {
        int a = n;
        int b = k;
        int x = (choose(n, k) - 1) - m; // x is the "dual" of m
        int[] ans = new int[k];

        for (int i = 0; i < k; i++) {
            ans[i] = getLargestV(a, b, x); // largest value v, where v < a and vCb < x
            x = x - choose(ans[i], b);
            a = ans[i];
            b--;
        }

        for (int i = 0; i < k; i++) {
            ans[i] = (n - 1) - ans[i];
        }

        return ans;
    }

    /**
     * Returns largest value v where v < a and Choose(v,b) <= x
     */
    private int getLargestV(int a, int b, int x) {
        int v = a - 1;

        while (choose(v, b) > x) {
            v--;
        }

        return v;
    }

    /**
     * Grab an array of edges to check for cliques. These are represented
     * as a tuple of indices into the array returned by getElement(m), e.g.:
     *
     * {
     * 	{0, 1},
     * 	{0, 2},
     *  {1, 2}
     * }
     *
     * For k = 3
     *
     */
    private int[][] getEdgePossibilities() {
        int numEdges = choose(k, 2);
        int[][] ans = new int[numEdges][2];

        for (int i = 0; i < numEdges; i++) {
            ans[i] = getElement(i, k, 2);
        }

        return ans;
    }

    private boolean[][] getAdjMatrixFromBoolArray(boolean[] arr) {
        int idx = 0;
        boolean[][] adj = new boolean[n][n];

        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                adj[i][j] = arr[idx];
                adj[j][i] = arr[idx];
                idx++;
            }
        }

        return adj;
    }

    /**
     * Maps 1 to true, 0 to false
     */
    private boolean[] charToBoolean(char[] bits) {
        boolean[] result = new boolean[bits.length];

        for (int i = 0; i < bits.length; i++) {
            result[i] = bits[i] == '1';
        }

        return result;
    }
}