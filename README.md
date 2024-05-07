# mdrive-resetter
Application to connect to a MDrive 17 or 23 and send the correct sequence of commands to reset to factory default

# prerequisites
- Mdrive 17 or 23 with an exposed P2 connector
- serial com device that can send data using the RS-422 protocol

In my case this means: 
- USB to RS-232 serial cable
- 9 to 25 pin converter
- 25 to 25 pin passive pass through serial monitor (red and green LEDs per data or CTRL signal)
- 25 pin to 9 pin converter
- KK K2 RS-232 to RS-422/485 converter (Dip switch 2 and 3 switched to on for RS-422)
- A Delta 9 pin to 10 pin ribbon cable to connect to the Stepper motor (driver)

There is a 'secret' message that can be sent to the driver to reset it to factory default. It basicly means that the drive needs to be brought in a state where it can be reset.
This involves issueing it in a timely fasion with delays between the different commands.

The application
- Go application that can list the available serial ports
- a port can be selected
- The app iterates over the different supported baud rates with the most logical first (9600 baud default)
- It issues the commands with the provided timing (delays between commands)
- it tries to read the returned string after the reset.
- When the copyright notice is returned the app stops (should be confirmation of success), otherwise it continues with other baud rates.

### Pinout 10 pin header, notches at the left hand side 
``` text 
     --- ---
TX+  |9  10|  TX-
RX+   7   8|  RX-
Aux  |5   6|  RX+
RX-   3   4|  TX-
TX+  |1   2|  GND
     --- ---
```

### Delta 9 header
``` text 
TX+  9   10  TX-

RX+  7    8  RX-

Aux  5    6  RX+

RX-  3    4  TX-

TX+  1    2  GND
```
### Connection diagram
``` text
Drive(r)         KK K2 D9     
1 TX+     -->    7 RXB 
4 TX-     -->    2 RXA
3 RX-     <--    3 TXB
6 RX+     <--    8 TXA
```
