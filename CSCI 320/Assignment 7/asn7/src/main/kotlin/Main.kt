fun main(args: Array<String>) {

    val words = args.toList() // list of words in command line

    // map and mean
    val mean = words.map { letter -> letter.length }.average()
    println("Mean word length: $mean")

   // map and deviation
    val dev = words.map { (it.length - mean) * (it.length - mean) }.average().let { Math.sqrt(it) }
    println("Standard deviation: $dev")
}