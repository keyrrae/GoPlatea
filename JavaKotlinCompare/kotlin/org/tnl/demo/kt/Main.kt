package org.tnl.demo.kt

fun printPersonAddresss(person: Person?) {
    println("Name: ${person?.fullName ?: "Unknown"}")
    println("Street: ${person?.address?.street ?: "Not specified"}")
    println("City: ${person?.address?.city ?: "Not specified"}")

    //person?.firstName.length
    person?.firstName.orEmpty().length
    person?.firstName.orEmpty()
}
