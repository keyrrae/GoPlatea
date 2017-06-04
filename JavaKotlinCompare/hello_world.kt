package helloworld

import kotlin.properties.Delegates

data class Name(val firstName: String,
                val middleName: String?,
                val lastName: String?)

trait NameFactory : Function0<Name?> {

    fun createName(): Name?

    fun defaultName(): Name

    override fun invoke(): Name? = this.createName()
}

class ConsoleNameFactory : NameFactory {

    public override fun createName(): Name? = readLine()?.let {
        println("[LOG] Creating name...")
        val nameParts = it.split(" ")
        val firstName = nameParts[0]

        val (middleName, lastName) =
                when (nameParts.size()) {
                    1 -> Pair(null, null)
                    2 -> Pair(null, nameParts[1])
                    else -> Pair(nameParts[1], nameParts[2])
                }

        Name(firstName, middleName, lastName)
    }

    override fun defaultName() = Name("", null, null)

}

class CachingNameFactory(delegate: NameFactory = ConsoleNameFactory()) : NameFactory by delegate {

    private val name by Delegates.lazy {
        delegate.createName()
    }

    override fun createName(): Name? = name

    override fun invoke(): Name? = super.invoke()
}

val nameProducer = { ->
    println("[LOG] Creating name")
    Name("", null, null)
}

inline fun memoize<reified T>(inlineOptions(InlineOption.ONLY_LOCAL_RETURN) body: () -> T): () -> T {
    var cache: T = null
    return { ->
        if (cache == null) {
            cache = body()
        }
        cache
    }
}

fun main(args: Array<String>) {
    app(memoize(nameProducer))
}

private fun app(nameFactory: () -> Name?) {
    println("Hello! I'm Bond. James Bond.")
    val name = nameFactory() ?: throw RuntimeException("Unnamed user")

    val line1 = "I'm ${name.firstName}..."
    val line2 = name.lastName?.let { "${name.lastName} ${name.firstName}..." } ?: ""
    val line3 = name.middleName?.let { "${name.lastName} ${name.firstName} ${name.middleName}" } ?: ""
    println(+"""
    Hello!
     $line1
     $line2
     $line3
    """)

    nameFactory()
}

fun String.plus() = this.trim().split("\n").
        map { it.trim() }.
        join("\n")