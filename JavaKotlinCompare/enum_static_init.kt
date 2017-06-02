package enum_static_init

public enum class Test(val id: String) {
    A : Test("id")

    companion object {

        val byId = values().map { Pair(it.id, it) }.toMap()

        fun byId(id: String) = byId[id]

    }
}

fun main(args: Array<String>) {
    println(Test.byId("id"))
}