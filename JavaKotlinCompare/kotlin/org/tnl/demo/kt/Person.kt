package org.tnl.demo.kt

class Person(firstName: String, lastName: String, var address: Address? = null) {

    var firstName = firstName
        set(value) {
            field = value.capitalize()
        }

    var lastName = lastName
        set(value) {
            field = value.capitalize()
        }

    val fullName : String
        get() = "$firstName $lastName"

    override fun toString(): String {
        return "Person(firstName='$firstName', lastName='$lastName')"
    }

    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (other !is Person) return false

        if (firstName != other.firstName) return false
        if (lastName != other.lastName) return false

        return true
    }

    override fun hashCode(): Int {
        var result = firstName.hashCode()
        result = 31 * result + lastName.hashCode()
        return result
    }
}