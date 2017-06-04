    fun main(args: Array<String>) {
        var total = 0
        for (i in 0..998) {
            for (j in 1..998) {
                total += i / j
            }
        }
        System.out.println(total)
    }
