import 'dart:io';

class RiceCooker {
  bool isOn = false;
  bool isPluggedIn = false;
  bool containsWater = false;
  bool cookedRice = false;

  void turnOnOrOff() {
    isOn = !isOn;
  }

  void plugInOrOff() {
    isPluggedIn = !isPluggedIn;
  }

  void addWaterOrEmpty() {
    containsWater = !containsWater;
  }

  void cookRice() {
    if (isOn && isPluggedIn && containsWater) {
      print('Cooking rice...');
      cookedRice = true;
    } else {
      print('Cannot cook rice. Make sure the rice cooker is on, plugged in, and contains water.');
    }
  }

  void displayStatus() {
    var status = isOn ? 'ON' : 'OFF';
    var pluggedInStatus = isPluggedIn ? 'Plugged In' : 'Not Plugged In';
    var waterStatus = containsWater ? 'Water Present' : 'No Water';
    var cookedRiceStatus = cookedRice ? 'Cooked Rice Ready' : 'No Cooked Rice';

    print('Status: $status\nPower: $pluggedInStatus\nWater: $waterStatus\nCooked Rice: $cookedRiceStatus');
  }
}

void main() {
  var riceCooker = RiceCooker();
  var scanner = stdin;

  while (true) {
    print('Rice Cooker Emulator:');
    print('1. Turn On/Off');
    print('2. Plug In/Off');
    print('3. Add Water/Empty');
    print('4. Cook Rice');
    print('5. Display Status');
    print('0. Exit');

    stdout.write('Enter your choice: ');
    var choice = scanner.readLineSync()?.trim();

    switch (choice) {
      case '1':
        riceCooker.turnOnOrOff();
        break;
      case '2':
        riceCooker.plugInOrOff();
        break;
      case '3':
        riceCooker.addWaterOrEmpty();
        break;
      case '4':
        riceCooker.cookRice();
        break;
      case '5':
        riceCooker.displayStatus();
        break;
      case '0':
        print('Exiting Rice Cooker Emulator.');
        exit(0);
      default:
        print('Invalid choice. Please enter a valid option.');
    }
  }
}
