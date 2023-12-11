package model

class RiceCooker {
  private var isOn: Boolean = false
  private var isPluggedIn: Boolean = false
  private var containsWater: Boolean = false
  private var cookedRice: Boolean = false

  def turnOnOrOff(): Unit = {
    isOn = !isOn
  }

  def plugInOrOff(): Unit = {
    isPluggedIn = !isPluggedIn
  }

  def addWaterOrEmpty(): Unit = {
    containsWater = !containsWater
  }

  def cookRice(): Unit = {
    if (isOn && isPluggedIn && containsWater) {
      println("Cooking rice...")
      cookedRice = true
    } else {
      println("Cannot cook rice. Make sure the rice cooker is on, plugged in, and contains water.")
    }
  }

  def displayStatus(): Unit = {
    val status: String = if (isOn) "ON" else "OFF"
    val pluggedInStatus: String = if (isPluggedIn) "Plugged In" else "Not Plugged In"
    val waterStatus: String = if (containsWater) "Water Present" else "No Water"
    val cookedRiceStatus: String = if (cookedRice) "Cooked Rice Ready" else "No Cooked Rice"

    println(s"Status: $status\nPower: $pluggedInStatus\nWater: $waterStatus\nCooked Rice: $cookedRiceStatus")
  }
}