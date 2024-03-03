# Second_Monitor_Swicher
Keyboard Display Switcher

This Go program listens for specific keyboard events and switches the display between internal and external monitors accordingly. It utilizes the go-hook package to capture keyboard events and triggers actions based on the keys pressed.
Features

Listens for specific keys to trigger display switching.     Uses the DisplaySwitch.exe utility from the Windows System32 directory to toggle between internal and external displays.  

Usage

Clone the repository to your local machine.     
Build the program using Go compiler. 
Run the executable.   
Press the  ` key to switch display to second screen only mod  and F2 for pc only mod   
or you can set your own keys


Dependencies

   moutend/go-hook: A Go package for Windows global hook.
   
Note

Currently, the program is designed for Windows platforms due to its reliance on DisplaySwitch.exe.
    
Feel free to contribute, report issues, or suggest improvements!
