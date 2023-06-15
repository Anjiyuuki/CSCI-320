import java.io.File

fun main(args: Array<String>) {
    // Check if program ran with 1 command line argument
    if (args.size != 1) {
        println("Error: Enter 1 command line argument only")
        return
    }
    // input file name from the command line argument
    val fileName = args[0]

    // Create an empty map to hold the name initials and corresponding names
    val nameInitialsMap = mutableMapOf<String, String>()

    // reads the content of each line
    File(fileName).forEachLine { line ->
        //remove leading and trailing whitespace from the line
        val name = line.trim()
        //grabs the first letter
        val initials = name.split(" ").map { word -> word.first() }.joinToString("")
        nameInitialsMap[initials] = name
    }

    while (true) {
        try {
            println("Enter initials:")
            // if it = null then break such as typing ctrl D
            val initialsInput = readLine() ?: break
            // doesn't match
            val name = nameInitialsMap[initialsInput] ?: "Not Found"
            println(name)
        } catch (e: Exception) {
            break
        }
    }
}




