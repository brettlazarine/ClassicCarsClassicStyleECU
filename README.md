# ClassicCarsClassicStyleECU
The ECU (Electronic Control Unit) repository for the Classic Cars Classic Style project -> https://github.com/brettlazarine/ClassicCarsClassicStyleMobile

## Why does this project exist?
Classic Cars Classic Style is a large project, divided into two componenets. This repo focuses on the back-end, the ECU itself.
Older vehicles did not always have many luxuries, and some had none at all. Take my old 1964 F250 for example:
- No radio
- No A/C (I live in Texas, this is *unforgivable*)
- A gas tank that is stored behind the driver in the cab of the pickup (keep your windows rolled down)
The problem is that these vehicles have a timeless aesthetic that would only by compromised by modernizing.
My 250 is fully steel on the inside, it would look ridiculous to slap an iPad on the dash just to have maps and music available.

## The solution
I decided to address this myself. 
In tandem with the Mobile App I am developing, this repo will serve as the ECU and logic for my project.
Using a Raspberry Pi 5, I am building an intelligent solution to a timeless problem: 
**How do I get the most out of my classic car without ruining the way it looks?**

## How
This ECU is going to be built out, at least primarily, in Go. 
The Raspberry Pi will act as a server host for multiple I/O outlets:
- A/C
- Speakers and Sound System
- Backup Camera
Communication will begin client-side with the Mobile App, which will connect directly to the Pi via Wireless Access Point.
Once connected, all aforementioned modules will be controlled remotely with control loops running on the Pi.

## The nitty gritty
As this is a live implementation, there will be physical implementations as well.
The Pi ECU and each of its modules will be powered by the vehicle's battery, with protective relays, fuses, and converters in place.

## Contributing
Collaboration is always appreciated, feel free to reach out to me on LinkedIn!
https://www.linkedin.com/in/brettlazarine
