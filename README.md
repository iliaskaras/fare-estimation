# fare-estimation CLI
#### A CLI tool written in GO for calculating fare estimates from files containing ride position entries.

## Running instructions

* Download the repository.
* Change directory to **fare-estimation** where the Makefile is located.
* In the Makefile there are two commands:
  * make install-cli-tool
    * This will build and install the CLI tool.
  * make run-tests
    * Will run all the project tests.
* Now you can run the fare-estimation using the installed cli tool:
```
fare-estimation estimate -f <input-file.csv> -o <output-filename>
```
* Run example with the paths file in the resources, note that you will have to be in the project's root directory:
```
fare-estimation estimate -f resources/paths.csv -o resources/estimated_fares.csv
```

## Fare estimation process logic
* File parsing: The file is parsed line by line, and pushes to the ridePositionsChan the RidePositions of a specific 
  RideID. 
  * Pusher to the ridePositionsChan.
* Filtering on Segment speed: Receiving all the RideID's RidePositions, and creates the RideSegments.
  * Here is where the filtering on segment speed is happening, only Segments that passes the sanity check are pushed
    to the rideSegmentsChan, for fare estimation later on. A single RideSegments that is pushed to the channel,
    contain all the filtered segments of a single RideID.
  * Receiver to the ridePositionsChan.
  * Pusher to the rideSegmentsChan.
  * Due to the fact that I found after stress test that there is a bottleneck between File parsing and Filtering steps,
    I wrapped the filtering step into a wait group of 4, making more concurrent receivers on the ridePositionsChan. 
* Fare estimation: Calculates the fares on the filtered ride segments.
  * Receiver to the rideSegmentsChan.
  * Pusher to the faresChan.
* File writer: Writes line by line the produced fares.
  * Receiver to the faresChan.

## Project Information
- General Information: This was the first "real" project I wrote in Golang, I hope there aren't many mistakes. Thanks
for the review!
- Structure: I tried to keep the project structure exactly as I used to in the projects I am developing in Python / Java.
That is, having a main app directory, and submodules that represent the app's resources. Each module have its own
models, services, factories and errors, if all are needed, depending on the resource. Usually, I have separated the
modules containing the resource models, services and errors with resource modules containing the controllers. In this
project, kept the "controllers" in the cmd directory, which the CLI Cobra tool needs to have the commands under.
- In general, I did some "over-engineering" in some parts, like for example, the distance and file reader factories,
which initializes an interface implementor depending on some logic. This is to make project potentially easier to
be modified and be mocked.
- Services: Contain the methods with the business logic.
- Models: The resource models, in Go, simple resource structs, which have no methods with business logic assigned to
them.
- Errors: Resource specific errors, keep note that all the app's resources errors implements the base application error
from infrastructure.
- Factories: Methods that take care of injecting the dependencies and initializing the Services. Used by the commands,
and decouples the initialization logic and maintainability from the client.

## To improve
- The commands: I did not have time to remove some synchronization logic from the estimate command. A service should
be created which taking care of wrapping all the calls necessary for the fare estimation, which would be called by
the command.
- Testing: I have created unit-tests only but tried to cover as much logic as possible, but currently there aren't
e2e tests, like testing the command for example, and all the inner pieces together.
- Code: As I said, first project in GO, I strongly believe that there could be more efficient way of solving this
problem, I hope I am not that far away.

## Stress test
- The given file is parsed in few milliseconds. 
- On 2GB file that I created and used for stress test: ~20seconds e2e.
  - Note: Just parsing the file line by line and unmarshal the rows, took close to the full amount. Making all
    the rest steps seem pretty efficient.
