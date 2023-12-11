import model.RiceCooker

import scala.io.StdIn.readLine

object Main extends App {
  val riceCooker = new RiceCooker

  while (true) {
    println("Rice Cooker Emulator:")
    println("1. Turn On/Off")
    println("2. Plug In/Off")
    println("3. Add Water/Empty")
    println("4. Cook Rice")
    println("5. Display Status")
    println("0. Exit")

    print("Enter your choice: ")
    val choice: String = readLine().trim

    choice match {
      case "1" => riceCooker.turnOnOrOff()
      case "2" => riceCooker.plugInOrOff()
      case "3" => riceCooker.addWaterOrEmpty()
      case "4" => riceCooker.cookRice()
      case "5" => riceCooker.displayStatus()
      case "0" =>
        println("Exiting Rice Cooker Emulator.")
        System.exit(0)
      case _ => println("Invalid choice. Please enter a valid option.")
    }
  }
}
