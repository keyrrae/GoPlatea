package def_methods

import kotlin.properties.Delegates

trait NameFactory {

    fun createName(): String

    fun trimmedName() = createName().trim()
}

class ConsoleNameFactory : NameFactory {

    override fun createName(): String {
        println("Creating name")
        return "name"
    }

}

class CachingNameFactory(delegate: NameFactory = ConsoleNameFactory()) : NameFactory by delegate {

    private val name by Delegates.lazy {
        delegate.createName()
    }

}

fun main(args: Array<String>) {
    val f = CachingNameFactory()
    f.createName() // Creating name are printed
    f.createName() // Creating name are printed too, while expected, that it will be cached
}
